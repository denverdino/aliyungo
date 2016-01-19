package ram

import (
	"os"
)

var (
	AccessKeyId     = os.Getenv("AccessKeyId")
	AccessKeySecret = os.Getenv("AccessKeySecret")
)

func NewTestClient() RamClientInterface {
	client := NewClient(AccessKeyId, AccessKeySecret)
	return client
}
