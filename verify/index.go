package verify

import (
	"regexp"

	"github.com/xm-chentl/go-basics/verify/enum"
)

// IsDate is 是否是日期
func IsDate(value string) bool {
	regex, err := regexp.Compile(enum.RegexpDate)
	if err != nil {
		return false
	}

	return regex.MatchString(value)
}

// IsDateTime is 是否是日期时间
func IsDateTime(value string) bool {
	regex, err := regexp.Compile(enum.RegexpDateTime)
	if err != nil {
		return false
	}

	return regex.MatchString(value)
}

// IsPhone is 手机格式是否正确
func IsPhone(value string) bool {
	regex, err := regexp.Compile(enum.RegexpPhone)
	if err != nil {
		return false
	}

	return regex.MatchString(value)
}

// IsIP is IP格式是否正确
func IsIP(value string) bool {
	regex, err := regexp.Compile(enum.RegexpIP)
	if err != nil {
		return false
	}

	return regex.MatchString(value)
}

// IsEmpty 判断是否为空
func IsEmpty(keys []string, values []string) ([]string, bool) {
	results := make([]string, 0)
	for index := 0; index < len(keys); index++ {
		if values[index] == "" {
			results = append(results, keys[index])
		}
	}
	if len(results) > 0 {
		return results, true
	}

	return nil, false
}
