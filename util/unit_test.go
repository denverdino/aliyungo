package util

import (
	"time"
	"testing"
	"errors"
)

func TestWaitForSignalWithTimeout(t *testing.T) {

	attempts := AttemptStrategy{
		Min:   5,
		Total: 5 * time.Second,
		Delay: 200 * time.Millisecond,
	}

	timeoutFunc := func() (bool,error) {
		return false,nil
	}

	begin := time.Now()

	timeoutError := WaitForSignal(attempts, timeoutFunc);
	if(timeoutError != nil) {
		t.Logf("timeout func complete successful")
	} else {
		t.Error("Expect timeout result")
	}

	end := time.Now()
	duration := end.Sub(begin).Seconds()
	if( duration  > (float64(attempts.Min) -1)) {
		t.Logf("timeout func duration is enough")
	} else {
		t.Error("timeout func duration is not enough")
	}

	errorFunc := func() (b bool,e error) {
		err := errors.New("execution failed");
		return false,err
	}

	failedError := WaitForSignal(attempts, errorFunc);
	if(failedError != nil) {
		t.Logf("error func complete successful: " + failedError.Error())
	} else {
		t.Error("Expect error result")
	}

	successFunc := func() (bool,error) {
		return true,nil
	}

	successError := WaitForSignal(attempts, successFunc);
	if(successError != nil) {
		t.Error("Expect success result")
	} else {
		t.Logf("success func complete successful")
	}
}