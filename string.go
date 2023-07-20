package common

import "fmt"

// StrVal 任意类型转字符串
func StrVal(data any) string {
	return fmt.Sprintf("%v", data)
}
