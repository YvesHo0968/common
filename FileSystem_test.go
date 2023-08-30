package common

import (
	"fmt"
	"os"
	"testing"
)

func TestBasename(t *testing.T) {
	fmt.Println(Basename("/path/to/file.txt"))
	fmt.Println(Basename("/path/to/file.txt", ".txt"))
}

func TestChGrp(t *testing.T) {
	fmt.Println(ChGrp("/path/to/file.txt", "everyone"))
}

func TestChmod(t *testing.T) {
	fmt.Println(Chmod("/path/to/file.txt", "0777"))
}

func TestChown(t *testing.T) {
	fmt.Println(Chown("/path/to/file.txt", "501"))
	fmt.Println(Chown("/path/to/file.txt", "everyone"))
}

func TestCopy(t *testing.T) {
	fmt.Println(Copy("/path/to/file.txt", "/path/to/file_copy.txt"))
}

func TestDirname(t *testing.T) {
	fmt.Println(Dirname("/path/to/file.txt"))
}

func TestDiskFreeSpace(t *testing.T) {
	fmt.Println(DiskFreeSpace("/"))
}

func TestDiskTotalSpace(t *testing.T) {
	fmt.Println(DiskTotalSpace("/"))
}

func TestFile(t *testing.T) {
	for k, v := range File("./math.go") {
		fmt.Println(k, v)
	}
}

func TestFileExists(t *testing.T) {
	fmt.Println(FileExists("/path/to/file.txt"))
}

func TestFileGetContents(t *testing.T) {
	fmt.Println(FileGetContents("/path/to/file.txt"))
}

func TestFilePullContents(t *testing.T) {
	fmt.Println(FilePullContents("/path/to/file.txt", "O_TRUNC\n", os.O_TRUNC))
	fmt.Println(FilePullContents("/path/to/file.txt", "O_APPEND", os.O_APPEND))
}

func TestFileAtime(t *testing.T) {
	fmt.Println(FileAtime("/Users/hezhiyi/Downloads/check.env"))
}

func TestFileCtime(t *testing.T) {
	fmt.Println(FileCtime("/Users/hezhiyi/Downloads/check.env"))
}

func TestFileGroup(t *testing.T) {
	fmt.Println(FileGroup("/Users/hezhiyi/Downloads/check.env"))
}

func TestFileInode(t *testing.T) {
	fmt.Println(FileInode("/Users/hezhiyi/Downloads/check.env"))
}

func TestFileMtime(t *testing.T) {
	fmt.Println(FileMtime("/Users/hezhiyi/Downloads/check.env"))
}

func TestFileOwner(t *testing.T) {
	fmt.Println(FileOwner("/Users/hezhiyi/Downloads/check.env"))
}
