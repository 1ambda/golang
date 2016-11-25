package main

import (
	"fmt"
	glog "log"
	"os"

	httptransport "github.com/go-kit/kit/transport/http"

	"context"
	. "github.com/1ambda/golang/go-kit-tutorial/stringsvc2/endpoint"
	. "github.com/1ambda/golang/go-kit-tutorial/stringsvc2/middleware"
	. "github.com/1ambda/golang/go-kit-tutorial/stringsvc2/service"
	. "github.com/1ambda/golang/go-kit-tutorial/stringsvc2/transport"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"net/http"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	addr := fmt.Sprintf(":%s", port)

	ctx := context.Background()
	svc := StringServiceImpl{}
	logger := log.NewLogfmtLogger(os.Stderr)

	var uppercase endpoint.Endpoint = MakeUppercaseEndpoint(svc)
	uppercase = LoggingMiddleware(log.NewContext(logger).With("method", "uppercase"))(uppercase)

	uppercaseHandler := httptransport.NewServer(
		ctx,
		uppercase,
		DecodeUppercaseRequest,
		EncodeResponse,
	)

	var count endpoint.Endpoint = MakeCountEndpoint(svc)
	count = LoggingMiddleware(log.NewContext(logger).With("method", "count"))(count)

	countHandler := httptransport.NewServer(
		ctx,
		count,
		DecodeCountRequest,
		EncodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	glog.Printf("Starting :%s\n", port)
	glog.Fatal(http.ListenAndServe(addr, nil))
}
