package dm

import (
	"encoding/json"
	"github.com/johnzeng/aliyungo/util"
	"io/ioutil"
	"net/http"
	//	"net/url"
)

func (this *Client) SendSms(signatureId, templateId, recNum string, paramMap map[string]string) error {
	initMap := this.newParamMap()

	if bytes, err := json.Marshal(paramMap); nil != err {
		println("mash error")
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
		println("error")
		return rspErr
	} else {

		println(rsp.StatusCode)
		defer rsp.Body.Close()

		body, err := ioutil.ReadAll(rsp.Body)

		if err != nil {
			// handle error
			println("error")
			return err
		}

		println(string(body))

		return nil
	}
}
