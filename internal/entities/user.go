package entities

import (
	"errors"

	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/common"
)

const userEntityName = "users"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;size:300;"`
	Password        string        `json:"-" gorm:"column:password;size:500;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;size:255;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;size:300;"`
	Phone           string        `json:"phone" gorm:"column:phone;size:15;"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (u *User) GetUserId() uint64 {
	return u.ID
}

func (u *User) GetEmail() string {
	return u.Email
}

func (User) TableName() string {
	return userEntityName
}

var (
	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExited",
	)

	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password is incorrect"),
		"email or password is incorrect",
		"ErrEmailOrPasswordInvalid",
	)
)
