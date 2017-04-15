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

	user, err := ctrl.UserService.Register(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)

}

func (ctrl UserRest) Login(c *gin.Context) {
	var param service.LoginParam
	if err := c.BindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid params"})
		return
	}

	token, err := ctrl.UserService.Login(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	c.Writer.Header().Add("token", token)

	c.JSON(http.StatusOK, map[string]string{"message": "user berhasil login"})
}

func (ctrl UserRest) Invite(c *gin.Context) {
	v, _ := c.Get("jwt")
	jwt, _ := v.(*service.CustomJwt)
	c.JSON(http.StatusOK, map[string]string{"message": jwt.UserId})
}
