package rest

import (
	"bitbucket.org/mirzaakhena/miranc-go/service"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

type UserRest struct {
	UserService *service.UserService
}

func (ctrl UserRest) Register(c *gin.Context) {
	var param service.RegisterParam
	if err := c.BindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid params"})
		return
	}

	err := ctrl.UserService.Register(param)
	if err != nil {
		c.JSON(http.StatusCreated, map[string]string{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, map[string]string{"message": "user berhasil didaftarkan"})

}

func (ctrl UserRest) Login(c *gin.Context) {

}

func (ctrl UserRest) Invite(c *gin.Context) {

}
