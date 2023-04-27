# go-mocker

url: https://go-mocker.onrender.com
## Api
POST   /getAllMocks
POST   /getMockById
POST   /updateMockById
POST   /deleteMockById
POST   /createMock
POST   /reply/:api
GET    /ip

Model:
```
{
    id          string      
    apiName     string      
    key         string      
    value       any 
    reply       string      
    handlerType string      
    timeOut     int         
}
```

createMock && updateMockById
```
{
    id          string      
    apiName     string      
    key         string      
    value       any 
    reply       string      
    handlerType string      
    timeOut     int         
}
```

```
updateMockById && deleteMockById
{
    id          string
}
```
