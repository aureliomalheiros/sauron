package metrics

import (
	"log"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Exporter struct {
	GetMemoryFree 	prometheus.Gauge
	GetMemoryTotal 	prometheus.Gauge
}

func NewExporter() *Exporter {
	exporter := &Exporter{
		GetMemoryFree: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "memory_free", 
				Help: "The amount of free memory in the system",
			}),
		GetMemoryTotal: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "memory_total",
				Help: "The total amount of memory in the system",
			}),
	}

	prometheus.MustRegister(exporter.GetMemoryFree)
	prometheus.MustRegister(exporter.GetMemoryTotal)

	return exporter

}

func StartMetricsServer(port string) {
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
