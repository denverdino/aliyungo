package crm

import (
	"testing"
)

func TestQueryCustomerLabel(t *testing.T) {
	client := NewTestClientForDebug()

	labels, err := client.QueryCustomerLabel(TestLabelSeries)
	if err != nil {
		t.Errorf("Failed to query custom labels %v", err)
	} else {
		t.Logf("Successfully to query custom lables : %v", labels)
	}
}

func TestIsFinanceAuthor(t *testing.T) {
	client := NewTestClientForDebug()

	isFinanceUser := client.IsFinanceUser()
	if isFinanceUser {
		t.Logf("Finance user ")
	} else {
		t.Logf("Not finance user")
	}
}
