package util

import (
	"testing"
	"time"
)

type TestStruct struct {
	Format      string
	Version     string
	AccessKeyId string
	Timestamp   time.Time
	Empty       string
	IntValue    int
	BoolPtr     *bool
	StringArray []string
}

func TestConvertToQueryValues(t *testing.T) {

	request := TestStruct{
		Format:      "JSON",
		Version:     "1.0",
		Timestamp:   time.Date(2015, time.Month(5), 26, 1, 2, 3, 4, time.UTC),
		IntValue:    10,
		StringArray: []string{"abc", "xyz"},
	}
	result := ConvertToQueryValues(&request).Encode()
	const expected_result = "Format=JSON&IntValue=10&StringArray=%5B%22abc%22%2C%22xyz%22%5D&Timestamp=2015-05-26T01%3A02%3A03Z&Version=1.0"
	if result != expected_result {
		t.Error("Incorrect encoding: ", result)
	}

}
