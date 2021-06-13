package clients

import "github.com/prometheus/client_golang/prometheus"

type PrometheusClient interface {
	RecordTotalNumberOfOrders()
	RegisterMetrics()
}
type prometheusClient struct {
}

var totalNumberOfOrders = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "orders_total",
		Help: "To get number of Orders",
	},
)

func NewPrometheusClient() PrometheusClient {
	return &prometheusClient{}
}

func (p prometheusClient) RecordTotalNumberOfOrders() {
	totalNumberOfOrders.Inc()
}

func (p prometheusClient) RegisterMetrics() {
	prometheus.Register(totalNumberOfOrders)
}
