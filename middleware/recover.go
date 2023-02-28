package middleware

import (
	"fmt"
	"runtime/debug"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/common"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/components/logging"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func errorHandler(p interface{}) error {
	logger := logging.NewAPILogger()
	if appErr, ok := p.(*common.AppError); ok {
		logger.Errorf("[error] %v", appErr.Error())
		status := status.New(appErr.StatusCode, appErr.Message)
		detail := &errdetails.ErrorInfo{
			Reason:   appErr.Error(),
			Metadata: map[string]string{"Key": appErr.Key},
		}
		st, _ := status.WithDetails(detail)
		return st.Err()
	}
	if err, isErr := p.(error); isErr {
		if grpcErr, ok := status.FromError(err); ok {
			debug.PrintStack()
			logger.Errorf("[error] %v", grpcErr.Err())
			return grpcErr.Err()
		}
	}

	logger.Errorf("[error unknow] %v", p)
	debug.PrintStack()
	status := status.New(codes.Unknown, "Đã có lỗi xảy ra!")
	detail := &errdetails.ErrorInfo{
		Reason: fmt.Sprintf("error %v", p),
	}
	st, _ := status.WithDetails(detail)
	return st.Err()
}

var RecoverOptions = []grpc_recovery.Option{
	grpc_recovery.WithRecoveryHandler(errorHandler),
}
