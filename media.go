package livebarn

import (
	"encoding/json"
	"io/ioutil"
)

const getMediaURL = "https://livebarn.com/livebarn/api/v1.0.0/media/get?locale=en_ca"
const getMediaDownloadURL = "https://livebarn.com/livebarn/api/v1.0.0/media/download/get?locale=en_ca"

// GetMediaRequest is a struct for containing `X-LiveBarn-Data` parameters for
// the GetMedia request
type GetMediaRequest struct {
	Token      string   `json:"token"`
	UUID       string   `json:"uuid"`
	Surface    *Surface `json:"surface"`
	BeginDate  string   `json:"beginDate"`
	EndDate    string   `json:"endDate"`
	DeviceType string   `json:"deviceType"`
}

// GetMediaResponse represents the payload returned by GetMedia
type GetMediaResponse struct {
	Status int `json:"status"`
	Result []struct {
		Duration  int    `json:"duration"`
		BeginDate string `json:"beginDate"`
		URL       string `json:"url"`
	} `json:"result"`
	Timestamp int64  `json:"timestamp"`
	Date      string `json:"date"`
	Message   string `json:"message"`
}

// NullGetMediaResponse is a placeholder for an empty GetMediaResponse to be
// returned when an error occurs
var NullGetMediaResponse = &GetMediaResponse{}

// GetMediaDownloadRequest is a struct for containing `X-LiveBarn-Data`
// parameters for the GetMediaDownload request
type GetMediaDownloadRequest struct {
	Token    string `json:"token"`
	UUID     string `json:"uuid"`
	MediaURL string `json:"mediaUrl"`
}

// GetMediaDownloadResponse represents the payload returned by GetMediaDownload
type GetMediaDownloadResponse struct {
	Status int `json:"status"`
	Result struct {
		Duration int     `json:"duration"`
		Venue    Venue   `json:"venue"`
		Surface  Surface `json:"surface"`
		URL      string  `json:"url"`
	} `json:"result"`
	Timestamp int64  `json:"timestamp"`
	Date      string `json:"date"`
	Message   string `json:"message"`
}

// NullGetMediaDownloadResponse is a placeholder for an empty
// GetMediaDownloadResponse to be returned when an error occurs
var NullGetMediaDownloadResponse = &GetMediaDownloadResponse{}

// GetMedia returns video URLs for a given surface and date range
func (c *Client) GetMedia(surface *Surface, dateRange *DateRange) (*GetMediaResponse, error) {
	data := &GetMediaRequest{
		Token:      c.Token,
		UUID:       c.UUID,
		Surface:    surface,
		BeginDate:  dateRange.StartFormatted(),
		EndDate:    dateRange.EndFormatted(),
		DeviceType: "hls",
	}

	resp, err := c.do(getMediaURL, data)
	if err != nil {
		return NullGetMediaResponse, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return NullGetMediaResponse, err
	}

	var response GetMediaResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return NullGetMediaResponse, err
	}

	return &response, nil
}

// GetMediaDownload returns downloadable video URLs for a given media URL
func (c *Client) GetMediaDownload(mediaURL string) (*GetMediaDownloadResponse, error) {
	data := &GetMediaDownloadRequest{
		Token:    c.Token,
		UUID:     c.UUID,
		MediaURL: mediaURL,
	}

	resp, err := c.do(getMediaDownloadURL, data)
	if err != nil {
		return NullGetMediaDownloadResponse, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return NullGetMediaDownloadResponse, err
	}

	var response GetMediaDownloadResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return NullGetMediaDownloadResponse, err
	}

	return &response, nil
}
