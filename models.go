package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mock struct {
	ApiName     string      `json:"apiName"bson:"apiName"`
	Key         string      `json:"key" bson:"key"`
	Value       interface{} `json:"value" bson:"value"`
	Reply       string      `json:"reply" bson:"reply"`
	HandlerType string      `json:"handlerType" bson:"handlerType"`
	TimeOut     int         `json:"timeOut,omitempty" bson:"timeOut,omitempty"`
}

type KipiContext struct {
	*gin.Context
	*mongo.Client
}
