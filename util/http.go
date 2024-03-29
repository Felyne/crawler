package util

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"time"
)

var defaultTransport = http.Transport{
	Proxy: http.ProxyFromEnvironment,
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).DialContext,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
	TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
}

func NewHttpClient(httpProxy string) (*http.Client, error) {
	if httpProxy == "" {
		return &http.Client{Transport: &defaultTransport}, nil
	}
	proxy, err := url.Parse(httpProxy)
	if err != nil {
		return nil, err
	}
	ts := defaultTransport
	ts.Proxy = http.ProxyURL(proxy)
	return &http.Client{Transport: &ts}, nil
}

func NewHttpRequest(method, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	//req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", getUserAgent())
	return req, nil
}

func HttpGet(url, httpProxy string) (*http.Response, error) {
	client, err := NewHttpClient(httpProxy)
	if err != nil {
		return nil, err
	}
	req, err := NewHttpRequest("GET", url)
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}
