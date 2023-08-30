package common

import (
	"testing"
)

func TestDie(t *testing.T) {
	Die("die")
}

func TestExit(t *testing.T) {
	Exit("exit")
}
