package rest

import (
	"bitbucket.org/mirzaakhena/miranc-go/service"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

type AkunRest struct {
	AkunService *service.AkunService
}

func (ctrl AkunRest) CreateAkun(c *gin.Context) {
	var param service.CreateAkunParam
	if err := c.BindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid params"})
		return
	}

	usahaId := c.Param("usahaId")

	akun, err := ctrl.AkunService.CreateAkun(usahaId, param)
	if err != nil {
		c.JSON(http.StatusCreated, map[string]string{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, akun)

}

func (ctrl AkunRest) GetAllAkun(c *gin.Context) {

	usahaId := c.Param("usahaId")

	c.JSON(http.StatusOK, ctrl.AkunService.GetAllAkun(usahaId))
}
