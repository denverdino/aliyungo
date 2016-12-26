package grandcanal

import (
	"github.com/denverdino/aliyungo/common"
)

// CreateDomainArgs represents arguments to create domain
type CreateDomainArgs struct {
	DomainName      string
	Description     string
	RetentionPeriod int32
}

type CreateDomainResponse struct {
	common.Response
}

// CreateDomain create a new domain
func (client *Client) CreateDomain(args *CreateDomainArgs) (requestId string, err error) {
	response := &CreateDomainResponse{}
	err = client.Invoke("CreateDomain", args, &response)
	if err != nil {
		return "", err
	}
	return response.RequestId, nil
}

// DescribeDomainArgs represents arguments to describe domain info
type DescribeDomainArgs struct {
	DomainName string
}

type DescribeDomainResponse struct {
	common.Response
	Domain Domain
}

type Domain struct {
	DomainId        int64
	DomainName      string
	Description     string
	RetentionPeriod int32
	CreateTime      string
	Deleted         bool
}

// DescribeDomain describe domain information
func (client *Client) DescribeDomain(args *DescribeDomainArgs) (domain *Domain, err error) {
	response := &DescribeDomainResponse{}
	err = client.Invoke("DescribeDomain", args, &response)
	if err != nil {
		return nil, err
	}
	return &response.Domain, nil
}

// DeleteDomainArgs represents arguments to delete domain
type DeleteDomainArgs struct {
	DomainName string
}

type DeleteDomainResponse struct {
	common.Response
}

// DeleteDomain remove a domain
func (client *Client) DeleteDomain(args *DeleteDomainArgs) (requestId string, err error) {
	response := &DeleteDomainResponse{}
	err = client.Invoke("DeleteDomain", args, &response)
	if err != nil {
		return "", err
	}
	return response.RequestId, nil
}
