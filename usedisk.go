package main

import (
        "fmt"
	"time"
        "github.com/shirou/gopsutil/disk"
)

func main() {

        f, _ := disk.Partitions(false)
        fmt.Printf("DISK Partitions: %+v\n\n", f)
        fmt.Println(f)
        partitions := []string{}
        for _, part := range f {
                g, _ := disk.Usage(part.Mountpoint)
                partitions = append(partitions,part.Mountpoint)
                fmt.Printf("DISK Usage: %+v\n", g)
        }
        //partitionsIO, err := disk.IOCounters(partitions...)
        partitionsIO, _  := disk.IOCounters()
	time.Sleep(100 * time.Millisecond)
        partitionsIO1, _ := disk.IOCounters()
        fmt.Printf("DISK IOCounters: %+v\n\n", partitionsIO)
        fmt.Printf("DISK IOCounters: %+v\n\n", partitionsIO1)
}


