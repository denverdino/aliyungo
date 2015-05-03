package oss

import (
	"bytes"
	"io"
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
func (client *Client) GetObject(bucket, object string, headers http.Header) (body io.ReadCloser, err error) {
	if headers == nil {
		headers = make(http.Header)
	}
	response, err := client.objectOp("GET", bucket, object, headers, nil)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// DeleteObject deletes object
func (client *Client) DeleteObject(bucket, object string) error {
	_, err := client.objectOp("DELETE", bucket, object, nil, nil)
	return err
}
