package usercontroller

import (
	"GoDemo/src/main/mapper/usermapper"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SearchUser(c *gin.Context) {
	username := c.Query("username")
	user, err := usermapper.FindUserByUsername(username)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusOK, user)
	c.Abort()
	return
}
