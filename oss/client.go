package oss

import (
	"bytes"
	"fmt"

	"encoding/xml"
	"io"
	"io/ioutil"
	"log"

	"net/http"

	"strconv"

	"time"
)

const (
	//DefaultContentType is default Content-Type
	DefaultContentType = "application/octet-stream"

	//HeaderOSSPrefix is prefix of OSS Header
	HeaderOSSPrefix = "x-oss-"

	//HeaderOSSACL is x-oss-acl Header
	HeaderOSSACL = "x-oss-acl"
)

// Client represents client for OSS service
type Client struct {
	AccessKeyId     string
	AccessKeySecret string
	Region          Region
	Internal        bool
	httpClient      *http.Client
	debug           bool
}

// NewOSSClient create a new instance of OSS client
func NewOSSClient(region Region, internal bool, accessKeyId string, accessKeySecret string) *Client {
	return &Client{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Region:          region,
		Internal:        internal,
		httpClient:      &http.Client{},
		debug:           false,
	}
}

// SetAccessKeyId sets new AccessKeyId
func (client *Client) SetAccessKeyId(id string) {
	client.AccessKeyId = id
}

// SetAccessKeySecret sets new AccessKeySecret
func (client *Client) SetAccessKeySecret(secret string) {
	client.AccessKeySecret = secret
}

// SetDebug sets debug mode to log the request/response message
func (client *Client) SetDebug(debug bool) {
	client.debug = debug
}

func getOSSErrorFromString(str string) error {
	return &Error{
		ErrorResponse: ErrorResponse{
			Code:    "OSSClientFailure",
			Message: str,
		},
		StatusCode: -1,
	}
}

func getOSSError(err error) error {
	return getOSSErrorFromString(err.Error())
}

// Invoke sends the raw HTTP request for OSS service
func (client *Client) Invoke(method, url string, body io.Reader, headers http.Header) (httpResp *http.Response, err error) {
	host := client.Region.GetEndpoint(client.Internal)
	req, err := http.NewRequest(method, host+url, body)
	if err != nil {
		log.Printf("Error Creating Request: %v\n", err)
	}

	if headers == nil {
		headers = make(http.Header)
	}
	headers.Set("Date", getGMTime())
	req.Header = headers

	if client.AccessKeySecret != "" {
		headers.Add("Authorization", client.createAuthorizationHeader(req))
	} else {
		headers.Add("Authorization", client.AccessKeyId)
	}

	if contentLenStr := headers.Get("Content-Length"); contentLenStr != "" {
		if cLen, err := strconv.ParseInt(contentLenStr, 10, 64); err != nil {
			req.ContentLength = -1
		} else {
			req.ContentLength = cLen
		}
	}
	if client.debug {
		log.Printf("Request: %++v", req)
	}

	t0 := time.Now()
	httpResp, err = client.httpClient.Do(req)
	t1 := time.Now()

	if err != nil {
		return nil, getOSSError(err)
	}

	//defer httpResp.Body.Close()
	statusCode := httpResp.StatusCode
	log.Printf("Invoke %s %s %d (%v)", method, req.URL, statusCode, t1.Sub(t0))
	if client.debug {
		body, err := ioutil.ReadAll(httpResp.Body)
		if err != nil {
			return nil, getOSSError(err)
		}
		log.Printf(string(body))
		httpResp.Body = ioutil.NopCloser(bytes.NewReader(body))
	}

	if statusCode >= 400 && statusCode <= 599 {
		ossError := &Error{}
		decoder := xml.NewDecoder(httpResp.Body)
		err := decoder.Decode(ossError)
		if err != nil {
			return nil, getOSSError(err)
		}
		ossError.StatusCode = statusCode
		return nil, ossError
	}
	return httpResp, nil
}

func (client *Client) decodeResponse(httpResp *http.Response, response interface{}) error {
	if response != nil {
		decoder := xml.NewDecoder(httpResp.Body)
		err := decoder.Decode(response)
		if err != nil {
			return getOSSError(err)
		}
		if client.debug {
			log.Printf("Response: %++v", response)
		}
	}
	return nil
}

func copyHeader(header http.Header) (newHeader http.Header) {
	newHeader = make(http.Header)
	for k, v := range header {
		newSlice := make([]string, len(v))
		copy(newSlice, v)
		newHeader[k] = newSlice
	}
	return
}

func getGMTime() string {
	return time.Now().UTC().Format(http.TimeFormat)
}

func getExpires() string {
	return fmt.Sprintf("%d", time.Now().Unix()+60)
}
