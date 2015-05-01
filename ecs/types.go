package ecs

import (
	"fmt"
	"github.com/denverdino/aliyungo/util"
	"log"
	"time"
)

const (
	API_VERSION       = "2014-05-26"
	SIGNATURE_VERSION = "1.0"
	SIGNATURE_METHOD  = "HMAC-SHA1"
	JSON_FORMAT       = "JSON"
	XML_FORMAT        = "XML"
	REQUEST_METHOD    = "GET"
)

type Request struct {
	Format               string
	Version              string
	AccessKeyId          string
	Signature            string
	SignatureMethod      string
	Timestamp            time.Time
	SignatureVersion     string
	SignatureNonce       string
	ResourceOwnerAccount string
	Action               string
}

func (request *Request) init(action string, AccessKeyId string) {
	request.Format = JSON_FORMAT
	request.Timestamp = time.Now()
	request.Version = API_VERSION
	request.SignatureVersion = SIGNATURE_VERSION
	request.SignatureMethod = SIGNATURE_METHOD
	request.SignatureNonce = util.CreateRandomString()
	request.Action = action
	request.AccessKeyId = AccessKeyId
}

type CommonResponse struct {
	RequestId string
}

type ErrorResponse struct {
	CommonResponse
	HostId  string
	Code    string
	Message string
}

type ECSError struct {
	ErrorResponse
	StatusCode int //Status Code of HTTP Response
}

func (e *ECSError) Error() string {
	return fmt.Sprintf("ECS Error: Status Code %d: Code %s: Message %s", e.StatusCode, e.Code, e.Message)
}

type Pagination struct {
	PageNumber int
	PageSize   int
}

func (p *Pagination) SetPageSize(size int) {
	p.PageSize = size
}

type PaginationResult struct {
	TotalCount int
	PageNumber int
	PageSize   int
}

//Get the Next Page of the result set
func (r *PaginationResult) NextPage() *Pagination {
	if r.PageNumber*r.PageSize >= r.TotalCount {
		return nil
	}
	return &Pagination{PageNumber: r.PageNumber + 1, PageSize: r.PageSize}
}

func (p *Pagination) validate() {
	if p.PageNumber < 0 {
		log.Printf("Invalid PageNumber: %d", p.PageNumber)
		p.PageNumber = 1
	}
	if p.PageSize < 0 {
		log.Printf("Invalid PageSize: %d", p.PageSize)
		p.PageSize = 10
	} else if p.PageSize > 50 {
		log.Printf("Invalid PageSize: %d", p.PageSize)
		p.PageSize = 50
	}
}
