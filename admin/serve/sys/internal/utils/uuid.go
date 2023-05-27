package utils

import (
	"github.com/google/uuid"
)

// 生成uuid
func GenerateUUID() string {
	// 生成一个随机的 UUID
	u := uuid.New()
	// 返回 UUID 的字符串表示
	return u.String()
}
