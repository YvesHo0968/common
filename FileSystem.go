package common

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
)

// Basename 返回路径中的文件名部分
func Basename(filePath string, suffix ...string) string {
	baseName := filepath.Base(filePath)

	if len(suffix) > 0 {
		baseName = strings.TrimSuffix(baseName, suffix[0])
	}

	return baseName
}

// ChGrp 改变文件组
func ChGrp(filePath, groupName string) bool {
	group, err := user.LookupGroup(groupName)
	if err != nil {
		return false
	}

	groupID, err := strconv.Atoi(group.Gid)
	if err != nil {
		return false
	}

	err = os.Chown(filePath, -1, groupID)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return true
}

// Chmod 改变文件模式
func Chmod(filePath, mode string) bool {
	permissionsInt, err := strconv.ParseInt(mode, 8, 32)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	err = os.Chmod(filePath, os.FileMode(permissionsInt))
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return true
}
