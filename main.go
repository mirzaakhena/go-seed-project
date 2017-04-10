package main

import (
	"bitbucket.org/mirzaakhena/miranc-go/controller"
	"bitbucket.org/mirzaakhena/miranc-go/model"
	"bitbucket.org/mirzaakhena/miranc-go/service"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {

	// open db connection
	db, err := gorm.Open("sqlite3", "barang.db")
	if err != nil {
		panic("gak bisa konek ke database")
	}

	// // build table according to schema
	// // db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Akun{})

	// wiring "bean"
	akunService := service.AkunService{DB: db}
	akunController := controller.AkunCtrl{AkunService: &akunService}

	// prepare endpoint api
	router := gin.Default()

	// endpoints
	router.POST("/:usahaId/akun", akunController.CreateNewAkun)

	// start server
	router.Run()

}
