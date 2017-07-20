package common

// Region represents ECS region
type Region string

// Constants of region definition
const (
	Hangzhou        = Region("cn-hangzhou")
	Qingdao         = Region("cn-qingdao")
	Beijing         = Region("cn-beijing")
	BeiJingHpc      = Region("cn-beijing-hpc")
	Hongkong        = Region("cn-hongkong")
	Shenzhen        = Region("cn-shenzhen")
	ZhangJiaKou     = Region("cn-zhangjiakou")
	USWest1         = Region("us-west-1")
	USEast1         = Region("us-east-1")
	APSouthEast1    = Region("ap-southeast-1")
	APNorthEast1    = Region("ap-northeast-1")
	Shanghai        = Region("cn-shanghai")
	ShenZhenFinance = Region("cn-shenzhen-finance-1")
	ShanghaiFinance = Region("cn-shanghai-finance-1")
	EUCentral1      = Region("eu-central-1")   //德国
	MEEast1         = Region("me-east-1")      //中东
	APSouthEast2    = Region("ap-southeast-2") //澳洲
)

var ValidRegions = []Region{Hangzhou, Qingdao, Beijing, BeiJingHpc, Shenzhen, ZhangJiaKou, Hongkong,
	Shanghai, USWest1, USEast1, APSouthEast1, APNorthEast1, ShenZhenFinance, ShanghaiFinance,
	EUCentral1, MEEast1, APSouthEast2}
