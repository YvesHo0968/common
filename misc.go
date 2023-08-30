package common

import (
	"fmt"
	"os"
)

// Die 输出一条消息，并退出当前脚本
func Die(message string) {
	fmt.Println(message)
	os.Exit(1)
}

// Exit 输出一条消息，并退出当前脚本
func Exit(message string) {
	Die(message)
}
