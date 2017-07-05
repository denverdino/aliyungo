package dns

import (
	"testing"
)

func TestAddDomain(t *testing.T) {
	client := NewTestClient()
	args := AddDomainArgs{
		DomainName: TestDomainName,
	}

	if res, err := client.AddDomain(&args); err == nil {
		t.Logf("AddDomain %s success, %v", TestDomainName, res)

		deleteDomainArgs := DeleteDomainArgs{
			DomainName: TestDomainName,
		}
		client.DeleteDomain(&deleteDomainArgs)

	} else {
		t.Errorf("Failed to AddDomain, %v", err)
	}
}
