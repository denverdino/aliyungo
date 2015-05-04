package oss

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func (client *Client) objectOp(method, bucket, object string, headers http.Header, data io.Reader) (httpResp *http.Response, err error) {
	url := "/" + bucket + "/" + object
	httpResp, err = client.Invoke(method, url, data, headers)
	return
}

// PutObject creates object/updates with byte array
func (client *Client) PutObject(bucket, object string, data []byte, contentType string) error {
	headers := make(http.Header)
	body := bytes.NewBuffer(data)
	if contentType == "" {
		contentType = DefaultContentType
	}

	headers.Add("Content-Length", strconv.Itoa(len(data)))
	headers.Add("Content-Type", contentType)

	_, err := client.objectOp("PUT", bucket, object, headers, body)
	return err
}

// PutObjectFromFile creates/updates object with file
func (client *Client) PutObjectFromFile(bucket, object string, file *os.File) error {
	headers := make(http.Header)

	if dotPos := strings.LastIndex(file.Name(), "."); dotPos == -1 {
		headers.Add("Content-Type", DefaultContentType)
	} else {
		if mimeType := mime.TypeByExtension(file.Name()[dotPos:]); mimeType == "" {
			headers.Add("Content-Type", DefaultContentType)
		} else {
			headers.Add("Content-Type", mimeType)
		}
	}
	stats, err := file.Stat()
	if err != nil {
		log.Panicf("Unable to read file %s stats.", file.Name())
		return nil
	}
	headers.Add("Content-Length", strconv.FormatInt(stats.Size(), 10))
	//headers.Add("Expect", "100-Continue") //TODO: what's for?

	log.Printf("Header in file put: %v", headers)
	_, err = client.objectOp("PUT", bucket, object, headers, file)
	return err
}

// GetObject retrieves object content
func (client *Client) GetObject(bucket, object string, headers http.Header) (data []byte, err error) {

	body, err := client.GetObjectReader(bucket, object, headers)

	if err != nil {
		return nil, err
	}
	data, err = ioutil.ReadAll(body)
	return
}

// GetObjectReader retrieves object content as Reader
func (client *Client) GetObjectReader(bucket, object string, headers http.Header) (body io.ReadCloser, err error) {
	response, err := client.GetObjectResponse(bucket, object, headers)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

func (client *Client) GetObjectResponse(bucket, object string, headers http.Header) (httpResp *http.Response, err error) {
	return client.objectOp("GET", bucket, object, headers, nil)
}

// DeleteObject deletes object
func (client *Client) DeleteObject(bucket, object string) error {
	_, err := client.objectOp("DELETE", bucket, object, nil, nil)
	return err
}

// GetObjectMetadata gets object metadata with HEAD request
func (client *Client) GetObjectMetadata(bucket, object string, headers http.Header) (httpResp *http.Response, err error) {
	return client.objectOp("HEAD", bucket, object, headers, nil)
}
