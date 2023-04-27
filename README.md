# go-mocker

url: https://go-mocker.onrender.com
## Api
```
POST   /getAllMocks
POST   /getMockById
POST   /updateMockById
POST   /deleteMockById
POST   /createMock
POST   /reply/:api
GET    /ip
```
Model:
```
{
    id          string      
    apiName     string      
    reqKey      string      
    reqValue    any 
    reply       string      
    handlerType string      //api
    timeOut     int         
}
```

createMock && updateMockById
```
{
    id          string      
    apiName     string      
    reqKey      string      
    reqValue    any 
    reply       string      
    handlerType string      
    timeOut     int         
}
```

```
getMockById && updateMockById && deleteMockById
{
    id          string
}
```
