package oss

import (
	"io/ioutil"
	"os"
	"testing"
)

var (
	client     = NewOSSClient(TEST_REGION.GetInternetEndpoint(), TEST_ACCESS_KEY_ID, TEST_ACCESS_KEY_SECRET)
	testObject = "api_handler.go"
)

func TestPutObjectFromFile(t *testing.T) {
	file, err := os.Open("../README.md")
	if err != nil {
		t.Fatal(err)
	}
	err = client.PutObjectFromFile(TEST_BUCKET, testObject, nil, file)

	if err != nil {
		t.Errorf("Unable to put Object: %++v", err)
	}

}

func TestGetObject(t *testing.T) {
	body, err := client.GetObject(TEST_BUCKET, testObject, nil)
	//defer body.Close()
	if err != nil {
		t.Errorf("Unable to get object: %v", err)
	}
	contents, err := ioutil.ReadAll(body)
	if err != nil {
		t.Log(err)
	}
	t.Log(string(contents))
}

func TestDeleteObject(t *testing.T) {
	err := client.DeleteObject(TEST_BUCKET, testObject)
	if err != nil {
		t.Errorf("Unable to del object: %v", err)
	}
}

func TestGetService(t *testing.T) {
	bucketList, err := client.GetService()
	if err != nil {
		t.Errorf("Unable to get service: %v", err)
	} else {
		t.Logf("GetService: %++v", bucketList)
	}
}

func TestGetBucket(t *testing.T) {
	result, err := client.GetBucket(TEST_BUCKET, "", "", "", "")
	if err != nil {
		t.Errorf("Unable to list Bucket with no params: %v", err)
	} else {
		t.Logf("Result: %++v", result)
	}
	result, err = client.GetBucket(TEST_BUCKET, "", "", "", "10")
	if err != nil {
		t.Error("Unable to list Bucket with 10 maxkeys: %v", err)
	} else {
		t.Logf("Result: %++v", result)
	}
}

func TestGetBucketACL(t *testing.T) {
	result, err := client.GetBucketAcl(TEST_BUCKET)
	if err != nil {
		t.Errorf("Unable to get Bucket ACL: %v", err)
	} else {
		t.Logf("Result: %++v", result)
	}
}

func TestPutBucket(t *testing.T) {
	err := client.PutBucket("denverdino-test", "")
	if err != nil {
		t.Error("Unable to create a new bucket with no acl specified: %v", err)
	}
	err = client.PutBucket("denverdino-test", "private")
	if err != nil {
		t.Error("Unable to create a new bucket with private acl: %v", err)
	}
	err = client.DeleteBucket("denverdino-test")
	if err != nil {
		t.Error("Unable to delete the test bucket: %v", err)
	}
}
