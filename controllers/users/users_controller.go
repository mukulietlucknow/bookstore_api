package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser for creatinf the user
func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement Me!")
}

// GetUser this will ve getting the usr
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement Me!")
}

// SearchUser it will be for finding the user
// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implement Me!")
// }
