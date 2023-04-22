package main

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
