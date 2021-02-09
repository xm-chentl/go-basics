package stringex

import (
	"strings"

	objectid "github.com/xm-chentl/go-basics/stringex/guid/object-id"
	randstring "github.com/xm-chentl/go-basics/stringex/rand-string"
)

var (
	// Generator is 创建Guid生成器
	Generator = objectid.New()
)

// GUID is 生成GUID
func GUID() string {
	return Generator.Generate()
}

// Join is 使用分隔符拼接字符串
func Join(sep string, s ...string) string {
	if s == nil || len(s) == 0 {
		return ""
	}

	return strings.Join(s, sep)
}

// Rand is 随机字符串
func Rand(n int) string {
	return randstring.Create(n)
}
