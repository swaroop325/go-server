package main

type Response struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type Customer struct {
	CustId   int    `json:"custId" binding:"required"`
	CustName string `json:"custName" binding:"required"`
}