package service

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	jujuratelimit "github.com/juju/ratelimit"
	"github.com/pkg/errors"
	"github.com/sony/gobreaker"
	"golang.org/x/net/context"
	httptransport "github.com/go-kit/kit/transport/http"
)

const loggingTag = "proxy_to"

type proxyMiddleware struct {
	ctx  context.Context
	StringService
	e    endpoint.Endpoint
}

func (mw proxyMiddleware) Uppercase(s string) (string, error) {
	res, err := mw.e(mw.ctx, UppercaseRequest{s})
	if err != nil {
		return "", errors.Wrap(err, "Failed to proxy UppercaseRequest")
	}

	resp := res.(UppercaseResponse)
	if resp.Err != "" {
		return resp.V, errors.New(resp.Err)
	}
	return resp.V, nil
}

func (mw proxyMiddleware) Count(s string) int {
	return mw.StringService.Count(s)
}

func CreateUppercaseProxyMiddleware(
	instances string,
	ctx context.Context,
	logger log.Logger) StringMiddleware {

	if instances == "" {
		logger.Log(loggingTag, "none")
		return func(next StringService) StringService { return next }
	}

	var (
		maxRequest = 100
		maxRetry   = 3
		maxTime    = 250 * time.Millisecond
	)

	var (
		proxyList  = parseProxyList(instances)
		subscriber sd.FixedSubscriber
	)
	logger.Log("proxy_to", fmt.Sprint(proxyList))
	for _, proxy := range proxyList {
		var e endpoint.Endpoint
		e = createUppercaseProxy(proxy)
		e = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(e)
		e = ratelimit.NewTokenBucketLimiter(jujuratelimit.NewBucketWithRate(float64(maxRequest), int64(maxRequest)))(e)
		subscriber = append(subscriber, e)
	}

	balancer := lb.NewRoundRobin(subscriber)
	retry := lb.Retry(maxRetry, maxTime, balancer)

	return func(next StringService) StringService {
		return proxyMiddleware{ctx, next, retry}
	}
}

func createUppercaseProxy(proxyURL string) endpoint.Endpoint {
	if !strings.HasPrefix(proxyURL, "http") {
		proxyURL = "http://" + proxyURL
	}
	u, err := url.Parse(proxyURL)
	if err != nil {
		err = errors.Wrap(err, "Failed to parse proxy URL")
		panic(err)
	}
	if u.Path == "" {
		u.Path = "/uppercase"
	}

	return httptransport.NewClient(
		"GET",
		u,
		EncodeRequest,
		DecodeUppercaseResponse,
	).Endpoint()
}

func parseProxyList(s string) []string {
	ps := strings.Split(s, ",")
	for i := range ps {
		ps[i] = strings.TrimSpace(ps[i])
	}

	return ps
}
