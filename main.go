package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

var mongoDb *mongo.Client

func main() {
	r := initRouters()
	r.Run(":" + getPort())
}
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9999"
	}
	return port
}
