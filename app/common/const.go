package common

import "time"

const (
	DbTypeCommon = iota + 1
	DbTypeUser
	DbTypeRole
	DbTypePermission
)

const (
	CurrentUser = "__user__"
	TokenUser   = "__token__"
)

const TimeLayout = time.RFC3339

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRoles() []int
	IsAdmin() bool
}
