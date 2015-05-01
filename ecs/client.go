package ecs

import (
	"bytes"
	"encoding/json"
	"github.com/denverdino/aliyungo/util"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const ECS_API_ENDPOINT = "https://ecs.aliyuncs.com"

type Client struct {
	AccessKeyId     string
	AccessKeySecret string
	debug           bool
}

func NewClient(accessKeyId, accessKeySecret string) *Client {
	client := &Client{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		debug:           false,
	}
	return client
}

func getECSErrorFromString(str string) *ECSError {
	return &ECSError{
		ErrorResponse: ErrorResponse{
			Code:    "ECSClientFailure",
			Message: str,
		},
		StatusCode: -1,
	}
}

func getECSError(err error) *ECSError {
	return getECSErrorFromString(err.Error())
}

func (client *Client) SetDebug(debug bool) {
	client.debug = debug
}
func (client *Client) Invoke(action string, args interface{}, response interface{}) *ECSError {

	request := Request{}
	request.init(action, client.AccessKeyId)

	query := util.ConvertToQueryValues(request)
	util.SetQueryValues(args, &query)

	// Sign request
	signature := util.CreateSignatureForRequest(REQUEST_METHOD, &query, client.AccessKeySecret)

	// Generate the request URL
	requestURL := ECS_API_ENDPOINT + "?" + query.Encode() + "&Signature=" + url.QueryEscape(signature)

	httpClient := &http.Client{}
	httpReq, err := http.NewRequest(REQUEST_METHOD, requestURL, nil)
	if err != nil {
		return getECSError(err)
	}

	t0 := time.Now()
	httpResp, err := httpClient.Do(httpReq)
	t1 := time.Now()
	if err != nil {
		return getECSError(err)
	}
	statusCode := httpResp.StatusCode

	log.Printf("Invoke %s %s %d (%v)", REQUEST_METHOD, requestURL, statusCode, t1.Sub(t0))

	body, err := ioutil.ReadAll(httpResp.Body)
	defer httpResp.Body.Close()

	if err != nil {
		return getECSError(err)
	}

	if client.debug {
		var prettyJSON bytes.Buffer
		err = json.Indent(&prettyJSON, body, "", "    ")
		log.Println(string(prettyJSON.Bytes()))
	}

	var ecsError *ECSError

	if statusCode >= 400 && statusCode <= 599 {
		errorResponse := ErrorResponse{}
		err = json.Unmarshal(body, &errorResponse)
		ecsError = &ECSError{
			ErrorResponse: errorResponse,
			StatusCode:    statusCode,
		}
	} else {
		err = json.Unmarshal(body, response)
		//log.Printf("%++v", response)
	}

	if err != nil {
		ecsError = getECSError(err)
	}

	return ecsError
}

func (client *Client) GenerateClientToken() string {
	return util.CreateRandomString()
}
