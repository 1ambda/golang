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
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	addr := fmt.Sprintf(":%s", port)

        logger := log.NewLogfmtLogger(os.Stderr)
	ctx := context.Background()
	var svc StringService
        svc = StringServiceImpl{}
        svc = LoggingMiddleware{logger, svc}

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
	glog.Printf("Starting :%s\n", port)
	glog.Fatal(http.ListenAndServe(addr, nil))
}
