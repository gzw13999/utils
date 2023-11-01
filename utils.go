package utils

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
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

func TimeSubDays(t1, t2 time.Time) int {
	if t1.Location().String() != t2.Location().String() {
		return -1
	}
	hours := t1.Sub(t2).Hours()

	if hours <= 0 {
		return -1
	}
	// sub hours less than 24
	if hours < 24 {
		// may same day
		t1y, t1m, t1d := t1.Date()
		t2y, t2m, t2d := t2.Date()
		isSameDay := (t1y == t2y && t1m == t2m && t1d == t2d)

		if isSameDay {

			return 0
		} else {
			return 1
		}

	} else { // equal or more than 24

		if (hours/24)-float64(int(hours/24)) == 0 { // just 24's times
			return int(hours / 24)
		} else { // more than 24 hours
			return int(hours/24) + 1
		}
	}

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

func StrLen(s string) int {
	if s == "" {
		return 0
	} else {
		return len([]rune(s))
	}

}

func Atoi(s any, defaultVal ...int) int {
	switch s := s.(type) {
	case string:
		if i, err := strconv.Atoi(s); err != nil {
			if len(defaultVal) == 0 {
				return 0
			}
			return defaultVal[0]
		} else {
			return i
		}
	default:
		if len(defaultVal) == 0 {
			return 0
		}
		return defaultVal[0]
	}
}

func Atob(s any) bool {
	switch s := s.(type) {
	case string:
		if i, err := strconv.ParseBool(s); err != nil {
			return false
		} else {
			return i
		}
	default:
		return false
	}
}

func Itoa(i int) string {
	return strconv.Itoa(i)
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

// uint 代表小数位数，格式位 0.000001 如果是几位就指定为几位,有多余小数情况
func truncateNaive(f float64, unit float64) float64 {
	return math.Trunc(f/unit) * unit
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.7f", value), 64)
	return value
}
func ToFixed(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	fv := 0.0000000001 + f //对浮点数产生.xxx999999999 计算不准进行处理
	//return math.Floor(fv*shift+.5) / shift //四舍五入
	return math.Floor(fv*shift) / shift
}
