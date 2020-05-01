package utils

import (
	"regexp"
)

//匹配电子邮箱
func VerifyEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//匹配国内手机号码
func VerifyPhone(s string) bool {
	reg := regexp.MustCompile(`^1(3|4|5|6|7|8)\d{9}$`)
	return reg.MatchString(s)
}
