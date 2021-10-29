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

func v1(context *gin.Context) {
	context.JSON(http.StatusOK, response)
}
func v2(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"statusCode": http.StatusFound, "message": "ID entered is" + context.Param("id")})
}
func v3(context *gin.Context) {
	var custId = context.Query("id")
	// fetches the value of custname parameter which of string type, if it does not exist it return empty string
	var custName = context.Query("custname")
	context.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK,
		"Customer Id": custId, "Customer Name": custName})
}
