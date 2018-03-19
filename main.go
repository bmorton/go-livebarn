package livebarn

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

// AppKey is the application key from https://livebarn.com/js/base.6.0.1.141.js
const AppKey = "a614b6c8-8f6f-4b14-a81d-9d0a14638dc8"

// APIKey is the API key from https://livebarn.com/js/base.6.0.1.141.js
const APIKey = "b991e51e-ddab-4193-8652-70e622478c3a"

// Client is the primary interface to make calls to the LiveBarn API
type Client struct {
	Token     string
	UUID      string
	DebugMode bool
}

// New creates a new Client with the specified credentials
func New(token string, uuid string) *Client {
	return &Client{
		Token:     token,
		UUID:      uuid,
		DebugMode: false,
	}
}

// NullHTTPResponse is an empty placeholder to be used when an error occurs
var NullHTTPResponse = &http.Response{}

func (c *Client) do(url string, data interface{}) (*http.Response, error) {
	request, err := NewRequest(url, data)
	if err != nil {
		return NullHTTPResponse, err
	}

	req, err := request.HTTPRequest()
	if err != nil {
		return NullHTTPResponse, err
	}

	if c.DebugMode {
		debug(httputil.DumpRequestOut(req, true))
	}

	resp, err := http.DefaultClient.Do(req)

	if c.DebugMode {
		debug(httputil.DumpResponse(resp, true))
	}

	return resp, err
}

func debug(data []byte, err error) {
	if err == nil {
		fmt.Printf("%s\n\n", data)
	} else {
		log.Fatalf("%s\n\n", err)
	}
}

// DateRange represents a start and end time with proper formatting for the
// LiveBarn API
type DateRange struct {
	Start time.Time
	End   time.Time
}

const dateFormat = "2006-01-02T15:04"

// StartFormatted returns a properly formatted start time for the LiveBarn API
func (d *DateRange) StartFormatted() string {
	return d.Start.Format(dateFormat)
}

// EndFormatted returns a properly formatted end time for the LiveBarn API
func (d *DateRange) EndFormatted() string {
	return d.End.Format(dateFormat)
}
