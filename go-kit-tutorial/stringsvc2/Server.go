package main

import (
	"fmt"
	glog "log"
	"os"

	httptransport "github.com/go-kit/kit/transport/http"

	"context"
	. "github.com/1ambda/golang/go-kit-tutorial/stringsvc2/endpoint"
	. "github.com/1ambda/golang/go-kit-tutorial/stringsvc2/service"
	. "github.com/1ambda/golang/go-kit-tutorial/stringsvc2/transport"
	"github.com/go-kit/kit/log"
	"net/http"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	addr := fmt.Sprintf(":%s", port)
	metricGroup := "my_group"
	metricSystem := "string_servie"

	logger := log.NewLogfmtLogger(os.Stderr)
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: metricGroup,
		Subsystem: metricSystem,
		Name:      "request_count",
		Help:      "Number of requests received",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: metricGroup,
		Subsystem: metricSystem,
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: metricGroup,
		Subsystem: metricSystem,
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	ctx := context.Background()
	var svc StringService
	svc = StringServiceImpl{}
	svc = LoggingMiddleware{logger, svc}
	svc = InstrumentMiddleware{requestCount, requestLatency, countResult, svc}

	uppercaseHandler := httptransport.NewServer(
		ctx,
		MakeUppercaseEndpoint(svc),
		DecodeUppercaseRequest,
		EncodeResponse,
	)

	countHandler := httptransport.NewServer(
		ctx,
		MakeCountEndpoint(svc),
		DecodeCountRequest,
		EncodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	http.Handle("/metrics", stdprometheus.Handler())
	glog.Printf("Starting :%s\n", port)
	glog.Fatal(http.ListenAndServe(addr, nil))
}
