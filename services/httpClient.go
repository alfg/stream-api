package services

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type HTTPClient struct {
	*http.Client
}

func (c *HTTPClient) loadResponse(url string, i interface{}, format string) error {

	fmt.Println("querying..." + url)
	rsp, e := c.Get(url)
	if e != nil {
		return e
	}

	defer rsp.Body.Close()

	b, e := ioutil.ReadAll(rsp.Body)
	if e != nil {
		return e
	}
	if rsp.Status[0] != '2' {
		return fmt.Errorf("expected status 2xx, got %s: %s", rsp.Status, string(b))
	}

	if format == "json" {
		return json.Unmarshal(b, &i)
	} else if format == "xml" {
		return xml.Unmarshal(b, &i)
	}
	return fmt.Errorf("expected format, got %s", format)
}

func NewClient() (*HTTPClient, error) {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	return &HTTPClient{Client: &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}}, nil
}
