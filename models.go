package main

import "encoding/json"

const (
	mongo_DataBase   = "kipiMocks"
	mongo_Collection = "mocks"
)

type getAllMocks struct {
}

type getMockById struct {
	Id string `json:"id"`
}

type updateMockById struct {
	Id string `json:"id"`
}

type deleteMockById struct {
	Id string `json:"id"`
}

type createMock struct {
	Id   string `json:"id" bson:"id"` //TODO: unique
	Mock `json:"mock"`
}

type Mock struct {
	ApiName     string          `json:"apiName"bson:"apiName"`
	Key         string          `json:"key" bson:"key"`
	Value       string          `json:"value" bson:"value"`
	Reply       json.RawMessage `json:"reply" bson:"reply"`
	HandlerType string          `json:"handlerType" bson:"handlerType"`
	TimeOut     *int            `json:"timeOut,omitempty" bson:"timeOut"`
}
