package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-kipi/go-mocker/db"
	"github.com/go-kipi/go-mocker/reply"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func getAllMocks(c *gin.Context) {
	if mock, err := getAllMocksFromDb(c); err != nil {
		reply.ErrorReply(c, "getAllMocks - getAllMocksFromDb", err)
	} else {
		reply.SuccessReply(c, mock)
	}
}

func createMock(c *gin.Context) {
	var mock Mock
	if err := c.ShouldBindJSON(&mock); err != nil {
		reply.ErrorReply(c, "createMock - ShouldBindJSON", err)
		return
	}

	isValid, err := validateApiNameKeyValue(c, mock)
	if err != nil {
		reply.ErrorReply(c, "createMock - validateApiNameKeyValue", err)
		return
	}
	if isValid {
		if res, err := mongoDb.Database(db.Mongo_DataBase).Collection(db.Mongo_Collection).InsertOne(c, Mock{
			ApiName:     mock.ApiName,
			Key:         mock.Key,
			Value:       mock.Value,
			Reply:       mock.Reply,
			HandlerType: mock.HandlerType,
			TimeOut:     mock.TimeOut,
		}, nil); err != nil {
			reply.ErrorReply(c, "createMock - InsertOne", err)
			return
		} else {
			reply.SuccessReply(c, res)
		}
	} else {
		reply.ErrorUserReply(c, fmt.Sprintf("%s,%s,%v are already exist", mock.ApiName, mock.Key, mock.Value), err)
		return
	}

}

func getMockById(c *gin.Context) {
	var reqData = make(map[string]string)
	if err := c.ShouldBindJSON(&reqData); err != nil {
		reply.ErrorReply(c, "getMockById - ShouldBindJSON", err)
		return
	}
	objectId, err := primitive.ObjectIDFromHex(reqData["id"])
	if err != nil {
		reply.ErrorReply(c, "getMockById - primitive.ObjectIDFromHex", err)
		return
	}
	var mockRes Mock

	filter := bson.M{"_id": objectId}

	res := mongoDb.Database(db.Mongo_DataBase).Collection(db.Mongo_Collection).FindOne(c, filter, nil)
	if res.Err() != nil { //check if empty
		reply.ErrorReply(c, "getMockById - FindOne", res.Err())
		return
	}
	err = res.Decode(&mockRes)
	if res.Err() != nil {
		reply.ErrorReply(c, "getMockById - Decode", err)
		return
	}
	reply.SuccessReply(c, mockRes)
}

func updateMockById(c *gin.Context) {
	var mock Mock
	if err := c.ShouldBindJSON(&mock); err != nil {
		reply.ErrorReply(c, "updateMockById - ShouldBindJSON", err)
		return
	}

	objectId, err := primitive.ObjectIDFromHex(mock.Id)
	if err != nil {
		reply.ErrorReply(c, "updateMockById - primitive.ObjectIDFromHex", err)
		return
	}

	filter := bson.M{"_id": objectId}

	update := bson.D{{"$set", Mock{
		ApiName:     mock.ApiName,
		Key:         mock.Key,
		Value:       mock.Value,
		Reply:       mock.Reply,
		HandlerType: mock.HandlerType,
	}}}

	if res, err := mongoDb.Database(db.Mongo_DataBase).Collection(db.Mongo_Collection).UpdateOne(c, filter, update, nil); err != nil {
		reply.ErrorReply(c, "updateMockById - UpdateOne", err)
	} else {
		reply.SuccessReply(c, res)
	}
}

func deleteMockById(c *gin.Context) {
	var reqData = make(map[string]string)
	if err := c.ShouldBindJSON(&reqData); err != nil {
		reply.ErrorReply(c, "deleteMockById - ShouldBindJSON", err)
		return
	}
	objectId, err := primitive.ObjectIDFromHex(reqData["id"])
	if err != nil {
		reply.ErrorReply(c, "deleteMockById - primitive.ObjectIDFromHex", err)
		return
	}

	filter := bson.M{"_id": objectId}

	if res, err := mongoDb.Database(db.Mongo_DataBase).Collection(db.Mongo_Collection).DeleteOne(c, filter, nil); err != nil {
		reply.ErrorReply(c, "deleteMockById - DeleteOne", err)
	} else {
		reply.SuccessReply(c, res)
	}
}

func dynamicApi(c *gin.Context) {
	var reqData = make(map[string]interface{})
	if err := c.ShouldBindJSON(&reqData); err != nil {
		reply.ErrorReply(c, "dynamicApi - ShouldBindJSON", err)
		return
	}
	param := c.Param("api")

	filter := bson.M{"apiName": param}

	if res, err := mongoDb.Database(db.Mongo_DataBase).Collection(db.Mongo_Collection).Find(c, filter, nil); err != nil {
		reply.ErrorReply(c, "dynamicApi - Find", err)
		return
	} else {
		var mockRes []Mock
		if err := res.All(c, &mockRes); err != nil {
			reply.ErrorReply(c, "dynamicApi - All", err)
			return
		}
		for _, mock := range mockRes {
			if reqData[mock.Key] == mock.Value {
				time.Sleep(time.Duration(mock.TimeOut) * time.Second)
				c.JSON(200, mock.mockReply())
				return
			}
		}
	}
}

func getIp(c *gin.Context) {
	var ips []string
	for _, h := range []string{XForwardedFor, XRealIp} {
		ip := c.Request.Header.Get(h)
		fmt.Println(ip)
		ips = append(ips, ip)
	}

	reply.SuccessReply(c, ips)
}
