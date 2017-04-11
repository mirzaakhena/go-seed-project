package rest

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func Authenticate(c *gin.Context) {
	c.JSON(http.StatusCreated, map[string]string{"message": "okee"})
}
