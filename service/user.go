package service

import (
	"bitbucket.org/mirzaakhena/miranc-go/model"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	_ "strconv"
)

type UserService struct {
	DB *gorm.DB
}

type RegisterParam struct {
	Nama     string `json:"nama" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

type LoginParam struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (serv UserService) Register(param RegisterParam) error {

	var count int

	tx := serv.DB.Begin()

	tx.Model(&model.User{}).Where("email = ?", param.Email).Count(&count)

	if count > 0 {
		s := "User dengan email " + param.Email + " sudah terdaftar"
		log.Errorf(s)
		return errors.New(s)
	}

	tx.Create(&model.User{
		ID:       uuid.NewV4().String(),
		Nama:     param.Nama,
		Email:    param.Email,
		Password: param.Password,
		Phone:    param.Phone,
		Address:  param.Address,
	})

	tx.Commit()

	log.Infof("user %s berhasil didaftarkan", param.Nama)

	return nil
}

func (serv UserService) Login(param LoginParam) (string, error) {

	var user model.User
	serv.DB.Where("email = ? AND password = ?", param.Email, param.Password).First(&user)
	if user.ID == "" {
		s := "email atau password tidak cocok"
		log.Errorf(s)
		return "", errors.New(s)
	}

	token, err := GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (serv UserService) Invite(usahaId string, param RegisterParam) error {
	return nil
}
