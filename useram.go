package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	// run as taskset 0x1 go run useram.go 50 [MB]  600
	ctr := 1024
	dur := 120
	if len(os.Args) > 1 { ctr, _ = strconv.Atoi(os.Args[1])}
	if len(os.Args) > 2 { dur, _ = strconv.Atoi(os.Args[2])}
	fmt.Printf("Starting for  %d MB memory utilization for %d seconds [taskset 0x1 go run %s.go %d %d]\n\n",ctr,dur,os.Args[0],ctr,dur)
	blob := make([]byte, 1024 * 1024 * ctr, 1024 * 1024 * ctr)
	fmt.Printf("Initializing %d MB\n",len(blob))
	for i := range(blob) {blob[i] = byte(i % 240) ; if i % 1024 == 0 {fmt.Printf("%10d / %10d \r",i/1024,ctr*1024)} }
	fmt.Printf("%10d / %10d \r",ctr*1024,ctr*1024)
	fmt.Printf("\n\nAllocated %d MB\n",len(blob))
	for i:= 1; i <= dur; i++ {
		time.Sleep( 1 *  time.Second)
		fmt.Printf("loop %d / %d seconds at %d MB RAM utilization\r",i,dur,ctr)
	}
	fmt.Println("\nDone\n")
}

