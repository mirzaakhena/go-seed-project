package rest

import (
	"bitbucket.org/mirzaakhena/miranc-go/service"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

type SubAkunRest struct {
	SubAkunService *service.SubAkunService
}

func (ctrl SubAkunRest) CreateSubAkun(c *gin.Context) {
	var param service.CreateSubAkunParam
	if err := c.BindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid params"})
		return
	}

	usahaId := c.Param("usahaId")
	akunId := c.Param("akunId")

	err := ctrl.SubAkunService.CreateSubAkun(usahaId, akunId, param)
	if err != nil {
		c.JSON(http.StatusCreated, map[string]string{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, map[string]string{"message": "subakun berhasil dibuat"})

}

func (ctrl SubAkunRest) GetAllSubAkun(c *gin.Context) {

	usahaId := c.Param("usahaId")
	akunId := c.Param("akunId")

	c.JSON(http.StatusOK, ctrl.SubAkunService.GetAllSubAkun(usahaId, akunId))
}
