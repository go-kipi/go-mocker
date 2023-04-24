package main

type Mock struct {
	ApiName     string      `json:"apiName"bson:"apiName"`
	Key         string      `json:"key" bson:"key"`
	Value       interface{} `json:"value" bson:"value"`
	Reply       string      `json:"reply" bson:"reply"`
	HandlerType string      `json:"handlerType" bson:"handlerType"`
	TimeOut     int         `json:"timeOut,omitempty" bson:"timeOut,omitempty"`
}

type updateMock struct {
	Id   string `json:"id"`
	Mock `json:"mock"`
}
