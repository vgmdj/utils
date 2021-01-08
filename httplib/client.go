package httplib

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

const (
	MinRead = 16 * 1024 // 16kb
)

var (
	hc   *Client
	once sync.Once

	DefaultClientConf = &ClientConf{
		Auth: &DefaultAuth{},
	}
)

//ClientConf client config
type ClientConf struct {
	Auth      Authorization
	Timeout   time.Duration
	KeepAlive time.Duration
}

//Client httplib.Client
type Client struct {
	conf      *ClientConf
	httpCli   *http.Client
	dialer    *net.Dialer
	transport http.RoundTripper
	mutex     sync.RWMutex
}

//UniqueClient return the only client
func UniqueClient(conf *ClientConf) *Client {
	once.Do(func() {
		hc = NewClient(conf)
	})

	return hc
}

//NewClient return httplib.Client
func NewClient(conf *ClientConf) *Client {
	if conf == nil {
		conf = DefaultClientConf
	}

	client := &Client{
		conf: conf,
		dialer: &net.Dialer{
			Timeout:   conf.Timeout,
			KeepAlive: conf.KeepAlive,
		},
	}

	client.transport = &http.Transport{
		DialContext:     client.dialer.DialContext,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client.httpCli = &http.Client{
		Transport: client.transport,
	}

	if err := conf.Auth.CheckFormat(); err != nil {
		panic("must use correct http auth params")
	}

	return client
}

//SetTransport set the transport
func (c *Client) SetTransport(transport *http.Transport) {
	c.transport = transport
	c.httpCli.Transport = transport
}

//SetConfig set config
func (c *Client) SetConfig(conf *ClientConf) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if conf.Auth.CheckFormat() != nil {
		c.conf.Auth = conf.Auth
	}

	if conf.Timeout > 0 {
		c.dialer.Timeout = conf.Timeout
		c.conf.Timeout = conf.Timeout
	}

	if conf.KeepAlive > 0 {
		c.dialer.KeepAlive = conf.KeepAlive
		c.conf.KeepAlive = conf.KeepAlive
	}
}

//Raw sends an HTTP request and use client do
func (c *Client) Raw(method, uri string, body []byte, v interface{}, headers map[string]string) (err error) {

	if headers == nil {
		headers = make(map[string]string)
	}

	request, err := c.NewRequest(method, uri, body, headers)
	if err != nil {
		return
	}

	rb, rh, err := c.DoWithData(request)
	if err != nil {
		return
	}

	config := make(map[string]string)
	config[ResponseResultContentType] = headers[ResponseResultContentType]
	config[ResponseResultAllowNull] = headers[ResponseResultAllowNull]
	config[ContentType] = rh.Get(ContentType)

	p, err := newParser(config)
	if err != nil {
		return
	}

	return p.respParser(rb, v)

}

//NewRequest return the client new request
func (c *Client) NewRequest(method, uri string, body []byte, headers map[string]string) (request *http.Request, err error) {
	request, err = http.NewRequest(method, uri, bytes.NewReader(body))
	if err != nil {
		return
	}

	c.conf.Auth.SetAuth(request)

	for k, v := range headers {
		request.Header.Set(k, v)
	}

	return

}

//Do sends an HTTP request and return response data
func (c *Client) Do(request *http.Request) (resp *http.Response, err error) {
	return c.httpCli.Do(request)
}

//DoWithData sends an HTTP request and return response data
func (c *Client) DoWithData(request *http.Request) ([]byte, http.Header, error) {
	resp, err := c.Do(request)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return nil, nil, fmt.Errorf("bad request %d", resp.StatusCode)
	}

	bts, err := readAll(resp.Body, MinRead)
	if err != nil {
		return nil, nil, err
	}

	return bts, resp.Header, nil
}
