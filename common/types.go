package common

type InternetChargeType string

const (
	PayByBandwidth = InternetChargeType("PayByBandwidth")
	PayByTraffic   = InternetChargeType("PayByTraffic")
)

type InstanceChargeType string

const (
	PrePaid  = InstanceChargeType("PrePaid")
	PostPaid = InstanceChargeType("PostPaid")
)

type ServiceCode string

type DescribeEndpointArgs struct {
	Id          Region
	ServiceCode string
	Type        string
}

type EndpointItem struct {
	Protocols struct {
		Protocols []string
	}
	Type        string
	Namespace   string
	Id          Region
	SerivceCode string
	Endpoint    string
}

type DescribeEndpointResponse struct {
	Response
	EndpointItem
}
