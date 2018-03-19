package livebarn

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const curlFormat = `curl '%s' -XPOST \
-H 'X-LiveBarn-URL: %s' \
-H 'X-LiveBarn-Data: %s' \
-H 'X-LiveBarn-Timestamp: %s' \
-H 'X-LiveBarn-Signature: %s' \
-H 'X-LiveBarn-Application-Id: %s' \
-H 'Content-Type: application/json' \
-H 'Accept: application/json'`

// Request is a struct for containing all parameters for a LiveBarn request
type Request struct {
	URL       string
	Data      string
	Timestamp string
}

// NullRequest is a Request struct that is returned as a placeholder when an
// error occurs
var NullRequest = &Request{}

// NullHTTPRequest is an http.Request struct that is returned as a placeholder
// when an error occurs
var NullHTTPRequest = &http.Request{}

// NewRequest takes a URL string and Data struct and prepares the request for
// sending to the server
func NewRequest(url string, data interface{}) (*Request, error) {
	dataJSON, err := toJSON(data)
	if err != nil {
		return NullRequest, err
	}
	return &Request{
		URL:       url,
		Data:      dataJSON,
		Timestamp: makeTimestamp(),
	}, nil
}

// HTTPRequest returns an *http.Request populated with the required headers and
// properly signed
func (r *Request) HTTPRequest() (*http.Request, error) {
	req, err := http.NewRequest("POST", r.URL, nil)
	if err != nil {
		return NullHTTPRequest, err
	}

	req.Header.Set("X-Livebarn-Data", r.Data)
	req.Header.Set("X-Livebarn-Url", r.URL)
	req.Header.Set("X-Livebarn-Signature", r.Signature())
	req.Header.Set("X-Livebarn-Timestamp", r.Timestamp)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Livebarn-Application-Id", AppKey)

	return req, nil
}

// ToCurl returns a string containing a properly formatted cURL command
func (r *Request) ToCurl() string {
	return fmt.Sprintf(curlFormat, r.URL, r.URL, r.Data, r.Timestamp, r.Signature(), AppKey)
}

// JSON generates the JSON payload for `X-LiveBarn-Data`
func toJSON(d interface{}) (string, error) {
	data, err := json.Marshal(d)
	return string(data), err
}

// Signature generates the `X-LiveBarn-Signature` field by concatenating the
// header values for URL, data, and timestamp (in that order) and generating
// an HMAC SHA256 signature using the API key
func (r *Request) Signature() string {
	return computeHmac256Signature(r.URL + string(r.Data) + r.Timestamp)
}

func makeTimestamp() string {
	return strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
}

func computeHmac256Signature(message string) string {
	h := hmac.New(sha256.New, []byte(APIKey))
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
