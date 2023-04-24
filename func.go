package main

import (
	"encoding/json"
	"github.com/google/martian/v3/log"
)

//func getByFilter(c *KipiContext, filter interface{}, opt *options.FindOptions) ([]mongoVal, error) {
//	res, err := c.Client.Database(mongo_DataBase).Collection(mongo_Collection).Find(c, filter, opt)
//	if err != nil {
//		return nil, err
//	}
//
//	var result []mongoVal
//	err = res.All(c, &result)
//	if err != nil {
//		return nil, err
//
//	}
//	fmt.Println(result)
//	return result, nil
//}

func (mock Mock) MockReply() interface{} {
	var jsonReply = make(map[string]interface{})
	err := json.Unmarshal([]byte(mock.Reply), &jsonReply)
	if err != nil {
		log.Errorf("mock reply", err)
	}
	return jsonReply
}
