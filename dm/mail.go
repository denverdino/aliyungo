package dm

import (
	"errors"
	"github.com/denverdino/aliyungo/util"
	"io/ioutil"
	"net/http"
)

//remember to setup the accountName in your aliyun console
//addressType should be "1" or "0",
//0:random address, it's recommanded
//1:sender's address
//tagName is optional, you can use "" if you don't wanna use it
//please set the receiverName and template in the console of Aliyun before you call this API,if you use tagName, you should set it as well
func (this *Client) SendBatchMail(accountName, addressType, templateName, receiverName, tagName string) error {
	initMap := this.newParamMap()

	if addressType != "0" && addressType != "1" {
		return errors.New("invalid addressType")
	}

	initMap.Add("Action", "BatchSendMail")
	initMap.Add("AccountName", accountName)
	initMap.Add("AddressType", addressType)
	initMap.Add("TemplateName", templateName)
	initMap.Add("ReceiversName", receiverName)
	if tagName != "" {
		initMap.Add("TagName", tagName)
	}

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

//remember to setup the accountName in your aliyun console
//addressType should be "1" or "0",
//0:random address, it's recommanded
//1:sender's address
//please set the receiverName and template in the console of Aliyun before you call this API,if you use tagName, you should set it as well

//formAlias, subject, htmlBody, textBody are optional
func (this *Client) SendSingleMail(accountName, replyToAddress, addressType, toAddress, formAlias, subject, htmlBody, textBody string) error {
	initMap := this.newParamMap()

	initMap.Add("Action", "SingleSendMail")
	initMap.Add("AccountName", accountName)
	initMap.Add("ReplyToAddress", replyToAddress)
	initMap.Add("AddressType", addressType)
	initMap.Add("ToAddress", toAddress)
	if formAlias != "" {
		initMap.Add("FormAlias", formAlias)
	}
	if subject != "" {
		initMap.Add("Subject", subject)
	}
	if htmlBody != "" {
		initMap.Add("HtmlBody", htmlBody)
	}
	if textBody != "" {
		initMap.Add("TextBody", textBody)
	}

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
