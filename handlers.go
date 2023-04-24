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
		//return nil, NewDbError("test", err)
	} else {
		var result []map[string]interface{}
		err = res.All(c, &result)
		if err != nil {
			//return nil, NewIoError("Decode from mongo", err)
		}
		reply.SuccessReply(c, result)
	}
}

func createMock(c *gin.Context) {
	var mock Mock
	if err := c.ShouldBindJSON(&mock); err != nil {
		c.JSON(400, gin.H{
			"status": "error",
			"error":  err,
		})
		return
	}
	if res, err := mongoDb.Database(db.Mongo_DataBase).Collection(db.Mongo_Collection).InsertOne(c, Mock{
		ApiName:     mock.ApiName,
		Key:         mock.Key,
		Value:       mock.Value,
		Reply:       mock.Reply,
		HandlerType: mock.HandlerType,
		TimeOut:     mock.TimeOut,
	}, nil); err != nil {
		//return nil, NewDbError("createMock", err)
	} else {
		reply.SuccessReply(c, res)
	}

}

func getMockById(c *gin.Context) {
	var reqData = make(map[string]string)
	if err := c.ShouldBindJSON(&reqData); err != nil {
		c.JSON(400, gin.H{
			"status": "error",
			"error":  err,
		})
		return
	}
	objectId, err := primitive.ObjectIDFromHex(reqData["id"])
	if err != nil {
	}
	var mockRes Mock

	filter := bson.M{"_id": objectId}

	if err := mongoDb.Database(db.Mongo_DataBase).Collection(db.Mongo_Collection).FindOne(c, filter, nil).Decode(&mockRes); err != nil {
	} else {
		reply.SuccessReply(c, mockRes)
	}

}

//	func (mock updateMockById) Run(c *KipiContext) (interface{}, ServiceError) {
//		if res, err := c.Client.Database(mongo_DataBase).Collection(mongo_Collection).UpdateOne(c, mock, nil); err != nil {
//			return nil, NewDbError("addRecord", err)
//		} else {
//			return res, nil
//		}
//	}
//
//	func (mock deleteMockById) Run(c *KipiContext) (interface{}, ServiceError) {
//		if res, err := c.Client.Database(mongo_DataBase).Collection(mongo_Collection).InsertOne(c, mock, nil); err != nil {
//			return nil, NewDbError("addRecord", err)
//		} else {
//			return res, nil
//		}
//	}
//
//	func (mock dynamicApi) Run(c *KipiContext) (interface{}, ServiceError) {
//		param := c.Param("api")
//		c.JSON(200, gin.H{
//			"ttttt": param,
//		})
//
//		return nil, nil
//	}
func dynamicApi(c *gin.Context) {
	var reqData = make(map[string]interface{})
	if err := c.ShouldBindJSON(&reqData); err != nil {
		c.JSON(400, gin.H{
			"status": "error",
			"error":  err,
		})
		return
	}
	param := c.Param("api")

	filter := bson.M{"apiName": param}

	if res, err := mongoDb.Database(db.Mongo_DataBase).Collection(db.Mongo_Collection).Find(c, filter, nil); err != nil {
		//error
	} else {
		var mockRes []Mock
		res.All(c, &mockRes)
		for _, mock := range mockRes {
			if reqData[mock.Key] == mock.Value {
				time.Sleep(time.Duration(mock.TimeOut) * time.Second)
				c.JSON(200, mock.Reply)
			}
		}
		//for _, v := range reqList {
		//	if strings.Contains(v.Endpoint, dynamicEndPointParam) {
		//		if reqData[v.Key] == v.Value {
		//			if res, err := getContentFile(v.ResponseID); err != nil {
		//				c.JSON(400, gin.H{
		//					"status": "error",
		//					"error":  err,
		//				})
		//				return
		//			} else {
		//				time.Sleep(time.Duration(v.TimeOut) * time.Second)
		//				fmt.Println(res)
		//				c.JSON(200, res)
		//				return
		//			}
		//
		//		}
		//	} else {
		//		c.JSON(404, gin.H{
		//			"status": "error",
		//			"error":  "page not found",
		//		})
		//	}
		//}

		//c.JSON(200, mockRes.Reply)
	}

}
