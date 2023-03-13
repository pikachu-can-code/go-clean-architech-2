package common

import (
	"context"
)

func ContextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return ErrRequestCanceled()
	case context.DeadlineExceeded:
		return ErrDeadlineExceeded()
	default:
		return nil
	}
}
