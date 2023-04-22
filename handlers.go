package main

import (
	. "github.com/go-kipi/kipimanager"
	"go.mongodb.org/mongo-driver/bson"
)

func (mocks getAllMocks) Run(c *KipiContext) (interface{}, ServiceError) {
	if res, err := c.Client.Database(mongo_DataBase).Collection(mongo_Collection).Find(c, bson.D{{}}, nil); err != nil {
		return nil, NewDbError("test", err)
	} else {
		var result []map[string]interface{}
		err = res.All(c, &result)
		if err != nil {
			return nil, NewIoError("Decode from mongo", err)

		}
		return result, nil
	}

}

func (mock createMock) Run(c *KipiContext) (interface{}, ServiceError) {
	mock.Id = "1234"
	if res, err := c.Client.Database(mongo_DataBase).Collection(mongo_Collection).InsertOne(c, mock, nil); err != nil {
		return nil, NewDbError("createMock", err)
	} else {
		return res, nil
	}

}

func (mock getMockById) Run(c *KipiContext) (interface{}, ServiceError) {
	var mockRes createMock
	filter := bson.M{
		"id": bson.M{
			"$eq": mock.Id,
		},
	}

	if err := c.Client.Database(mongo_DataBase).Collection(mongo_Collection).FindOne(c, filter, nil).Decode(&mockRes); err != nil {
		return nil, NewDbError("getMockById", err)
	} else {
		return mockRes, nil
	}

}

func (mock updateMockById) Run(c *KipiContext) (interface{}, ServiceError) {
	if res, err := c.Client.Database(mongo_DataBase).Collection(mongo_Collection).UpdateOne(c, mock, nil); err != nil {
		return nil, NewDbError("addRecord", err)
	} else {
		return res, nil
	}
}

func (mock deleteMockById) Run(c *KipiContext) (interface{}, ServiceError) {
	if res, err := c.Client.Database(mongo_DataBase).Collection(mongo_Collection).InsertOne(c, mock, nil); err != nil {
		return nil, NewDbError("addRecord", err)
	} else {
		return res, nil
	}
}
