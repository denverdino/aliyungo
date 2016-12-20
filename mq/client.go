package mq

import (
	"strconv"
	"fmt"
	"errors"
	"encoding/json"
)

type Client struct {
	AccessKey  string
	SecretKey  string
	Endpoint   string
	Topic      string
	ProducerId string
	ConsumerId string
	Key        string
	Tag        string
}

func NewClient(ak string, sk string, endpoint string, topic string,
	producerId string, consumerId string, key string, tag string) (client *Client) {
	client = &Client{
		AccessKey:  ak,
		SecretKey:  sk,
		Endpoint:   endpoint,
		Topic:      topic,
		ProducerId: producerId,
		ConsumerId: consumerId,
		Key:        key,
		Tag:        tag,
	}
	return client
}

func getSendUrl(endpoint string, topic string, time int64, tag string, key string) string {
	return endpoint + "/message/?topic=" + topic + "&time=" +
		strconv.FormatInt(time, 10) + "&tag=" + tag + "&key=" + key
}

func getSendSign(topic string, producerId string, messageBody string, time int64, sk string) (sign string) {
	signStr := topic + newline + producerId + newline + Md5(messageBody) + newline + strconv.FormatInt(time, 10)
	sign = HamSha1(signStr, []byte(sk))
	return sign
}

func getSendHeader(ak string, sign string, producerId string) (header map[string]string, err error) {
	if producerId == "" {
		return nil, fmt.Errorf("producer id is not provided")
	}
	header = make(map[string]string, 0)
	header["AccessKey"] = ak
	header["Signature"] = sign
	header["ProducerId"] = producerId
	return header, nil
}

func (client *Client) send(time int64, message string) (msgId string, err error) {
	url := getSendUrl(client.Endpoint, client.Topic, time, client.Tag, client.Key)
	sign := getSendSign(client.Topic, client.ProducerId, message, time, client.SecretKey)
	header, err := getSendHeader(client.AccessKey, sign, client.ProducerId)
	if err != nil {
		return "", err
	}
	response, status, err := httpPost(url, header, []byte(message))
	if err != nil {
		return "", err
	}

	statusMessage := getStatusCodeMessage(status)
	if statusMessage != "" {
		return "", errors.New(statusMessage)
	}

	var rs interface{}
	err = json.Unmarshal(response, &rs)
	if err != nil {
		return "", err
	}

	result := rs.(map[string]interface{})

	sendStatus := result["sendStatus"].(string)
	if sendStatus != "SEND_OK" {
		return "", errors.New(sendStatus)
	}

	return result["msgId"].(string), nil
}

func (client *Client) receive() {

}
