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

func TestFClose(t *testing.T) {
	file, err := os.Open("./math.go")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		err := FClose(file)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Close")
		}
	}()
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
	fmt.Println(FileAtime("/path/to/file.txt"))
}

func TestFileCtime(t *testing.T) {
	fmt.Println(FileCtime("/path/to/file.txt"))
}

func TestFileGroup(t *testing.T) {
	fmt.Println(FileGroup("/path/to/file.txt"))
}

func TestFileInode(t *testing.T) {
	fmt.Println(FileInode("/path/to/file.txt"))
}

func TestFileMtime(t *testing.T) {
	fmt.Println(FileMtime("/path/to/file.txt"))
}

func TestFileOwner(t *testing.T) {
	fmt.Println(FileOwner("/path/to/file.txt"))
}

func TestFilePerms(t *testing.T) {
	fmt.Println(fmt.Sprintf("%o", FilePerms("/path/to/file.txt")))
}

func TestFileSize(t *testing.T) {
	fmt.Println(FileSize("/path/to/file.txt"))
}

func TestFileType(t *testing.T) {
	fmt.Println(FileType("/path/to/file.txt"))
}

func TestIsDir(t *testing.T) {
	fmt.Println(IsDir("/path/to/file.txt"))
}

func TestIsExecutable(t *testing.T) {
	fmt.Println(IsExecutable("/path/to/file.txt"))
}

func TestIsFile(t *testing.T) {
	fmt.Println(IsFile("/path/to/file.txt"))
}

func TestIsLink(t *testing.T) {
	fmt.Println(IsLink("/path/to/file.txt"))
}
