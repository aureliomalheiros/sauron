package main

import (
	"fmt"
	"time"

	"github.com/aureliomalheiros/sauron/resources/metrics"
	memory "github.com/aureliomalheiros/sauron/resources/pkg"
)

func main() {

	exporter := metrics.NewExporter()

	go metrics.StartMetricsServer("7777")
	fmt.Printf("Prometheus server running in port http://localhost:7777/metrics\n")

	for {

		exporter.GetMemoryFree.Set(memory.GetMemoryFree())
		exporter.GetMemoryTotal.Set(memory.GetMemoryTotal())
		fmt.Printf("Update memory_free to %f\n", memory.GetMemoryFree())
		fmt.Printf("Update memory_total to %f\n", memory.GetMemoryTotal())
		time.Sleep(10 * time.Second)
	}

}
