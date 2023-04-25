package main

import (
	"encoding/json"
	"fmt"
)

type Mock struct {
	Id          string      `json:"id,omitempty" bson:"_id,omitempty"`
	ApiName     string      `json:"apiName"bson:"apiName"`
	Key         string      `json:"key" bson:"key"`
	Value       interface{} `json:"value" bson:"value"`
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
