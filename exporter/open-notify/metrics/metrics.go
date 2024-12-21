
package metrics


import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Exporter struct {
	PeopleInSpaceGauge 		prometheus.Gauge
	PeopleInScapceDetails 	*prometheus.GaugeVec
}

func NewExporter() *Exporter {
	exporter := &Exporter{
		PeopleInSpaceGauge: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "people_in_space",
				Help: "The current number of people in space",
			},
		),
		PeopleInScapceDetails: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "people_in_space_details",
				Help: "Details about people in space",
			},
			[]string{"craft", "name"},
		),
	}

	prometheus.MustRegister(exporter.PeopleInSpaceGauge)
	prometheus.MustRegister(exporter.PeopleInScapceDetails)

	return exporter
}

func StartMetricsServer(port string) {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Printf("Prometheus server running in port %s\n", port)
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			panic(err)
		}
	}()
}
