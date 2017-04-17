package rest

import (
	"bitbucket.org/mirzaakhena/miranc-go/service"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

type JurnalRest struct {
	JurnalService *service.JurnalService
}

func (ctrl JurnalRest) CreateJurnal(c *gin.Context) {

	var param service.CreateJurnalParam
	if err := c.BindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid params"})
		return
	}

	usahaId := c.Param("usahaId")

	v, _ := c.Get("jwt")

	jwt, _ := v.(*service.CustomJwt)

	jurnal, err := ctrl.JurnalService.CreateJurnal(usahaId, jwt.UserId, param)
	if err != nil {
		c.JSON(http.StatusCreated, map[string]string{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, jurnal)

}
