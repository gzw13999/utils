package utils

import (
	"os"
	"strconv"
)

func SubStr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

func FileIsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func InIntSlice(islice []int, i int) bool {
	for _, value := range islice {
		if value == i {
			return true
		}
	}
	return false
}

func InStringSlice(islice []string, i string) bool {
	for _, value := range islice {
		if value == i {
			return true
		}
	}
	return false
}

func Atoi(s string, defaultVal ...int) int {

	if i, err := strconv.Atoi(s); err != nil {
		return defaultVal[0]
	} else {
		return i
	}

}

func Itoa(i int) string {

	return strconv.Itoa(i)

}
