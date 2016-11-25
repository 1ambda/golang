package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"golang.org/x/net/context"

	. "github.com/1ambda/golang/go-kit-tutorial/stringsvc3/service"
)

func main() {

	var (
		listen = flag.String("listen", ":8080", "HTTP listen address")
		proxy  = flag.String("proxy", "", "Optional comma-separated list of URLS to proxy requests")
	)
	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.NewContext(logger).With("listen", *listen).With("caller", log.DefaultCaller)
	ctx := context.Background()

	var svc StringService
	svc = StringServiceImpl{}
	svc = CreateUppercaseProxyMiddleware(*proxy, ctx, logger)(svc)
	svc = CreateLoggingMiddleware(logger)(svc)
	svc = CreatePrebuiltInstrumentMiddleware()(svc)

	uppercaseHandler := CreateUppercaseHandler(ctx, svc)
	countHandler := CreateCountHandler(ctx, svc)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	http.Handle("/metrics", stdprometheus.Handler())
	logger.Log("msg", "HTTP", "addr", *listen)
        logger.Log("err", http.ListenAndServe(*listen, nil))
}
