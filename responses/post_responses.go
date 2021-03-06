package responses

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostResponse struct {
}

func (postResponse PostResponse) Postv1(context *gin.Context) {
	var cust Customer
	// With ShouldBind function, the request parameters are bound to the struct object.
	if err := context.ShouldBind(&cust); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": fmt.Sprintf("%v", err)})
	} else {
		context.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "custId": cust.CustId, "custName": "customer id is: " + cust.CustName})
	}
}
