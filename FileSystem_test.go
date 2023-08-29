package common

import (
	"fmt"
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
