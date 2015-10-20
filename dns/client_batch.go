package dns

import (
	"bytes"
	"github.com/denverdino/aliyungo/common"
	"github.com/denverdino/aliyungo/util"
	"time"
	"net/url"
	"encoding/json"
	"net/http"
	"log"
	"io/ioutil"
)

// Http Post Request for Domain Batch Command
// read doc at https://docs.aliyun.com/?spm=5176.100054.201.106.OeZ3dN#/pub/dns/api-reference/batch-related&BaseSpec
func (client *Client) InvokePost(action string,paramKey string,args interface{}, response interface{}) error {
	request := common.Request{}
	request.Format = common.JSONResponseFormat
	request.Timestamp = util.NewISO6801Time(time.Now().UTC())
	request.Version = DNSAPIVersion
	request.SignatureVersion = common.SignatureVersion
	request.SignatureMethod = common.SignatureMethod
	request.SignatureNonce = util.CreateRandomString()
	request.Action = action
	request.AccessKeyId = client.AccessKeyId	
	
	query := util.ConvertToQueryValues(request)
	
	// Sign request
	signature := util.CreateSignatureForRequest("POST", &query, client.AccessKeySecret)

	// Generate the request URL
	requestURL := DNSDefaultEndpoint + "?" + query.Encode() + "&Signature=" + url.QueryEscape(signature)

	// Generate http body
	argsjson, err := json.Marshal(args)
	httpbody := paramKey + "=" + string(argsjson)
	
	httpReq, err := http.NewRequest("POST", requestURL, bytes.NewBuffer([]byte(httpbody)))

	// TODO move to util and add build val flag
	httpReq.Header.Set("X-SDK-Client", `AliyunGO/`+DNSAPIVersion)

	if err != nil {
		return common.GetClientError(err)
	}

	t0 := time.Now()
	httpClient := &http.Client{}
	httpResp, err := httpClient.Do(httpReq)
	t1 := time.Now()
	if err != nil {
		return common.GetClientError(err)
	}
	statusCode := httpResp.StatusCode

	log.Printf("InvokePost %s %d (%v) body:%s", requestURL, statusCode, t1.Sub(t0), httpbody)

	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)

	if err != nil {
		return common.GetClientError(err)
	}

	{
		var prettyJSON bytes.Buffer
		err = json.Indent(&prettyJSON, body, "", "    ")
		log.Println(string(prettyJSON.Bytes()))
	}

	if statusCode >= 400 && statusCode <= 599 {
		errorResponse := common.ErrorResponse{}
		err = json.Unmarshal(body, &errorResponse)
		ecsError := &common.Error{
			ErrorResponse: errorResponse,
			StatusCode:    statusCode,
		}
		return ecsError
	}

	err = json.Unmarshal(body, response)
	//log.Printf("%++v", response)
	if err != nil {
		return common.GetClientError(err)
	}

	return nil
}
