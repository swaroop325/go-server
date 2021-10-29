package responses

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var response = Response{StatusCode: 200, Message: "Successful"}

type GetResponse struct {
}

func (getResponse GetResponse) Getv1(context *gin.Context) {
	context.JSON(http.StatusOK, response)
}
func (getResponse GetResponse) Getv2(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"statusCode": http.StatusFound, "message": "ID entered is" + context.Param("id")})
}
func (getResponse GetResponse) Getv3(context *gin.Context) {
	var custId = context.Query("id")
	// fetches the value of custname parameter which of string type, if it does not exist it return empty string
	var custName = context.Query("custname")
	context.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK,
		"Customer Id": custId, "Customer Name": custName})
}

func (getResponse GetResponse) GetApi(context *gin.Context) {

	url := "http://localhost:3001/v2/100"
	//Invoking ‘Get’ method available from the ‘net/http’ package
	response, err := http.Get(url)
	//Handling Error in case the request url is not valid
	if err != nil {
		//Logging Error Information
		log.Fatal(err)
	}
	//’ReadAll’ method from ‘io’ package is used to read the Response Body
	// body, err := io.ReadAll(response.Body)
	//Closing the connection to the request once the response is read
	response.Body.Close()
	//Handling any error while reading the response
	if err != nil {
		//Logging Error Information
		log.Fatal(err)
	}
	fmt.Println(response.StatusCode)
	context.JSON(http.StatusOK, gin.H{"response": response.StatusCode})
}
