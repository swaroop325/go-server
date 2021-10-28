package main
import (
    "net/http"
    "github.com/gin-gonic/gin"
)
func main() {
    //Initialising a Router
    ginEngine := gin.Default()
    //Creating a Route handler for GET request with anonymous function as handler function
    ginEngine.GET("/test", hello)
    //Serving the application on port number 8080
    ginEngine.Run()
}
func hello(context *gin.Context) {
    context.String(http.StatusOK, "Hello World!")
}