package main

import (
	"api_server/responses"

	"github.com/gin-gonic/gin"
)

var (
	getResponse   = responses.GetResponse{}
	postResponse  = responses.PostResponse{}
	otherResponse = responses.OtherResponse{}
)

func main() {
	//Initialising a Router
	ginEngine := gin.Default()
	//Creating a Route handler for GET request with anonymous function as handler function

	//GET Requests
	//http://localhost:8080/v1
	ginEngine.GET("/v1", getResponse.Getv1)

	//http://localhost:8080/v2/200
	ginEngine.GET("/v2/:id", getResponse.Getv2)

	//http://localhost:8080/v3?custname=300&&id=1
	ginEngine.GET("/v3", getResponse.Getv3)

	//http://localhost:8080/v3?custname=300&&id=1
	ginEngine.GET("/getApi", getResponse.GetApi)

	//POST Reqquests
	ginEngine.POST("/v1", postResponse.Postv1)

	//XML Rendereing
	ginEngine.GET("/xml", otherResponse.XmlResponse)

	//YAML response
	ginEngine.GET("/yaml", otherResponse.YamlResponse)
	//Serving the application on port number 8080
	ginEngine.Run()
}
