package oss

import (
	"bytes"
	"fmt"

	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	//"errors"
	"github.com/denverdino/aliyungo/util"
)

const (
	DefaultContentType     = "application/octet-stream"
	SelfDefineHeaderPrefix = "x-oss-"
	DefaultHost            = "http://oss.aliyuncs.com"
	HEADER_X_OSS_ACL       = "x-oss-acl"
)

type Client struct {
	AccessKeyId     string
	AccessKeySecret string
	Endpoint        string
	httpClient      *http.Client
	debug           bool
}

func NewOSSClient(endpoint string, accessKeyId string, accessKeySecret string) *Client {
	return &Client{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Endpoint:        endpoint,
		httpClient:      &http.Client{},
		debug:           false,
	}
}

func (client *Client) SetAccessKeyId(id string) {
	client.AccessKeyId = id
}

func (client *Client) SetAccessKeySecret(secret string) {
	client.AccessKeySecret = secret
}

func (client *Client) createSignature(method string, headers http.Header, resource string) string {

	contentMd5 := headers.Get("Content-Md5")
	contentType := headers.Get("Content-Type")
	date := headers.Get("Date")
	canonicalizedResource := resource

	_, canonicalizedHeader := CanonicalizeHeader(headers)

	stringToSign := method + "\n" + contentMd5 + "\n" + contentType + "\n" + date + "\n" + canonicalizedHeader + canonicalizedResource
	return util.CreateSignature(stringToSign, client.AccessKeySecret)
}

//func (client *Client) SignUrlAuthWithExpireTime(method, urladdr string, headers http.Header, resource string, timeout int) string {
//	sendTime := headers.Get("Date")
//	if sendTime == "" {
//		headers.Add("Date", getGMTime())
//	}
//	auth := client.getAuthorization(method, headers, resource)
//	params := make(url.Values)
//	params.Add("OSSAccessKeyId", client.AccessKeyId)
//	params.Add("Expires", getExpires())
//	params.Add("Signature", auth)
//	return url.QueryEscape(urladdr + "?" + params.Encode())
//}

func (client *Client) createSignForNormalAuth(method string, headers http.Header, resource string) string {
	return "OSS " + client.AccessKeyId + ":" + client.createSignature(method, headers, resource)
}

//Have to break the abstraction to append keys with lower case.
func CanonicalizeHeader(headers http.Header) (newHeaders http.Header, result string) {
	var canonicalizedHeaders []string = make([]string, 0)
	newHeaders = http.Header{}

	for k, v := range headers {
		if lower := strings.ToLower(k); strings.HasPrefix(lower, SelfDefineHeaderPrefix) {
			newHeaders[lower] = v
			canonicalizedHeaders = append(canonicalizedHeaders, lower)
		} else {
			newHeaders[k] = v
		}
	}

	sort.Strings(canonicalizedHeaders)

	var canonicalizedHeader string

	for _, k := range canonicalizedHeaders {
		canonicalizedHeader += k + ":" + headers.Get(k) + "\n"
	}
	return newHeaders, canonicalizedHeader
}

func (client *Client) BucketOp(method, bucket string, headers http.Header, params url.Values, response interface{}) error {

	if params == nil {
		params = make(url.Values)
	}
	url := "/" + bucket + "?" + params.Encode()
	var resource string

	if _, ok := params["acl"]; ok {
		resource = "/" + bucket + "?acl"
	} else {
		resource = "/" + bucket
	}
	log.Printf("resource: %s", resource)
	_, err := client.Invoke(method, url, resource, strings.NewReader(""), headers, response)

	return err
	//return client.Do(method, url, resource, strings.NewReader(""), headers)
}

func (client *Client) GetService() (bucketList *BucketList, err error) {
	method := "GET"
	url := "/"
	headers := make(http.Header)
	resource := "/"
	//response := client.Do(method, url, resource, strings.NewReader(""), headers)
	bucketList = &BucketList{}
	_, err = client.Invoke(method, url, resource, strings.NewReader(""), headers, &bucketList)
	return
}

func (client *Client) GetBucketAcl(bucket string) (result *AccessControlPolicy, err error) {
	params := make(url.Values)
	params.Add("acl", "")
	result = &AccessControlPolicy{}
	err = client.BucketOp("GET", bucket, nil, params, result)
	if err != nil {
		return nil, err
	}
	return
}

func (client *Client) GetBucket(bucket, prefix, marker, delimiter, maxkeys string) (result *ListBucketResult, err error) {
	return client.listBucket(bucket, prefix, marker, delimiter, maxkeys)
}

func (client *Client) ListBucket(prefix, marker, delimiter, maxkeys string) (result *ListBucketResult, err error) {
	return client.listBucket("", prefix, marker, delimiter, maxkeys)
}

func (client *Client) listBucket(bucket, prefix, marker, delimiter, maxkeys string) (result *ListBucketResult, err error) {
	params := make(url.Values)
	if prefix != "" {
		params.Add("prefix", prefix)
	}
	if marker != "" {
		params.Add("marker", marker)
	}
	if delimiter != "" {
		params.Add("delimiter", delimiter)
	}
	if maxkeys != "" {
		params.Add("max-keys", maxkeys) //TODO: check max-keys <= 1000?
	}
	result = &ListBucketResult{}
	err = client.BucketOp("GET", bucket, nil, params, result)
	if err != nil {
		return nil, err
	}
	return
}

func (client *Client) PutBucket(bucket, acl string) error {
	headers := make(http.Header)
	if acl != "" {
		headers.Add(HEADER_X_OSS_ACL, acl)
	}
	return client.BucketOp("PUT", bucket, headers, nil, nil)
}

func (client *Client) CreateBucket(bucket, acl string) error {
	return client.PutBucket(bucket, acl)
}

func (client *Client) DeleteBucket(bucket string) error {
	return client.BucketOp("DELETE", bucket, nil, nil, nil)
}

func (client *Client) ObjectOp(method, bucket, object string, headers http.Header, data io.Reader, response interface{}) (httpResp *http.Response, err error) {
	resource := "/" + bucket + "/" + object
	urladdr := resource
	return client.Invoke(method, urladdr, resource, data, headers, response)
}

func (client *Client) PutObjectFromFile(bucket, object string, headers http.Header, file *os.File) error {
	if headers == nil {
		headers = make(http.Header)
	}
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
	headers.Add("Expect", "100-Continue") //TODO: what's this?

	log.Printf("Header in file put: %v", headers)
	_, err = client.ObjectOp("PUT", bucket, object, headers, file, nil)
	return err
}

func (client *Client) GetObject(bucket, object string, headers http.Header) (body io.ReadCloser, err error) {
	if headers == nil {
		headers = make(http.Header)
	}
	response, err := client.ObjectOp("GET", bucket, object, headers, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

func (client *Client) DeleteObject(bucket, object string) error {
	_, err := client.ObjectOp("DELETE", bucket, object, nil, nil, nil)
	return err
}

func getOSSErrorFromString(str string) error {
	return &OSSError{
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

func (client *Client) Invoke(method, url, resource string, body io.Reader, headers http.Header, response interface{}) (httpResp *http.Response, err error) {
	if headers == nil {
		headers = make(http.Header)
	}
	headers.Set("Date", getGMTime())
	if client.AccessKeySecret != "" {
		headers.Add("Authorization", client.createSignForNormalAuth(method, headers, resource))
	} else {
		headers.Add("Authorization", client.AccessKeyId)
	}

	var host string
	if h := headers.Get("Host"); h != "" {
		host = h
	} else {
		host = client.Endpoint
	}
	req, err := http.NewRequest(method, host+url, body)
	if err != nil {
		log.Printf("Error Creating Request: %v\n", err)
	}
	req.Header = headers
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
		httpResp.Body = NOPCloser{bytes.NewReader(body)}
	}

	if statusCode >= 400 && statusCode <= 599 {
		ossError := &OSSError{}
		decoder := xml.NewDecoder(httpResp.Body)
		err := decoder.Decode(ossError)
		if err != nil {
			return nil, getOSSError(err)
		}
		ossError.StatusCode = statusCode
		return nil, ossError
	} else {
		if response != nil {
			decoder := xml.NewDecoder(httpResp.Body)
			err := decoder.Decode(response)
			if err != nil {
				return nil, getOSSError(err)
			}
			if client.debug {
				log.Printf("Response: %++v", response)
			}
		}
	}
	return httpResp, nil
}

func CopyHeader(header http.Header) (newHeader http.Header) {
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
	return fmt.Sprint("%d", time.Now().Unix()+60)
}
