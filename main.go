package main

import (
	"bitbucket.org/mirzaakhena/miranc-go/model"
	"bitbucket.org/mirzaakhena/miranc-go/rest"
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

	db.LogMode(true)

	// // build table according to schema
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Akun{})
	db.AutoMigrate(&model.Usaha{})
	db.AutoMigrate(&model.UserUsaha{})

	// wiring "bean"
	userService := service.UserService{DB: db}
	userRest := rest.UserRest{UserService: &userService}

	akunService := service.AkunService{DB: db}
	akunRest := rest.AkunRest{AkunService: &akunService}

	usahaService := service.UsahaService{DB: db}
	usahaRest := rest.UsahaRest{UsahaService: &usahaService}

	// prepare endpoint api
	router := gin.Default()

	// endpoints
	router.POST("/register", userRest.Register)
	router.POST("/login", userRest.Login)

	authorized := router.Group("/usaha")

	authorized.Use(rest.Authenticate)
	{

		authorized.POST("/", usahaRest.CreateUsaha)
		authorized.GET("/", usahaRest.GetAllUsahaByUser)

		authorized.POST("/:usahaId/invite", userRest.Invite)
		authorized.POST("/:usahaId/akun", akunRest.CreateNewAkun)

	}

	// start server
	router.Run()

}
