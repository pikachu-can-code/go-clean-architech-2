package common

const (
	ping = iota + 1
	userInfo
)

var PermissionRules = map[string][]int{
	"/user/info": {userInfo},
}
