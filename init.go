package main

import (
	. "github.com/go-kipi/kipimanager"
	cors "github.com/itsjamie/gin-cors"
	"time"
)

func initRouters(router *KipiRouter) {
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	router.POST("/getAllMocks", Handle(&getAllMocks{}))
	router.POST("/getMockById", Handle(&getMockById{}))
	router.POST("/updateMockById", Handle(&updateMockById{}))
	router.POST("/deleteMockById", Handle(&deleteMockById{}))

	router.POST("/createMock", Handle(&createMock{}))
	//router.Use(spa.Middleware("/", "./dist"))

}
