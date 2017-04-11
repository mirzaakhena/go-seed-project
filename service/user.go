package service

import (
	"bitbucket.org/mirzaakhena/miranc-go/model"
	_ "errors"
	"github.com/jinzhu/gorm"
	_ "github.com/satori/go.uuid"
	_ "strconv"
)

type UserService struct {
	DB *gorm.DB
}

type RegisterParam struct {
	Nama     string `json:"nama" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Telepon  string `json:"telepon"`
	Alamat   string `json:"alamat"`
}

func (serv UserService) Register(param RegisterParam) error {

	var count int
	serv.DB.Model(&model.User{}).Where("email = ?", param.Email).Count(&count)
	log.Debugf("count => %d", count)
	if count > 0 {
		s := "User dengan email " + param.Email + " sudah terdaftar"
		log.Errorf(s)
		return errors.New(s)
	}

	serv.DB.Create(&model.User{
		ID:       uuid.NewV4().String(),
		Nama:     param.Nama,
		Email:    param.Email,
		Password: param.Password,
		Telepon:  param.Telepon,
		Alamat:   param.Alamat,
	})

	return nil
}

func (serv UserService) Login(param RegisterParam) error {
	return nil
}

func (serv UserService) Invite(usahaId string, param RegisterParam) error {
	return nil
}
