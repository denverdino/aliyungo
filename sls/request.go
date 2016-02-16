package sls

import (
	"bytes"
	"errors"
	"github.com/denverdino/aliyungo/util"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"encoding/json"
)

type request struct {
	endpoint    string
	method      string
	path        string
	contentType string
	params      map[string]string
	headers     map[string]string
	payload     []byte
}

func (req *request) url() string {
	params := &url.Values{}

	if req.params != nil {
		for k, v := range req.params {
			params.Set(k, v)
		}
	}

	u := url.URL{
		Scheme:   "http",
		Host:     req.endpoint,
		Path:     req.path,
		RawQuery: params.Encode(),
	}
	return u.String()
}

func (client *Client) doRequest(req *request) (*http.Response, error) {

	payload := req.payload

	if req.headers == nil {
		req.headers = make(map[string]string)
	}

	if req.endpoint == "" {
		req.endpoint = client.endpoint
	}

	contentLength := "0"

	if payload != nil {
		contentLength = strconv.Itoa(len(payload))
	}

	req.headers["Content-Type"] = req.contentType
	req.headers["Content-Length"] = contentLength
	req.headers["x-log-bodyrawsize"] = contentLength
	req.headers["Date"] = util.GetGMTime()
	req.headers["Host"] = req.endpoint
	req.headers["x-log-apiversion"] = client.version
	req.headers["x-log-signaturemethod"] = "hmac-sha1"

	client.signRequest(req, payload)

	var reader io.Reader

	if payload != nil {
		reader = bytes.NewReader(payload)
	}

	hreq, err := http.NewRequest(req.method, req.url(), reader)

	for k, v := range req.headers {
		if v != "" {
			hreq.Header.Set(k, v)
		}
	}

	if err != nil {
		return nil, err
	}
	resp, err := client.httpClient.Do(hreq)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(data))
	}
	return resp, nil
}

func (client *Client) requestWithJsonResponse(req *request, v interface{}) error {
	resp, err := client.doRequest(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

func (client *Client) requestWithClose(req *request) error {
	resp, err := client.doRequest(req)
	if err != nil {
		return err
	}

	resp.Body.Close()
	return nil
}
