package rest

import (
	"bitbucket.org/mirzaakhena/miranc-go/service"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"strings"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	log.Debug(token)
	s := strings.Fields(token)
	if len(s) != 2 {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	if s[0] != "Bearer" {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	jwt, err := service.ValidateToken(s[1])
	if err != nil {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	c.Set("jwt", jwt)
}
