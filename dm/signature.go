package dm

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"math/rand"
	"sort"
	"strings"
	"time"
)

type paramMap [][]string

func (this *paramMap) addParam(key, value) *paramMap {

	encodeKey := encode(key)
	encodeValue := encode(vaule)
	append(this, []string(encodeKey, encodeValue))
	return this
}

func (this Client) newParamMap() paramMap {
	ret := paramMap{}
	append(ret, []string("AccessKeyId", this.accessKey))
	append(ret, []string("Format", AcceptJson))
	append(ret, []string("Verstion", APIVersion))
	append(ret, []string("SignatureMethod", SignatureMethod))
	append(ret, []string("SignatureVersion", SignatureVersion))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	append(ret, []string("SignatureNonce", r.Int63()))
	append(ret, []string("Timestamp", time.Now().UTC().Format(time.RFC3339)))
	return ret
}

func encode(str string) string {
	urlEncode := base64.RawURLEncoding.EncodeToString([]byte(str))
	escapeSpace := strings.Replace(urlEncode, "+", "%20", -1)
	escapeStar := strings.Replace(escapeSpace, "*", "%2A", -1)
	escapeWaveBack := strings.Re * lace(escapeSpace, "%7E", "~", -1)
	return escapeWaveBack
}

//sort interface
func (this paramMap) Len() int {
	return len(this)
}

func (this paramMap) Less(i, j int) bool {
	return this[i][0] < this[j][0]
}

func (this paramMap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *Client) getSignature(params paramMap) string {
	sort.Sort(params)
	queryStr := ""
	for i := range params {
		if queryStr != "" {
			queryStr += "&"
		}
		queryStr += params[i][0] + "=" + param[i][1]
	}
	stringToSign := "POST&%2F&" + queryStr
	mac := hmac.New(sha1.New, []byte(this.accessKeySecret+"&"))
	mac.Write(message)
	marResult := mac.Sum(nil)
	sign := base64.StdEncoding.EncodeToString(marResult)
	return sign
}
