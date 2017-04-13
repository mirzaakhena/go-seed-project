package rest

import (
	"bitbucket.org/mirzaakhena/miranc-go/service"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

type UsahaRest struct {
	UsahaService *service.UsahaService
}

func (rest UsahaRest) CreateUsaha(c *gin.Context) {
	var param service.CreateUsahaParam
	if err := c.BindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid params"})
		return
	}
	v, _ := c.Get("jwt")
	jwt, _ := v.(*service.CustomJwt)
	rest.UsahaService.CreateUsaha(jwt.UserId, param)
	c.JSON(http.StatusOK, map[string]string{"message": "Usaha telah dibuat"})
}

func (rest UsahaRest) GetAllUsahaByUser(c *gin.Context) {
	v, _ := c.Get("jwt")
	jwt, _ := v.(*service.CustomJwt)
	c.JSON(http.StatusOK, rest.UsahaService.GetAllUsahaByUser(jwt.UserId))
}
