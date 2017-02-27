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

type NetType string

const (
	Internet = NetType("Internet")
	Intranet = NetType("Intranet")
)

type TimeType string

const (
	Hour  = TimeType("Hour")
	Day   = TimeType("Day")
	Month = TimeType("Month")
	Year  = TimeType("Year")
)

type NetworkType string

const (
	Classic = NetworkType("Classic")
	VPC     = NetworkType("VPC")
)

type BusinessInfo struct {
	Pack       string `json:"pack,omitempty"`
	ActivityId string `json:"activityId,omitempty"`
}
