package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cloudfoundry/dropsonde"
	"github.com/cloudfoundry/dropsonde/metrics"
)

func main() {
	var (
		endpoint     string
		metricPrefix string
		metricName   string
		metricValue  float64
		metricUnit   string
	)

	flag.StringVar(&endpoint, "metron-endpoint", "127.0.0.1:3457", "Metron endpoint")
	flag.StringVar(&metricPrefix, "prefix", "", "Metric prefix")
	flag.StringVar(&metricName, "name", "", "Metric name")
	flag.Float64Var(&metricValue, "value", 0.0, "Metric value")
	flag.StringVar(&metricUnit, "unit", "", "Metric unit")

	flag.Parse()

	if err := dropsonde.Initialize(endpoint, metricPrefix); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialise dropsonde: %s", err.Error())
		os.Exit(1)
	}

	if err := metrics.SendValue(metricName, float64(metricValue), metricUnit); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to send metric: %s", err.Error())
		os.Exit(1)
	}
}
