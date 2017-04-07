package controller

import (
	"bitbucket.org/mirzaakhena/seed-project/dto"
	"bitbucket.org/mirzaakhena/seed-project/service"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

type Barang struct {
	Service service.Barang
}

func (ctrl Barang) AddBarang(c *gin.Context) {
	var json dto.AddBarangParam
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": "error"})
		return
	}
	ctrl.Service.AddBarang(json)
	c.JSON(http.StatusCreated, map[string]string{"message": "created"})
}

func (ctrl Barang) GetAllBarang(c *gin.Context) {
	models := ctrl.Service.GetAllBarang()
	c.JSON(http.StatusOK, models)
}
