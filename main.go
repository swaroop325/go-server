package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//Initialising a Router
	ginEngine := gin.Default()
	//Creating a Route handler for GET request with anonymous function as handler function

	//GET Requests
	//http://localhost:8080/v1
	ginEngine.GET("/v1", getv1)

	//http://localhost:8080/v2/200
	ginEngine.GET("/v2/:id", getv2)

	//http://localhost:8080/v3?custname=300&&id=1
	ginEngine.GET("/v3", getv3)

	//POST Reqquests
	ginEngine.POST("/v1", postv1)

	//XML Rendereing
	ginEngine.GET("/xml", xmlResponse)

	//Serving the application on port number 8080
	ginEngine.Run()
}
