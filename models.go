package main

import (
	"encoding/json"
	"fmt"
)

const (
	XForwardedFor = "X-Forwarded-For"
	XRealIp       = "X-Real-Ip"
)

type Mock struct {
	Id          string      `json:"id,omitempty" bson:"_id,omitempty"`
	ApiName     string      `json:"apiName"bson:"apiName"`
	Key         string      `json:"reqKey" bson:"key"`
	Value       interface{} `json:"reqValue" bson:"value"`
	Reply       string      `json:"reply" bson:"reply"`
	HandlerType string      `json:"handlerType" bson:"handlerType"`
	TimeOut     int         `json:"timeOut,omitempty" bson:"timeOut,omitempty"`
}

func (mock *Mock) mockReply() (jsonMap map[string]interface{}) {
	err := json.Unmarshal([]byte(mock.Reply), &jsonMap)
	if err != nil {
		fmt.Println("mockReply ", err)
	}
	return jsonMap
}
