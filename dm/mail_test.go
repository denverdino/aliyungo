package dm

import (
	"os"
	"testing"
)

func TestBatchMail(t *testing.T) {
	ID := os.Getenv("ALI_DM_ACCESS_KEY_ID")
	SECRET := os.Getenv("ALI_DM_ACCESS_KEY_SECRET")
	accountName := os.Getenv("ALI_DM_ACCOUNT_NAME")
	templateName := os.Getenv("ALI_DM_TEMPLATE_NAME")
	receiverName := os.Getenv("ALI_DM_RECEIVER_NAME")
	client := NewClient(ID, SECRET)
	err := client.SendBatchMail(accountName, "0", templateName, receiverName, "")
	if nil != err {
		t.Error(err.Error())
	}
}

func TestSingleMail(t *testing.T) {
	ID := os.Getenv("ALI_DM_ACCESS_KEY_ID")
	SECRET := os.Getenv("ALI_DM_ACCESS_KEY_SECRET")
	accountName := os.Getenv("ALI_DM_ACCOUNT_NAME")
	replyToAddress := os.Getenv("ALI_DM_REPLY_TO_ADDRESS")
	toAddress := os.Getenv("ALI_DM_TO_ADDRESS")
	client := NewClient(ID, SECRET)
	err := client.SendSingleMail(accountName, replyToAddress, "0", toAddress, "", "", "", "")
	if nil != err {
		t.Error(err.Error())
	}
}
