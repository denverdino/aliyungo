package common

import (
	"testing"
)

func TestLoadEndpointFromFile(t *testing.T) {

}

func TestClient_DescribeOpenAPIEndpoint(t *testing.T) {
	client := NewTestClientForDebug()
	endpoint := client.DescribeOpenAPIEndpoint(Region(TestRegionId), TestServiceCode)
	t.Logf("endpoint=%s", endpoint)
}
