package responses

type Response struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type Customer struct {
	CustId   int    `json:"custId" binding:"required"`
	CustName string `json:"custName" binding:"required"`
}

type Plan struct {
	PlanId   int
	PlanName string
}
type Friends struct {
	Friend []int
}
type Customer1 struct {
	CustId   int    `json:"custId" binding:"required"`
	CustName string `json:"custName" binding:"required"`
	Plan
	Friends
}
