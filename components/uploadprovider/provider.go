package uploadprovider

import (
	"context"

	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/common"
)

type UploadProvider interface {
	SaveImageUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
