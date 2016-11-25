package main

import (
        "fmt"
        "log"
        "os"

        httptransport "github.com/go-kit/kit/transport/http"

        "context"
        "github.com/1ambda/golang/go-kit-tutorial/stringsvc/endpoint"
        "github.com/1ambda/golang/go-kit-tutorial/stringsvc/service"
        "github.com/1ambda/golang/go-kit-tutorial/stringsvc/transport"
        "net/http"
)

func main() {
        port := os.Getenv("PORT")
        if port == "" {
                port = "9090"
        }
        addr := fmt.Sprintf(":%s", port)

        ctx := context.Background()
        svc := service.StringServiceImpl{}

        uppercaseHandler := httptransport.NewServer(
                ctx,
                endpoint.MakeUppercaseEndpoint(svc),
                transport.DecodeUppercaseRequest,
                transport.EncodeResponse,
        )

        countHandler := httptransport.NewServer(
                ctx,
                endpoint.MakeCountEndpoint(svc),
                transport.DecodeCountRequest,
                transport.EncodeResponse,
        )

        http.Handle("/uppercase", uppercaseHandler)
        http.Handle("/count", countHandler)
        log.Printf("Starting :%s\n", port)
        log.Fatal(http.ListenAndServe(addr, nil))
}
