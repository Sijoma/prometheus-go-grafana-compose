package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	appVersion string
	version    = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "my_own_version",
		Help: "Version information about this binary",
		ConstLabels: map[string]string{
			"version": appVersion,
		},
	})

	httpRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "my_own_http_requests_total",
		Help: "Count of all HTTP requests",
	}, []string{"code", "method"})

	httpRequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "my_own_http_request_duration_seconds",
		Help: "Duration of all HTTP requests",
	}, []string{"code", "handler", "method"})
)

func main() {
	version.Set(1)
	bind := ""
	flagset := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flagset.StringVar(&bind, "bind", ":9200", "The socket to bind to.")
	flagset.Parse(os.Args[1:])

	r := prometheus.NewRegistry()
	r.MustRegister(httpRequestsTotal)
	r.MustRegister(httpRequestDuration)
	r.MustRegister(version)

	foundHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Println("Scraped Success")
		w.Write([]byte("Hello from example application."))
	})
	notfoundHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Scraped not found")
		w.WriteHeader(http.StatusNotFound)
	})

	foundChain := promhttp.InstrumentHandlerDuration(
		httpRequestDuration.MustCurryWith(prometheus.Labels{"handler": "found"}),
		promhttp.InstrumentHandlerCounter(httpRequestsTotal, foundHandler),
	)

	http.Handle("/", foundChain)
	http.Handle("/err", promhttp.InstrumentHandlerCounter(httpRequestsTotal, notfoundHandler))

	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(bind, nil))
}
