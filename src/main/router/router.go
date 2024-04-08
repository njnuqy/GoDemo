package router

import (
	"GoDemo/src/main/controller/usercontroller"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	UserGroup := r.Group("/User")
	{
		UserGroup.GET("/SearchUser", usercontroller.SearchUser)
	}
}
