package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var response = Response{StatusCode: 200, Message: "Successful"}

func getv1(context *gin.Context) {
	context.JSON(http.StatusOK, response)
}
func getv2(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"statusCode": http.StatusFound, "message": "ID entered is" + context.Param("id")})
}
func getv3(context *gin.Context) {
	var custId = context.Query("id")
	// fetches the value of custname parameter which of string type, if it does not exist it return empty string
	var custName = context.Query("custname")
	context.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK,
		"Customer Id": custId, "Customer Name": custName})
}
