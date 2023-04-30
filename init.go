package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-kipi/go-mocker/db"
	cors "github.com/itsjamie/gin-cors"
	"github.com/mandrigin/gin-spa/spa"
	"os"
	"time"
)

func initRouters() *gin.Engine {
	router := gin.New()
	router.TrustedPlatform = gin.PlatformGoogleAppEngine
	router.TrustedPlatform = "X-CDN-IP"
	router.SetTrustedProxies([]string{})
	router.RemoteIPHeaders = []string{"X-Forwarded-For", "X-Real-IP"}
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	mongoStr := os.Getenv("mongo")
	if mongoStr != "" {
		mongoDb = db.InitMongo(mongoStr)
	} else {
		panic("mongo env is missing")
	}

	router.POST("/getAllMocks", getAllMocks)
	router.POST("/getMockById", getMockById)
	router.POST("/updateMockById", updateMockById)
	router.POST("/deleteMockById", deleteMockById)
	router.POST("/createMock", createMock)
	router.POST("/reply/:api", dynamicApi)

	router.GET("/ip", getIp)
	router.Use(spa.Middleware("/", "./dist"))
	return router
}
