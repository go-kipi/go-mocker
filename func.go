package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kipi/go-mocker/db"
	"go.mongodb.org/mongo-driver/bson"
)

func validateApiNameKeyValue(c *gin.Context, mock Mock) (bool, error) {
	//TODO: replace in mongo count query
	// count > 0 return error
	if mockes, err := getAllMocksFromDb(c); err != nil {
		return false, err
	} else {
		for _, v := range mockes {
			if v.HandlerType == mock.HandlerType && v.Key == mock.Key && v.Value == mock.Value {
				return false, nil
			}
		}
	}
	return true, nil
}

func getAllMocksFromDb(c *gin.Context) ([]Mock, error) {
	var mock []Mock
	if res, err := mongoDb.Database(db.Mongo_DataBase).Collection(db.Mongo_Collection).Find(c, bson.D{{}}, nil); err != nil {
		return mock, err
	} else {
		err = res.All(c, &mock)
		if err != nil {
			return mock, err
		}
		return mock, nil
	}
}
