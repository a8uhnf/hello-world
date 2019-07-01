package pkg

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

func RegisterMetrics(name string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    name,
		Help:    fmt.Sprintf("this histogram is for %v api", name),
		Buckets: []float64{.002, .005, .01, .025, .05, .075, 0.1, 0.15, .25, 0.5, 0.75, 1, 1.5, 2, 2.5, 5},
	}, []string{"method", "endpoint"})
}
