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
	pattern := `^(((?:19|20)\d\d)-(0?[1-9]|1[0-2])-(0?[1-9]|[12][0-9]|3[01]))$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(date)
}

//匹配国内手机号码
func VerifyPhone(s string) bool {
	reg := regexp.MustCompile(`^1(3|4|5|6|7|8|9)\d{9}$`)
	return reg.MatchString(s)
}

//匹配密码
func VerifyUsername(s string) bool {
	reg := regexp.MustCompile("^[a-zA-Z0-9_-]{6,16}$")
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
