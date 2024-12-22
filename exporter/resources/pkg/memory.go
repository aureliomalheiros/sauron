package memory

import (
	"github.com/pbnjay/memory"
)

func BytesToMegabytes(bytes uint64) float64 {
	return float64(bytes) / (1024 * 1024)
}

func GetMemoryTotal() float64 {
	return BytesToMegabytes(memory.TotalMemory())
}

func GetMemoryFree() float64 {
	return BytesToMegabytes(memory.FreeMemory())
}
