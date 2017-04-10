package controller

import (
	"bitbucket.org/mirzaakhena/miranc-go/service"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

type AkunCtrl struct {
	AkunService *service.AkunService
}

func (ctrl AkunCtrl) CreateNewAkun(c *gin.Context) {
	var param service.CreateAkunParam
	if err := c.BindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": "error"})
		return
	}

	// usahaId := c.Param("usahaId")
	// userId := "123abc"
	// hakAkses := "CREATE_NEW_AKUN"

	ctrl.AkunService.CreateNewAkun(param)

	c.JSON(http.StatusCreated, map[string]string{"message": "created"})
}
