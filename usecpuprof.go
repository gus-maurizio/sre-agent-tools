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
	// run as taskset 0x1 go run usecpuprof.go [number of cycles|100] [ list of duples [cpu% milliseconds]...| 100 500 10 500] 
	// Calibrate function usecpu to understand # of loops per millisecond at 100% use for 500ms and 10% for 500ms
	// It is important to run this program pegged to a specific CPU (any) if possible via taskset 0x1
	// We need to seed the initial estimate, use 300K loops to start 
	calcount := int64(300 * 1000)
	calloops := 20
	calavg   := int64(0)
	for i := 1; i <= calloops; i++ {
		t0 := time.Now().UnixNano()
		cal := usecpu(calcount)
		t1 := time.Now().UnixNano()
		loopspermillisec := int64(1000000) * calcount / int64(t1 - t0)
		// loopback for convergence
		calcount = loopspermillisec
		calavg  += calcount 
		fmt.Printf("---> loop %02d calcount %d [cal %d]\n",i,calcount,cal)
	}
	loopspermillisec := calavg / int64(calloops)
	fmt.Printf("Calibration complete. Loops per millisecond %d \n\n",loopspermillisec)
	// get the parameters and initialize with defaults
	cpuloops	:= 100
	cpuprofile	:= [][2]int{	{100, 500},
								{10,  450},
	} 
	qargs 		:= len(os.Args)
	if qargs > 1 { cpuloops, _ = strconv.Atoi(os.Args[1])}
	if qargs > 2 { 
		// are arguments multiple of two?
		qargs -= 2		// remove program name and cpuloops
		if qargs > 0 && qargs % 2 == 0 {
			duples 		:= qargs / 2
			cpuprofile 	=  nil		// clear slice but do not deallocate
			for i := 0; i < duples; i++ {
				cpupercent, _ := strconv.Atoi(os.Args[2 * (i+1) + 0])
				cpumsecond, _ := strconv.Atoi(os.Args[2 * (i+1) + 1])
				fmt.Printf("---> profile %2d/%2d [%03d%% %03dms]\n",i, duples,cpupercent, cpumsecond)
				cpuprofile 	= append(cpuprofile,[2]int{ cpupercent, cpumsecond})
			}
		}
		cpuloops, _ = strconv.Atoi(os.Args[1])
	}
	// compute the total number of milliseconds per loop and the average cpu usage
	totms 	:= 0
	avgcpu 	:= 0
	for _, prof := range(cpuprofile) {
		avgcpu += prof[0] * prof[1]
		totms  += prof[1]
	}
	avgcpu  /= totms
	res		:= 0
	fmt.Printf("Executing %d %dms (%d seconds) loops at %d%% average cpu with profile: %v \n\n",cpuloops, totms, cpuloops*totms/1000, avgcpu, cpuprofile)
	for i := 1; i <= cpuloops; i++ {
		for j, prof := range(cpuprofile) {
			// we will do 10 msec loops to ensure some smooth distribution
			for k := 1; k <= prof[1] / 10; k++ {
				// inside each loop
				fmt.Printf("loop %3d / %3d seconds profile %02d %03d/%03d [%03d%% CPU, %3dms]\r",i,cpuloops,j, k, prof[1]/10, prof[0],prof[1])
				loopcount   := int64(prof[0] / 10) * loopspermillisec
				timetosleep := time.Duration((100 - prof[0]) / 10) * time.Millisecond 
				res += usecpu(loopcount)
				time.Sleep(timetosleep)				
			}
		}
	}
	fmt.Println("\nDone (lucky number is %d)\n", res)
/*
	// Executes a number of cycles (default is 100 cycles)
	// Each cycle
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
*/
}

