package main

import (
	"fmt"
	"time"
	"github.com/aureliomalheiros/sauron/open-notity-fy/configs"
	"github.com/aureliomalheiros/sauron/open-notity-fy/pkg/people"
	"github.com/aureliomalheiros/sauron/open-notity-fy/metrics"
)

func main() {
	apiURL := configs.Config.APIURL
	port   := configs.Config.HTTPPort

	exporter := metrics.NewExporter()

	metrics.StartMetricsServer(port)
	fmt.Printf("Prometheus server running in port http://localhost/%s/metrics\n", port)

	for {
		peopleInSpace, err := people.FetchData(apiURL)
		if err != nil {
			fmt.Println("Error fetching data: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		

		exporter.PeopleInSpaceGauge.Set(float64(peopleInSpace.Number))
		exporter.PeopleInScapceDetails.Reset()
		for _, p := range peopleInSpace.People {
			exporter.PeopleInScapceDetails.WithLabelValues(p.Craft, p.Name).Set(1)
		}
		fmt.Printf("Update people_in_space to %d\n", peopleInSpace.Number)
		time.Sleep(60 * time.Second)
	}
}

