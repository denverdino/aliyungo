package ros

const (
	ROS_ENDPOINT    = "https://ros.aliyuncs.com"
	ROS_API_VERSION = "2015-09-01"
	//stack uri
	STACK_URI               = "/stacks"
	STACK_INFO_URI          = STACK_URI + "/{StackName}/{StackId}"
	STACK_ABADON_URI        = STACK_INFO_URI + "/abandon"
	STACK_RESOURCE_URI      = STACK_INFO_URI + "/resources"
	STACK_RESOURCE_INFO_URI = STACK_RESOURCE_URI + "/{ResourceName}"
	STACK_TEMPLATE          = STACK_URI + "/template"
	STACK_EVENT             = STACK_URI + "/events"

	//resource types
	RESOURCE_URI          = "/resource_types"
	RESOURCE_INFO_URI     = RESOURCE_URI + "/{TypeName}"
	RESOURCE_TEMPLATE_URI = RESOURCE_INFO_URI + "/template"

	//OTHER
	VALIDATE_URI = "/validate"
)
