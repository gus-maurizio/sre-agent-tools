package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func usecpu(loop int64) int {
	a := 1
        for j := int64(1); j < loop; j++ {
                a = (a * 2 % 50000) + 1
        }
	return a
}

func main() {
	// run as taskset 0x1 go run resource.go 50 600
	ctr := 100
	dur := 120
	if len(os.Args) > 1 { ctr, _ = strconv.Atoi(os.Args[1])}
	if len(os.Args) > 2 { dur, _ = strconv.Atoi(os.Args[2])}
	calcount := int64(250 * 1000 * 1000)
	t0 := time.Now().UnixNano()
	cal := usecpu(calcount)
	t1 := time.Now().UnixNano()
	delta := t1 - t0
	loopspermillisec := int64(1000 * 1000 * calcount) / delta
	fmt.Printf("Starting for at %d%% utilization for %d seconds %d [taskset 0x1 go run %s.go %d %d]\n\n",ctr,dur,cal,os.Args[0],ctr,dur)
	timetosleep := time.Duration(100 - ctr) * time.Millisecond 
	timetoloop  := int64(ctr) * loopspermillisec
	res := 0
	for i:= 1; i <= dur; i++ {
		res += usecpu(timetoloop)
		time.Sleep(timetosleep)
		fmt.Printf("loop %d / %d seconds at %d%% CPU utilization\r",i,dur,ctr)
	}
	fmt.Println("\nDone\n")
}

