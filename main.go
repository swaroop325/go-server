package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

var response = Response{StatusCode: 200, Message: "Successful"}

func main() {
	//Initialising a Router
	ginEngine := gin.Default()
	//Creating a Route handler for GET request with anonymous function as handler function
	ginEngine.GET("/v1", v1)
	ginEngine.GET("/v2/:id", v2)
	//Serving the application on port number 8080
	ginEngine.Run()
}
func v1(context *gin.Context) {
	context.JSON(http.StatusOK, response)
}
func v2(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"statusCode": http.StatusFound, "message": "ID entered is" + context.Param("id")})
}
