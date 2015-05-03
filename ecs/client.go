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
	AccessKeyId     string //Access Key Id
	AccessKeySecret string //Access Key Secret
	debug           bool
	httpClient      *http.Client
}

//Create a new instance of ECS client
func NewClient(accessKeyId, accessKeySecret string) *Client {
	client := &Client{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret + "&",
		debug:           false,
		httpClient:      &http.Client{},
	}
	return client
}

func (client *Client) SetAccessKeyId(id string) {
	client.AccessKeyId = id
}

func (client *Client) SetAccessKeySecret(secret string) {
	client.AccessKeySecret = secret + "&"
}

func getECSErrorFromString(str string) error {
	return &ECSError{
		ErrorResponse: ErrorResponse{
			Code:    "ECSClientFailure",
			Message: str,
		},
		StatusCode: -1,
	}
}

func getECSError(err error) error {
	return getECSErrorFromString(err.Error())
}

func (client *Client) SetDebug(debug bool) {
	client.debug = debug
}

//Invoke the ECS request
func (client *Client) Invoke(action string, args interface{}, response interface{}) error {

	request := Request{}
	request.init(action, client.AccessKeyId)

	query := util.ConvertToQueryValues(request)
	util.SetQueryValues(args, &query)

	// Sign request
	signature := util.CreateSignatureForRequest(REQUEST_METHOD, &query, client.AccessKeySecret)

	// Generate the request URL
	requestURL := ECS_API_ENDPOINT + "?" + query.Encode() + "&Signature=" + url.QueryEscape(signature)

	httpReq, err := http.NewRequest(REQUEST_METHOD, requestURL, nil)
	if err != nil {
		return getECSError(err)
	}

	t0 := time.Now()
	httpResp, err := client.httpClient.Do(httpReq)
	t1 := time.Now()
	if err != nil {
		return getECSError(err)
	}
	statusCode := httpResp.StatusCode

	log.Printf("Invoke %s %s %d (%v)", REQUEST_METHOD, requestURL, statusCode, t1.Sub(t0))

	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)

	if err != nil {
		return getECSError(err)
	}

	if client.debug {
		var prettyJSON bytes.Buffer
		err = json.Indent(&prettyJSON, body, "", "    ")
		log.Println(string(prettyJSON.Bytes()))
	}

	if statusCode >= 400 && statusCode <= 599 {
		errorResponse := ErrorResponse{}
		err = json.Unmarshal(body, &errorResponse)
		ecsError := &ECSError{
			ErrorResponse: errorResponse,
			StatusCode:    statusCode,
		}
		return ecsError
	} else {
		err = json.Unmarshal(body, response)
		//log.Printf("%++v", response)
		if err != nil {
			return getECSError(err)
		}
	}

	return nil
}

// Generate the Client Token
func (client *Client) GenerateClientToken() string {
	return util.CreateRandomString()
}
