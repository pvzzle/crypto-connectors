package binance

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"
)

type connector struct {
	httpClient  *http.Client
	credentials credentials
}

type credentials struct {
	public string
	secret string
}

func NewConnector() *connector {
	return &connector{
		httpClient: &http.Client{
			Timeout: time.Second,
		},
	}
}

func (c *connector) hasCredentials() bool {
	return c.credentials.public != "" && c.credentials.secret != ""
}

func (c *connector) setCredentials(public, secret string) {
	c.credentials.public = public
	c.credentials.secret = secret
}

func (c *connector) privateFetch(method, path, queryString string, body []byte) (*http.Response, error) {
	if !c.hasCredentials() {
		return nil, errors.New("set up credentials for private API access")
	}

	r := &http.Request{
		Method: method,
		URL: &url.URL{
			Path:     path,
			RawQuery: queryString,
		},
		Body:   io.NopCloser(bytes.NewReader(body)),
		Header: http.Header{},
	}

	signature := generateSignature(queryString, string(body), c.credentials.secret)
	r.Header.Set("X-MBX-APIKEY", signature)

	return c.makeRequest(r)
}

func (c *connector) publicFetch(method, path, queryString string) (*http.Response, error) {
	r := &http.Request{
		Method: method,
		URL: &url.URL{
			Path:     path,
			RawQuery: queryString,
		},
		Header: http.Header{},
	}

	return c.makeRequest(r)
}

func (c *connector) makeRequest(r *http.Request) (*http.Response, error) {
	r.URL.Host = "api.binance.com"
	r.URL.Scheme = "https"

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.httpClient.Do(r)
	if err != nil {
		return nil, errors.New("request failed")
	}

	return res, nil
}
