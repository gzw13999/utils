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

func FileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func InIntSlice(s []int, i int) bool {
	for _, value := range s {
		if value == i {
			return true
		}
	}
	return false
}

func InSlice(s []*interface{}, i interface{}) bool {
	for _, value := range s {
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

func Atob(s string) bool {
	if s == "1" || s == "true" {
		return true
	} else {
		return false
	}
}

func Atof32(s string, defaultVal ...float32) float32 {
	v, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return defaultVal[0]
	} else {
		return float32(v)
	}
}

func Atof64(s string, defaultVal ...float64) float64 {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return defaultVal[0]
	} else {
		return v
	}
}
