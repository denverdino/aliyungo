package ros

type RosClientInterface interface {
	//stack
	ListStacks()
	CreateStack()
	DescribeStackInfo()
	DeleteStack()
	AbandonStack()

	//resource
	ListResources()
	DescribeResourceInfo()
	DescribeResourceTypes()
	DescribeResourceTemplateInfo()

	//templates
	DescribeTemplate()
	ValidateTemplate()

	//event
	ListAllEvents()
}
