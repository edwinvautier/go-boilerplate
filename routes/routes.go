package routes

import (
	"github.com/edwinvautier/go-boilerplate/controllers"
	"github.com/gin-gonic/gin"
)

// Init initializes router with the following routes
func Init(r *gin.Engine) {
	// api := r.Group("/api")

	r.POST("/register", controllers.CreateCustomer)
}