package httprequest

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/common"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/components/logging"
)

func MakeRequest(method string, url string, body io.Reader, headers map[string]string) (io.ReadCloser, error) {
	logger := logging.NewAPILogger()
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	if len(headers) == 0 {
		req.Header.Set("Content-Type", "application/json")
	} else {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, common.NewFullErrorResponse(http.StatusServiceUnavailable, err, "Lỗi dịch vụ!", err.Error(), "ErrServiceUnavailable")
	}

	logger.Infof("Request URL: %v", url)
	logger.Infof("Request Status: %v", res.Status)

	if res.StatusCode >= 300 && res.StatusCode < 600 {
		appErr := common.AppError{}
		byteData, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, common.ErrInternal(err)
		}
		err = json.Unmarshal(byteData, &appErr)
		if err != nil {
			return nil, common.NewFullErrorResponse(res.StatusCode, err, "Lỗi dịch vụ!", err.Error(), "ErrServiceUnavailable")
		}
		appErr.RootErr = errors.New(appErr.Log)
		return nil, &appErr
	}

	return res.Body, nil
}

func MakeHeader(ctx context.Context, authenToken, userAgent, cookie string) (header map[string]string) {
	header = make(map[string]string)
	if ctx.Value(common.TokenUser) != nil && ctx.Value(common.TokenUser).(string) != "" {
		authenToken = ctx.Value(common.TokenUser).(string)
	}
	header["accept-encoding"] = "gzip, deflate, br"
	header["user-agent"] = userAgent
	header["cookie"] = cookie
	header["authorization"] = "Bearer " + authenToken
	return
}
