package main

import (
	"bitbucket.org/mirzaakhena/go-seed-project/controller"
	"bitbucket.org/mirzaakhena/go-seed-project/model"
	"bitbucket.org/mirzaakhena/go-seed-project/repo"
	"bitbucket.org/mirzaakhena/go-seed-project/service"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/satori/go.uuid"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {

	// open db connection
	db, err := gorm.Open("sqlite3", "barang.db")
	if err != nil {
		panic("gak bisa konek ke database")
	}

	// build table according to schema
	db.AutoMigrate(&model.Barang{})

	// wiring "bean"
	barangRepo := repo.Barang{DB: db}
	barangService := service.Barang{Repo: barangRepo}
	barangController := controller.Barang{Service: barangService}

	// prepare endpoint api
	router := gin.Default()
	router.GET("/barang", barangController.GetAllBarang)
	router.POST("/barang", barangController.AddBarang)

	// start server
	router.Run()

}
