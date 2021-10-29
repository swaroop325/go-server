package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type OtherResponse struct {
}

func (otherResponse OtherResponse) XmlResponse(context *gin.Context) {
	var cust Customer
	cust.CustId = 123
	cust.CustName = "Alex"
	context.XML(http.StatusOK, cust)
}

func (otherResponse OtherResponse) YamlResponse(context *gin.Context) {
	cust1 := Customer1{
		CustId:   101,
		CustName: "Alex",
		Plan: Plan{
			PlanId:   10,
			PlanName: "ISD",
		},
		Friends: Friends{
			Friend: []int{10, 20},
		},
	}
	context.YAML(http.StatusOK, cust1)
}
