package dns

import (
	"testing"
)

func TestAddDomainRecord(t *testing.T) {
	client := NewTestClient()
	addDomainRecordArgs := AddDomainRecordArgs{
		DomainName: TestDomainName,
		RR:         "testaddrecord",
		Type:       ARecord,
		Value:      "8.8.8.8",
	}
	response, err := client.AddDomainRecord(&addDomainRecordArgs)
	if err == nil {
		t.Logf("AddDomainRecord: testaddr for domain: %s Success, %v",
			TestDomainName, response)

		deleteDomainRecordArgs := DeleteDomainRecordArgs{
			RecordId: response.RecordId,
		}
		client.DeleteDomainRecord(&deleteDomainRecordArgs)
	} else {
		t.Errorf("Failed to AddDomainRecord: testaddr for domain: %s", TestDomainName)
	}
}

func TestClient_AddDomainRecord(t *testing.T) {
	client := NewTestClientForDebug()
	addDomainRecordArgs := &AddDomainRecordArgs{
		DomainName: TestDomainName,
		RR:         "*.tunnel-cn-hangzhou-agility",
		Type:       ARecord,
		Value:      "118.31.132.109",
	}

	response, err := client.AddDomainRecord(addDomainRecordArgs)
	if err != nil {
		t.Fatalf("Failed to AddDomainRecord %++v", err)
	} else {
		t.Logf("Response = %++v", response)
	}
}
