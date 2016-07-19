package slb

import "github.com/hdksky/aliyungo/common"

type UploadServerCertificateArgs struct {
	RegionId              common.Region
	ServerCertificate     string
	ServerCertificateName string
	PrivateKey            string
}

type UploadServerCertificateResponse struct {
	common.Response
	ServerCertificateId   string
	ServerCertificateName string
	Fingerprint           string
}

// UploadServerCertificate Upload server certificate
//
// You can read doc at http://docs.aliyun.com/#pub/slb/api-reference/api-servercertificate&UploadServerCertificate
func (client *Client) UploadServerCertificate(args *UploadServerCertificateArgs) (response *UploadServerCertificateResponse, err error) {
	response = &UploadServerCertificateResponse{}
	err = client.Invoke("UploadServerCertificate", args, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

type DeleteServerCertificateArgs struct {
	RegionId            common.Region
	ServerCertificateId string
}

type DeleteServerCertificateResponse struct {
	common.Response
}

// DeleteServerCertificate Delete server certificate
//
// You can read doc at http://docs.aliyun.com/#pub/slb/api-reference/api-servercertificate&DeleteServerCertificate
func (client *Client) DeleteServerCertificate(regionId common.Region, serverCertificateId string) (err error) {
	args := &DeleteServerCertificateArgs{
		RegionId:            regionId,
		ServerCertificateId: serverCertificateId,
	}
	response := &DeleteServerCertificateResponse{}
	return client.Invoke("DeleteServerCertificate", args, response)
}

type SetServerCertificateNameArgs struct {
	RegionId              common.Region
	ServerCertificateId   string
	ServerCertificateName string
}

type SetServerCertificateNameResponse struct {
	common.Response
}

// SetServerCertificateName Set name of server certificate
//
// You can read doc at http://docs.aliyun.com/#pub/slb/api-reference/api-servercertificate&SetServerCertificateName
func (client *Client) SetServerCertificateName(regionId common.Region, serverCertificateId string, name string) (err error) {
	args := &SetServerCertificateNameArgs{
		RegionId:              regionId,
		ServerCertificateId:   serverCertificateId,
		ServerCertificateName: name,
	}
	response := &SetServerCertificateNameResponse{}
	return client.Invoke("SetServerCertificateName", args, response)
}

type DescribeServerCertificatesArgs struct {
	RegionId            common.Region
	ServerCertificateId string
}

type ServerCertificateType struct {
	RegionId              common.Region
	ServerCertificateId   string
	ServerCertificateName string
	Fingerprint           string
}

type DescribeServerCertificatesResponse struct {
	common.Response
	ServerCertificates struct {
		ServerCertificate []ServerCertificateType
	}
}

// DescribeServerCertificates Describe server certificates
//
// You can read doc at http://docs.aliyun.com/#pub/slb/api-reference/api-servercertificate&DescribeServerCertificates
// cookie fix func name DescribeServerCertificatesArgs to DescribeServerCertificates on 2016-07-19
func (client *Client) DescribeServerCertificates(regionId common.Region, serverCertificateId string) (serverCertificates []ServerCertificateType, err error) {
	args := &DescribeServerCertificatesArgs{
		RegionId:            regionId,
		ServerCertificateId: serverCertificateId,
	}
	response := &DescribeServerCertificatesResponse{}
	err = client.Invoke("DescribeServerCertificates", args, response)
	if err != nil {
		return nil, err
	}
	return response.ServerCertificates.ServerCertificate, err
}

type UploadCACertificateArgs struct {
	RegionId          common.Region
	CACertificate     string
	CACertificateName string
}

type UploadCACertificateResponse struct {
	common.Response
	CACertificateId   string
	CACertificateName string
	Fingerprint       string
}

// UploadCACertificate Upload CA certificate
//
// cookie add this func on 2016-07-19
func (client *Client) UploadCACertificate(args *UploadCACertificateArgs) (response *UploadCACertificateResponse, err error) {
	response = &UploadCACertificateResponse{}
	err = client.Invoke("UploadCACertificate", args, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

type DeleteCACertificateArgs struct {
	RegionId        common.Region
	CACertificateId string
}

type DeleteCACertificateResponse struct {
	common.Response
}

// DeleteCACertificate Delete CA certificate
//
// cookie add this func on 2016-07-19
func (client *Client) DeleteCACertificate(regionId common.Region, CACertificateId string) (err error) {
	args := &DeleteCACertificateArgs{
		RegionId:        regionId,
		CACertificateId: CACertificateId,
	}
	response := &DeleteCACertificateResponse{}
	return client.Invoke("DeleteCACertificate", args, response)
}

type SetCACertificateNameArgs struct {
	RegionId          common.Region
	CACertificateId   string
	CACertificateName string
}

type SetCACertificateNameResponse struct {
	common.Response
}

// SetCACertificateName Set name of CA certificate
//
// cookie add this func on 2016-07-19
func (client *Client) SetCACertificateName(regionId common.Region, CACertificateId string, name string) (err error) {
	args := &SetCACertificateNameArgs{
		RegionId:          regionId,
		CACertificateId:   CACertificateId,
		CACertificateName: name,
	}
	response := &SetCACertificateNameResponse{}
	return client.Invoke("SetCACertificateName", args, response)
}

type DescribeCACertificatesArgs struct {
	RegionId        common.Region
	CACertificateId string
}

type CACertificateType struct {
	RegionId          common.Region
	CACertificateId   string
	CACertificateName string
	Fingerprint       string
}

type DescribeCACertificatesResponse struct {
	common.Response
	CACertificates struct {
		CACertificate []CACertificateType
	}
}

// DescribeCACertificates Describe CA certificates
//
// You can read doc at http://docs.aliyun.com/#pub/slb/api-reference/api-CAcertificate&DescribeCACertificates
// cookie fix func name DescribeCACertificatesArgs to DescribeCACertificates on 2016-07-19
func (client *Client) DescribeCACertificates(regionId common.Region, CACertificateId string) (CACertificates []CACertificateType, err error) {
	args := &DescribeCACertificatesArgs{
		RegionId:        regionId,
		CACertificateId: CACertificateId,
	}
	response := &DescribeCACertificatesResponse{}
	err = client.Invoke("DescribeCACertificates", args, response)
	if err != nil {
		return nil, err
	}
	return response.CACertificates.CACertificate, err
}
