package routes

import (
	"net/http"
	user "go-rest-with-gin/controllers/users"
	auth "go-rest-with-gin/controllers/auth"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	router := gin.Default()
	api := router.Group("/api") 
	{
		api.GET("/users", user.GetAllUsers)
		api.POST("/register", user.RegisterUser)
		api.GET("/users/:id", user.GetUserById)
		api.POST("/login", auth.Login)
		api.PUT("/users/:id", user.UpdateUser)
		api.DELETE("/users/:id", user.DeleteUser)
	}
	router.NoRoute(func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "app is running"})
	})
	router.Run(":8002")
}