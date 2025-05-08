package routes

import (

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	AuthRoutes(r)
	FoodPrefrenceRoutes(r)
	UserPostRoutes(r)
	return r
}
