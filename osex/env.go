package osex

import (
	"fmt"
	"os"
)

// MockEnv is 环境模拟
var MockEnv = make(map[string]string)

// GetenvWithPainc is 获取环境变量
func GetenvWithPainc(k string) string {
	if v, ok := MockEnv[k]; ok {
		return v
	}

	v := os.Getenv(k)
	if v == "" {
		panic(
			fmt.Sprintf("请先配置环境变量: %s", k),
		)
	}

	return v
}
