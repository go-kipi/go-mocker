package reply

import (
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
	serviceReply := ServiceReply{}
	serviceReply.Status = "success"
	serviceReply.Data = reply
	c.JSON(http.StatusOK, serviceReply)
}
