package stringex

import (
	"github.com/google/uuid"
)

// GUID 返回GUID
func GUID() string {
	return uuid.NewString()
}
