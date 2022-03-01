package common

import (
	"regexp"
)

func ValidateEmail(email string) bool {
	regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return regex.MatchString(email)
}

func ValidatePhone(phone string) bool {
	regex := regexp.MustCompile(`(03|07|08|09|01[2|6|8|9])+([0-9]{8,9})\b`)
	return regex.MatchString(phone)
}
