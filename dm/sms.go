package dm

import (
	"encoding/json"
	"errors"
	"github.com/denverdino/aliyungo/util"
	"io/ioutil"
	"net/http"
	//	"net/url"
)

//please set the signature and template in the console of Aliyun before you call this API
func (this *Client) SendSms(signatureId, templateId, recNum string, paramMap map[string]string) error {
	initMap := this.newParamMap()

	if bytes, err := json.Marshal(paramMap); nil != err {
		println("marsh error")
	} else {
		value := string(bytes)
		initMap.Add("ParamString", value)
	}

	initMap.Add("Action", "SingleSendSms")
	initMap.Add("SignName", signatureId)
	initMap.Add("TemplateCode", templateId)
	initMap.Add("RecNum", recNum)

	signature := util.CreateSignatureForRequest("GET", initMap, this.accessKeySecret+"&")
	initMap.Add("Signature", signature)
	finalUrl := Url + "?" + initMap.Encode()
	println(finalUrl)

	if rsp, rspErr := http.Get(finalUrl); nil != rspErr {
		return rspErr
	} else {
		defer rsp.Body.Close()

		body, err := ioutil.ReadAll(rsp.Body)

		if rsp.StatusCode > 400 {
			return errors.New(string(body))
		}

		if err != nil {
			// handle error
			return err
		}
		//only print the request id for debuging
		println(string(body))

		return nil
	}
}
