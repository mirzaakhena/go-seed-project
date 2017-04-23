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
	router, _ := MainEngine("barang.db")
	router.Run()
}

func MainEngine(databaseName string) (*gin.Engine, *gorm.DB) {

	// open db connection
	db, err := gorm.Open("sqlite3", databaseName)
	if err != nil {
		panic("gak bisa konek ke database")
	}

	db.LogMode(true)

	// // build table according to schema
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Akun{})
	db.AutoMigrate(&model.Usaha{})
	db.AutoMigrate(&model.UserUsaha{})
	db.AutoMigrate(&model.SubAkun{})
	db.AutoMigrate(&model.Jurnal{})

	// wiring "bean"
	userService := service.UserService{DB: db}
	userRest := rest.UserRest{UserService: &userService}

	usahaService := service.UsahaService{DB: db}
	usahaRest := rest.UsahaRest{UsahaService: &usahaService}

	akunService := service.AkunService{DB: db}
	akunRest := rest.AkunRest{AkunService: &akunService}

	subakunService := service.SubAkunService{DB: db}
	subakunRest := rest.SubAkunRest{SubAkunService: &subakunService}

	jurnalService := service.JurnalService{
		DB:             db,
		AkunService:    &akunService,
		SubAkunService: &subakunService,
	}
	jurnalRest := rest.JurnalRest{JurnalService: &jurnalService}

	// prepare endpoint api
	router := gin.Default()

	// endpoints
	router.POST("/register", userRest.Register)
	router.POST("/login", userRest.Login)

	authorized := router.Group("/usaha")

	authorized.Use(rest.Authenticate)
	{

		authorized.POST("", usahaRest.CreateUsaha)
		authorized.GET("", usahaRest.GetAllUsahaByUser)

		authorized.POST("/:usahaId/akun", akunRest.Create)
		authorized.GET("/:usahaId/akun", akunRest.GetAll)

		authorized.POST("/:usahaId/akun/:akunId", subakunRest.CreateSubAkun)
		authorized.GET("/:usahaId/akun/:akunId", subakunRest.GetAllSubAkun)

		authorized.POST("/:usahaId/jurnal", jurnalRest.CreateJurnal)
	}

	return router, db
}
