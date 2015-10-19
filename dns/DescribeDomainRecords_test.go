package dns

import (
	"testing"
)

func TestDescribeDomainRecords(t *testing.T) {
	//prepare
	client := NewTestClient()
	describeArgs := DescribeDomainRecordsArgs{
		DomainName: TestDomainName,
		PageSize:   500,
	}

	describeResponse, err := client.DescribeDomainRecords(&describeArgs)
	if err == nil {
		t.Logf("DescribeDomainRecords success: TotalCount:%d ", describeResponse.TotalCount)
	} else {
		t.Fatalf("Failed to DescribeDomainRecords: %s", describeArgs.DomainName)
	}
}
