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

//匹配日期
func VerifyDate(date string) bool {
	pattern := `\d{4}-\d{2}-\d{2}`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(date)
}

//匹配国内手机号码
func VerifyPhone(s string) bool {
	reg := regexp.MustCompile(`^1(3|4|5|6|7|8)\d{9}$`)
	return reg.MatchString(s)
}

//匹配密码
func VerifyPassword(s string) bool {

	reg := regexp.MustCompile("^[a-zA-Z0-9-*/+.~!@#$%^&*()]{6,20}$")
	return reg.MatchString(s)
}

func IsGuid(s string) bool {
	if len(s) != 36 {
		return false
	}

	reg := regexp.MustCompile("^[a-fA-F0-9]{8}(-[a-fA-F0-9]{4}){3}-[a-fA-F0-9]{12}$")
	if reg.MatchString(s) {
		return true
	} else {
		return false
	}
}

func IsNumber(s string) bool {
	reg := regexp.MustCompile("^[0-9]+$")
	return reg.MatchString(s)
}

func IsABC(s string) bool {
	reg := regexp.MustCompile("^[a-zA-Z]+$")
	return reg.MatchString(s)
}
