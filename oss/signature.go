package oss

import (
	"github.com/denverdino/aliyungo/util"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

var PARAMS_TO_SIGN = map[string]bool{
	"acl":                          true,
	"location":                     true,
	"logging":                      true,
	"notification":                 true,
	"partNumber":                   true,
	"policy":                       true,
	"requestPayment":               true,
	"torrent":                      true,
	"uploadId":                     true,
	"uploads":                      true,
	"versionId":                    true,
	"versioning":                   true,
	"versions":                     true,
	"response-content-type":        true,
	"response-content-language":    true,
	"response-expires":             true,
	"response-cache-control":       true,
	"response-content-disposition": true,
	"response-content-encoding":    true,
}

func (client *Client) createSignature(request *http.Request) string {
	headers := request.Header
	contentMd5 := headers.Get("Content-Md5")
	contentType := headers.Get("Content-Type")
	date := headers.Get("Date")

	resource := request.URL.Path

	query := request.URL.Query()
	params := make(url.Values)
	for k, v := range query {
		if PARAMS_TO_SIGN[k] {
			params[k] = v
		}
	}

	if len(params) > 0 {
		resource = resource + "?" + util.Encode(params)
	}

	canonicalizedResource := resource

	_, canonicalizedHeader := CanonicalizeHeader(headers)

	stringToSign := request.Method + "\n" + contentMd5 + "\n" + contentType + "\n" + date + "\n" + canonicalizedHeader + canonicalizedResource
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

func (client *Client) createAuthorizationHeader(request *http.Request) string {
	return "OSS " + client.AccessKeyId + ":" + client.createSignature(request)
}

//Have to break the abstraction to append keys with lower case.
func CanonicalizeHeader(headers http.Header) (newHeaders http.Header, result string) {
	var canonicalizedHeaders []string = make([]string, 0)
	newHeaders = http.Header{}

	for k, v := range headers {
		if lower := strings.ToLower(k); strings.HasPrefix(lower, HEADER_X_OSS_PREFIX) {
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
