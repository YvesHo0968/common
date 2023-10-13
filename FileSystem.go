package common

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
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
		return false
	}

	err = os.Chmod(filePath, os.FileMode(permissionsInt))
	if err != nil {
		return false
	}
	return true
}

// Chown 改变文件所有者
func Chown(filePath string, owner string) bool {
	u, err := user.Lookup(owner)
	if err != nil {
		u, err = user.LookupId(owner)
		if err != nil {
			return false
		}
	}
	uid, err := strconv.Atoi(u.Uid)
	if err != nil {
		return false
	}
	gid, err := strconv.Atoi(u.Gid)
	if err != nil {
		return false
	}
	err = os.Chown(filePath, uid, gid)
	if err != nil {
		return false
	}
	return true
}

// Copy 复制文件
func Copy(source string, destination string) bool {
	srcFile, err := os.Open(source)
	if err != nil {
		return false
	}
	defer func(srcFile *os.File) {
		_ = srcFile.Close()
	}(srcFile)

	destFile, err := os.Create(destination)
	if err != nil {
		return false
	}
	defer func(destFile *os.File) {
		_ = destFile.Close()
	}(destFile)
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return false
	}
	return true
}

// Dirname 返回路径中的目录名称部分
func Dirname(filePath string) string {
	return filepath.Dir(filePath)
}

// DiskFreeSpace 返回目录的可用空间
func DiskFreeSpace(directory string) uint64 {
	var stat syscall.Statfs_t
	err := syscall.Statfs(directory, &stat)
	if err != nil {
		return 0
	}
	// 可用空间 = 块大小 × 可用块数
	return stat.Bavail * uint64(stat.Bsize)
}

// DiskTotalSpace 返回一个目录的磁盘总容量
func DiskTotalSpace(directory string) uint64 {
	var stat syscall.Statfs_t
	err := syscall.Statfs(directory, &stat)
	if err != nil {
		return 0
	}
	// 总空间大小 = 总块数 * 每块的大小
	return stat.Blocks * uint64(stat.Bsize)
}

// FClose 关闭打开的文件
func FClose(file *os.File) error {
	return file.Close()
}

// File 把文件读入一个数组中
func File(filePath string) []string {
	content := FileGetContents(filePath)
	return strings.Split(content, "\n")
}

// FileExists 把文件读入一个数组中
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return true
}

// FileGetContents 把文件读入字符串
func FileGetContents(filePath string) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return ""
	}
	return string(content)
}

// FilePullContents 把字符串写入文件
func FilePullContents(filePath, content string, flag int) int {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|flag, 0644)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return 0
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	n, err := io.WriteString(file, content)

	if err != nil {
		fmt.Println("Failed to write to file:", err)
		return 0
	}

	return n
}

// FileAtime 返回文件的上次访问时间
func FileAtime(filePath string) int64 {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0
	}
	// 获取文件的上次访问时间
	return fileInfo.Sys().(*syscall.Stat_t).Atimespec.Sec
}

// FileCtime 返回文件的上次修改时间
func FileCtime(filePath string) int64 {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0
	}
	// 获取文件的上次访问时间
	return fileInfo.Sys().(*syscall.Stat_t).Ctimespec.Sec
}

// FileGroup 返回文件的组 ID
func FileGroup(filePath string) int {
	uid := FileOwner(filePath)
	userInfo, err := user.LookupId(fmt.Sprint(uid))
	if err != nil {
		return 0
	}
	// 获取所属组名
	groupName, err := user.LookupGroupId(userInfo.Gid)
	if err != nil {
		return 0
	}

	groupId, err := strconv.Atoi(groupName.Gid)
	if err != nil {
		return 0
	}

	return groupId
}

// FileInode 返回文件的 inode 编号
func FileInode(filePath string) uint64 {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0
	}
	// 获取文件的inode号
	return fileInfo.Sys().(*syscall.Stat_t).Ino
}

// FileMtime 返回文件内容的上次修改时间
func FileMtime(filePath string) int64 {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0
	}
	// 获取文件的上次访问时间
	return fileInfo.Sys().(*syscall.Stat_t).Mtimespec.Sec
}

// FileOwner 返回文件的用户 ID （所有者）
func FileOwner(filePath string) int {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0
	}
	// 获取文件的所属用户
	uid := fileInfo.Sys().(*syscall.Stat_t).Uid
	return int(uid)
}

// FilePerms 返回文件的权限
func FilePerms(filePath string) uint16 {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0
	}
	return fileInfo.Sys().(*syscall.Stat_t).Mode
}

// FileSize 返回文件大小
func FileSize(filePath string) int64 {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0
	}
	return fileInfo.Sys().(*syscall.Stat_t).Size
}

// FileType 返回文件类型
func FileType(filePath string) string {
	if IsDir(filePath) {
		return "dir"
	}
	return "file"
}

// IsDir 判断文件是否是一个目录
func IsDir(filePath string) bool {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return false
	}

	if fileInfo.IsDir() {
		return true
	}

	return false
}

// IsExecutable 判断文件是否可执行
func IsExecutable(filename string) bool {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fileMode := fileInfo.Mode()
	return fileMode&0111 != 0
}

// IsFile 判断文件是否是常规的文件
func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// IsLink 判断文件是否是连接
func IsLink(path string) bool {
	info, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return info.Mode()&os.ModeSymlink != 0
}
