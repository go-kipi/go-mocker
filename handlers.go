package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kipi/go-mocker/db"
	"github.com/go-kipi/go-mocker/reply"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func getAllMocks(c *gin.Context) {
	if res, err := mongoDb.Database(db.Mongo_DataBase).Collection(db.Mongo_Collection).Find(c, bson.D{{}}, nil); err != nil {
		reply.ErrorReply(c, "getAllMocks - Find", err)
	} else {
		var result []map[string]interface{}
		err = res.All(c, &result)
		if err != nil {
			reply.ErrorReply(c, "getAllMocks - All", err)
		}
		reply.SuccessReply(c, result)
	}
}

func createMock(c *gin.Context) {
	var mock Mock
	if err := c.ShouldBindJSON(&mock); err != nil {
		reply.ErrorReply(c, "createMock - ShouldBindJSON", err)
	}
	if res, err := mongoDb.Database(db.Mongo_DataBase).Collection(db.Mongo_Collection).InsertOne(c, Mock{
		ApiName:     mock.ApiName,
		Key:         mock.Key,
		Value:       mock.Value,
		Reply:       mock.Reply,
		HandlerType: mock.HandlerType,
		TimeOut:     mock.TimeOut,
	}, nil); err != nil {
		reply.ErrorReply(c, "createMock - InsertOne", err)
	} else {
		reply.SuccessReply(c, res)
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
	}
	var mockRes Mock

	filter := bson.M{"_id": objectId}

	if err := mongoDb.Database(db.Mongo_DataBase).Collection(db.Mongo_Collection).FindOne(c, filter, nil).Decode(&mockRes); err != nil {
		reply.ErrorReply(c, "getMockById - FindOne", err)
	} else {
		reply.SuccessReply(c, mockRes)
	}

}

func updateMockById(c *gin.Context) {
	var mock updateMock
	if err := c.ShouldBindJSON(&mock); err != nil {
		reply.ErrorReply(c, "updateMockById - ShouldBindJSON", err)
	}

	objectId, err := primitive.ObjectIDFromHex(mock.Id)
	if err != nil {
		reply.ErrorReply(c, "updateMockById - primitive.ObjectIDFromHex", err)
	}

	filter := bson.M{"_id": objectId}
	if res, err := mongoDb.Database(db.Mongo_DataBase).Collection(db.Mongo_Collection).UpdateOne(c, filter, mock, nil); err != nil {
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
	}
	param := c.Param("api")

	filter := bson.M{"apiName": param}

	if res, err := mongoDb.Database(db.Mongo_DataBase).Collection(db.Mongo_Collection).Find(c, filter, nil); err != nil {
		reply.ErrorReply(c, "dynamicApi - Find", err)
	} else {
		var mockRes []Mock
		if err := res.All(c, &mockRes); err != nil {
			reply.ErrorReply(c, "dynamicApi - All", err)
		}
		for _, mock := range mockRes {
			if reqData[mock.Key] == mock.Value {
				time.Sleep(time.Duration(mock.TimeOut) * time.Second)
				c.JSON(200, mock.Reply)
			}
		}
	}
}
