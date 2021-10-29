package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//Initialising a Router
	ginEngine := gin.Default()
	//Creating a Route handler for GET request with anonymous function as handler function
	
	//http://localhost:8080/v1
	ginEngine.GET("/v1", v1)

	//http://localhost:8080/v2/200
	ginEngine.GET("/v2/:id", v2)

	//http://localhost:8080/v3?custname=300&&id=1
	ginEngine.GET("/v3", v3)
	
	//Serving the application on port number 8080
	ginEngine.Run()
}
