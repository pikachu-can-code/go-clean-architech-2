package tokenprovider

import (
	"errors"
	"net/http"

	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/common"
)

type Token struct {
	AuthId  int64  `json:"-"`
	Token   string `json:"token"`
	Created int64  `json:"created"`
	Expiry  int64  `json:"expiry"`
}

type TokenPayload struct {
	UserID uint64   `json:"user_id"`
	Role   []string `json:"role"`
}

type Provider interface {
	Generate(data TokenPayload, expiry int64) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

var (
	ErrNotFound = common.NewFullErrorResponse(
		http.StatusUnauthorized,
		errors.New("token not found"),
		"Không tìm thấy token!",
		"token not found",
		"ErrNotFound",
	)

	ErrEncodingToken = common.NewFullErrorResponse(
		http.StatusUnauthorized,
		errors.New("error encoding the token"),
		"token not found",
		"Lỗi khi tạo token",
		"ErrEncodingToken",
	)

	ErrInvalidToken = common.NewFullErrorResponse(
		http.StatusUnauthorized,
		errors.New("invalid token provided"),
		"Token xác thực không đúng!",
		"invalid token provideds",
		"ErrInvalidToken",
	)
)
