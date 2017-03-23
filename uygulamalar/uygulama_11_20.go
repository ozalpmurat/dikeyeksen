// uygulama_11_20
package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

func main() {
	fmt.Println(mem.VirtualMemory())
	fmt.Println(cpu.Info())
	fmt.Println(disk.GetDiskSerialNumber("sda"))
	fmt.Println(disk.Usage("/"))
	fmt.Println(host.PlatformInformation())
	fmt.Println(host.Info())
	fmt.Println(host.KernelVersion())
	fmt.Println(load.Avg())
	fmt.Println(net.Interfaces())
	fmt.Println(net.IOCounters(true))
	// Değerler genellikle yapı türündedir.
	// istenirse tek bir alana erişilebilir:
	fmt.Println("\nBellek bilgileri:")
	bellek, _ := mem.VirtualMemory()
	fmt.Printf("Toplam: %dMB, Boş:%dMB\n",
		bellek.Total/(1024*1024), bellek.Free/(1024*1024))
}
