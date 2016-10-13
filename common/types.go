package common

type InternetChargeType string
type InstanceChargeType string

const (
	PayByBandwidth = InternetChargeType("PayByBandwidth")
	PayByTraffic   = InternetChargeType("PayByTraffic")

	PrePaid  = InstanceChargeType("PrePaid")  //包年包月
	PostPaid = InstanceChargeType("PostPaid") //按量付费
)
