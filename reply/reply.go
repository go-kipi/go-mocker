package reply

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ServiceReply struct {
	BaseServiceReply
	Data interface{} `json:"data,omitempty"`
}

type BaseErrorReply struct {
	Msg  string `json:"msg"`
	Type string `json:"type"`
}

type BaseServiceReply struct {
	Status string          `json:"status"`
	Error  *BaseErrorReply `json:"error,omitempty"`
}

func SuccessReply(c *gin.Context, reply interface{}) {
	fmt.Println("success:", reply)
	serviceReply := ServiceReply{}
	serviceReply.Status = "success"
	serviceReply.Data = reply
	c.JSON(http.StatusOK, serviceReply)
}

func ErrorReply(c *gin.Context, msg string, err error) {
	fmt.Println("error: ", msg, err)
	serviceReply := ServiceReply{}
	serviceReply.Status = "error"
	serviceReply.Data = err
	c.JSON(http.StatusOK, serviceReply)
}
