package cli

import (
	"context"
	"net"
	"net/http"
	"runtime"
	"time"
)

const _defaultConcurrentDialsPerCPU = 5 << 1

var (
	_defaultTimeOut   = 30 * time.Second
	_defaultKeepAlive = 30 * time.Second
)

type dialRateLimiter struct {
	dialer *net.Dialer
	sem    chan struct{}
}

func DownloadClient(maxIdleConnsPerHost int) *http.Client {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	dialer := rateLimiter(&net.Dialer{
		Timeout:   _defaultTimeOut,
		KeepAlive: _defaultKeepAlive,
	})

	transport.DialContext = dialer.DialContext
	transport.MaxIdleConns = 0 // not limit
	transport.MaxIdleConnsPerHost = maxIdleConnsPerHost
	transport.DisableCompression = true
	return &http.Client{
		Transport: transport,
	}
}

func DefaultClient() *http.Client {
	return http.DefaultClient
}

func rateLimiter(dialer *net.Dialer) *dialRateLimiter {
	return &dialRateLimiter{
		dialer: dialer,
		sem:    make(chan struct{}, _defaultConcurrentDialsPerCPU*runtime.NumCPU()),
	}
}

func (d *dialRateLimiter) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	d.sem <- struct{}{}
	defer func() { <-d.sem }()
	return d.dialer.DialContext(ctx, network, address)
}
