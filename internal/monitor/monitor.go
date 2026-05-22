package monitor

import (
	"fmt"
	"runtime"
	"time"
)

func StartMonitor() {
	go func() {

		var m runtime.MemStats

		for {

			runtime.ReadMemStats(&m)

			fmt.Printf(`
====================================
GOROUTINES: %d
MEMÓRIA ALOCADA: %.2f MB
MEMÓRIA TOTAL: %.2f MB
GC EXECUTADO: %d
CPU CORES: %d
====================================

`,
				runtime.NumGoroutine(),
				float64(m.Alloc)/1024/1024,
				float64(m.Sys)/1024/1024,
				m.NumGC,
				runtime.NumCPU(),
			)

			time.Sleep(2 * time.Second)
		}
	}()
}