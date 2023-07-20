package common

import (
	"fmt"
	"testing"
)

func TestHostInfo(t *testing.T) {
	fmt.Println(HostInfo())
}

func TestCpuInfo(t *testing.T) {
	fmt.Println(CpuInfo())
}

func TestMemoryInfo(t *testing.T) {
	fmt.Println(MemoryInfo())
}

func TestDiskInfo(t *testing.T) {
	fmt.Println(DiskInfo())
}

func TestNetInterfaces(t *testing.T) {
	fmt.Println(NetInterfaces())
}

func TestNetIOCounters(t *testing.T) {
	fmt.Println(NetIOCounters())
}

func TestNetConnections(t *testing.T) {
	fmt.Println(NetConnections("tcp"))
}

func TestGetOS(t *testing.T) {
	fmt.Println(GetOS())
}

func TestGetArch(t *testing.T) {
	fmt.Println(GetArch())
}

func TestGetArchBit(t *testing.T) {
	fmt.Println(GetArchBit())
}

func TestGetCpuCores(t *testing.T) {
	fmt.Println(GetCpuCores())
}

func TestSetGoMaxProc(t *testing.T) {
	fmt.Println(SetGoMaxProc(0))
}
