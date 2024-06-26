package common

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	n "net"
	"os"
	"runtime"
)

// HostInfo 主机信息
func HostInfo() (map[string]any, error) {
	hostInfo, err := host.Info()
	if err != nil {
		return map[string]any{}, err
	}

	var infoMap map[string]any

	err = JsonDecode(JsonEncode(hostInfo), &infoMap)

	if err != nil {
		return map[string]any{}, err
	}

	return infoMap, nil
}

// CpuInfo cpu信息
func CpuInfo() ([]map[string]any, error) {
	cpuInfos, err := cpu.Info()
	if err != nil {
		return []map[string]any{}, err
	}

	var infoArr []map[string]any

	err = JsonDecode(JsonEncode(cpuInfos), &infoArr)

	if err != nil {
		return []map[string]any{}, err
	}

	return infoArr, nil
}

// MemoryInfo 内存信息
func MemoryInfo() (map[string]any, error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return map[string]any{}, err
	}

	var infoMap map[string]any

	err = JsonDecode(JsonEncode(memInfo), &infoMap)

	if err != nil {
		return map[string]any{}, err
	}

	return infoMap, nil
}

// DiskInfo 磁盘信息
func DiskInfo() ([]map[string]any, error) {
	partitions, err := disk.Partitions(true)
	if err != nil {
		return []map[string]any{}, err
	}

	var infoArr []map[string]any

	err = JsonDecode(JsonEncode(partitions), &infoArr)

	if err != nil {
		return []map[string]any{}, err
	}

	return infoArr, nil
}

// NetInterfaces 网络接口信息
func NetInterfaces() ([]map[string]any, error) {
	netStats, err := net.Interfaces()
	if err != nil {
		return []map[string]any{}, err
	}

	var infoArr []map[string]any

	err = JsonDecode(JsonEncode(netStats), &infoArr)

	if err != nil {
		return []map[string]any{}, err
	}

	return infoArr, nil
}

// NetIOCounters 网络IO统计信息
func NetIOCounters() ([]map[string]any, error) {
	stats, err := net.IOCounters(true)

	if err != nil {
		return []map[string]any{}, err
	}

	var infoArr []map[string]any

	err = JsonDecode(JsonEncode(stats), &infoArr)

	if err != nil {
		return []map[string]any{}, err
	}

	return infoArr, nil
}

// NetConnections 网络连接信息, kind可以为""、all、inet、inet4、inet6、tcp、tcp4、tcp6、udp、udp4、udp6、unix
func NetConnections(kind string) ([]map[string]any, error) {
	connections, err := net.Connections(kind)

	if err != nil {
		return []map[string]any{}, err
	}

	var infoArr []map[string]any

	err = JsonDecode(JsonEncode(connections), &infoArr)

	if err != nil {
		return []map[string]any{}, err
	}

	return infoArr, nil
}

// GetHostName 获取主机名
func GetHostName() string {
	name, err := os.Hostname()
	if err != nil {
		name = ""
	}
	return name
}

// GetOS 获取系统
func GetOS() string {
	return runtime.GOOS
}

// GetArch 获取系统架构
func GetArch() string {
	return runtime.GOARCH
}

// GetArchBit 获取架构位数
func GetArchBit() int {
	return 32 << (^uint(0) >> 63)
}

// GetCpuCores 获取cpu数
func GetCpuCores() int {
	return runtime.NumCPU()
}

// SetGoMaxProc 设置最大进程数
func SetGoMaxProc(n int) int {
	return runtime.GOMAXPROCS(n)
}

// GetHostByName 返回主机名对应的 IPv4地址
func GetHostByName(hostname string) string {
	ipAddr, err := n.LookupIP(hostname)
	if err != nil {
		return ""
	}
	for _, ip := range ipAddr {
		if ip4 := ip.To4(); ip4 != nil {
			return ip4.String()
		}
	}
	return ""
}

// GetHostByNameL 返回主机名对应的 IPv4地址列表
func GetHostByNameL(hostname string) []string {
	hosts := make([]string, 0)
	ipAddr, err := n.LookupIP(hostname)
	if err != nil {
		return hosts
	}
	for _, ip := range ipAddr {
		if ip4 := ip.To4(); ip4 != nil {
			hosts = append(hosts, ip4.String())
		}
	}
	return hosts
}

// GetHostByAddr 指定 IP 地址对应的 Internet 主机名
func GetHostByAddr(ip string) string {
	hostnames, err := n.LookupAddr(ip)
	if err != nil {
		return ""
	}
	host := ""
	for _, hostname := range hostnames {
		ipAddr, err := n.LookupIP(hostname)
		if err != nil {
			continue
		}
		for _, ipData := range ipAddr {
			if ipData.String() == ip {
				host = hostname
				break
			}
		}
	}
	return host
}
