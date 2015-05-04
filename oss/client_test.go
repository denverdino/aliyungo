package oss

import (
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"testing"
)

var (
	client     = NewOSSClient(TestRegion, false, TestAccessKeyId, TestAccessKeySecret)
	bucket     = client.Bucket(TestBucket)
	testObject = "api_handler.go"
)

func TestPutObject(t *testing.T) {
	err := bucket.Put(testObject, []byte("Just for text"), "")
	if err != nil {
		t.Errorf("Unable to put Object: %++v", err)
	}

}

func TestPutObjectFromFile(t *testing.T) {
	file, err := os.Open("../README.md")
	if err != nil {
		t.Fatal(err)
	}
	err = bucket.PutFile(testObject, file)
	if err != nil {
		t.Errorf("Unable to put object: %v", err)
	}

}

func TestGetObject(t *testing.T) {
	body, err := bucket.Get(testObject)
	if err != nil {
		t.Fatalf("Unable to get object: %v", err)
	}
	t.Logf("Content of object %s:", testObject)
	t.Log(string(body))
}

// ReadStream retrieves an io.ReadCloser for the content stored at "path" with a
// given byte offset.
func TestGetResponse(t *testing.T) {
	headers := make(http.Header)
	offset := int64(100)
	headers.Add("Range", "bytes="+strconv.FormatInt(offset, 10)+"-")
	httpResp, err := bucket.GetResponse(testObject)
	if err != nil {
		t.Fatalf("Unable to get object with offset: %v", err)
	}
	data, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		t.Fatalf("Unable to get object with offset: %v", err)
	}
	t.Logf("Content of object %s from offset %d:", testObject, offset)
	t.Log(string(data))
	return
}

func TestDeleteObject(t *testing.T) {
	err := bucket.Delete(testObject)
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
	result, err := bucket.List("", "", "", 0)
	if err != nil {
		t.Errorf("Unable to list Bucket with no params: %v", err)
	} else {
		t.Logf("Result: %++v", result)
	}
	result, err = bucket.List("", "", "", 10)
	if err != nil {
		t.Errorf("Unable to list Bucket with 10 maxkeys: %v", err)
	} else {
		t.Logf("Result: %++v", result)
	}
}

func TestGetBucketACL(t *testing.T) {
	result, err := bucket.GetBucketAcl()
	if err != nil {
		t.Errorf("Unable to get Bucket ACL: %v", err)
	} else {
		t.Logf("Result: %++v", result)
	}
}

func TestPutBucket(t *testing.T) {
	testBucket := client.Bucket("denverdino-test")
	err := testBucket.PutBucket("")
	if err != nil {
		t.Errorf("Unable to create a new bucket with no acl specified: %v", err)
	}
	err = testBucket.PutBucket("private")
	if err != nil {
		t.Errorf("Unable to create a new bucket with private acl: %v", err)
	}
	err = testBucket.DeleteBucket()
	if err != nil {
		t.Errorf("Unable to delete the test bucket: %v", err)
	}
}
