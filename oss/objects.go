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

func (client *Client) ObjectOp(method, bucket, object string, headers http.Header, data io.Reader) (httpResp *http.Response, err error) {
	url := "/" + bucket + "/" + object
	httpResp, err = client.Invoke(method, url, data, headers)
	return
}

func (client *Client) PutObject(bucket, object string, data []byte, contentType string) error {
	headers := make(http.Header)
	body := bytes.NewBuffer(data)
	if contentType == "" {
		contentType = DEFAULT_CONTENT_TYPE
	}

	headers.Add("Content-Length", strconv.Itoa(len(data)))
	headers.Add("Content-Type", contentType)

	_, err := client.ObjectOp("PUT", bucket, object, headers, body)
	return err
}

func (client *Client) PutObjectFromFile(bucket, object string, file *os.File) error {
	headers := make(http.Header)

	if dotPos := strings.LastIndex(file.Name(), "."); dotPos == -1 {
		headers.Add("Content-Type", DEFAULT_CONTENT_TYPE)
	} else {
		if mimeType := mime.TypeByExtension(file.Name()[dotPos:]); mimeType == "" {
			headers.Add("Content-Type", DEFAULT_CONTENT_TYPE)
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
	_, err = client.ObjectOp("PUT", bucket, object, headers, file)
	return err
}

func (client *Client) GetObject(bucket, object string, headers http.Header) (body io.ReadCloser, err error) {
	if headers == nil {
		headers = make(http.Header)
	}
	response, err := client.ObjectOp("GET", bucket, object, headers, nil)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

func (client *Client) DeleteObject(bucket, object string) error {
	_, err := client.ObjectOp("DELETE", bucket, object, nil, nil)
	return err
}
