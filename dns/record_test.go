package dns

import (
	"testing"
)

func TestAddDomainRecord(t *testing.T) {
	client := NewTestClientForDebug()
	addDomainRecordArgs := AddDomainRecordArgs{
		DomainName: TestDomainName,
		RR: "testaddrr", 
		Type: ARecord,
		Value: "8.8.8.8",
	}
	
	instanceId, recordId, err := client.AddDomainRecord(&addDomainRecordArgs)
	if err == nil{
		t.Logf("AddDomainRecord: testaddr for domain: %s Success, instanceId:%s recordId:%s",TestDomainName,instanceId,recordId)	
	}else{
		t.Fatalf("Failed to AddDomainRecord: testaddr for domain: %s",TestDomainName)
	}

}
