package main

import (
	"fmt"
	"log"
	localnet "net"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

func getCPUInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cup info failed, err:%v", err)
		return
	}

	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}

	for {
		percent, _ := cpu.Percent(time.Second, false)
		fmt.Printf("cpu percent:%v\n", percent)
	}
}

func getCPULoad() {
	info, _ := load.Avg()
	fmt.Printf("%v\n", info)
}

func getMemoryInfo() {
	memInfo, _ := mem.VirtualMemory()
	fmt.Printf("mem info:%v\n", memInfo)
}

func getHostInfo() {
	hInfo, _ := host.Info()
	fmt.Printf("host info:%v uptime:%v boottime:%v\n", hInfo, hInfo.Uptime, hInfo.BootTime)
}

func getDiskInfo() {
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get partitions failed, err:%v\n", err)
		return
	}

	for _, part := range parts {
		fmt.Printf("part:%v\n", part.String())
		diskInfo, _ := disk.Usage(part.Mountpoint)
		fmt.Printf("disk info:used:%v free:%v\n", diskInfo.UsedPercent, diskInfo.Free)
	}

	ioStat, _ := disk.IOCounters()
	for k, v := range ioStat {
		fmt.Printf("%v:%v\n", k, v)
	}

}

func getNetInfo() {
	info, _ := net.IOCounters(true)
	for index, v := range info {
		fmt.Printf("%v:%v send:%v recv:%v\n", index, v, v.BytesSent, v.BytesRecv)
	}
}

func getLocalIP() (ip string, err error) {
	address, err := localnet.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, addr := range address {
		ipAddr, ok := addr.(*localnet.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}

		return ipAddr.IP.String(), nil
	}

	return
}

func getOutBoundIP() string {
	conn, err := localnet.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*localnet.UDPAddr)
	fmt.Println(localAddr.String())

	return localAddr.IP.String()
}

func main() {
	// getCPUInfo()
	// getCPULoad()
	// getMemoryInfo()
	// getHostInfo()
	// getDiskInfo()
	// getNetInfo()
	// fmt.Println(getLocalIP())
	fmt.Println(getOutBoundIP())
}
