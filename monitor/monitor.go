package monitor

import (
	"fmt"
	"runtime"
)

func PrintMemUsage() {

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	alloc := fmt.Sprintf("Alloc=%v MiB", bToMb(m.Alloc))
	total := fmt.Sprintf("TotalAlloc=%v MiB", bToMb(m.TotalAlloc))
	sys := fmt.Sprintf("Sys=%v MiB", bToMb(m.Sys))
	gc := fmt.Sprintf("NumGC =%v", m.NumGC)

	fmt.Println(alloc, total, sys, gc)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
