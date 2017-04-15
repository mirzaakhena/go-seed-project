package service

import (
	"bitbucket.org/mirzaakhena/miranc-go/model"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	_ "strconv"
)

type UserService struct {
	DB *gorm.DB
}

type RegisterParam struct {
	Name     string `json:"nama" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

type LoginParam struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (serv UserService) Register(param RegisterParam) (*model.User, error) {

	var count int

	tx := serv.DB.Begin()

	tx.Model(&model.User{}).Where("email = ?", param.Email).Count(&count)

	if count > 0 {
		s := "User dengan email " + param.Email + " sudah terdaftar"
		log.Errorf(s)
		return nil, errors.New(s)
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(param.Password), 10)

	user := &model.User{
		ID:       uuid.NewV4().String(),
		Name:     param.Name,
		Email:    param.Email,
		Password: string(hashedPassword),
		Phone:    param.Phone,
		Address:  param.Address,
	}

	tx.Create(user)

	tx.Commit()

	log.Infof("user %s berhasil didaftarkan", param.Name)

	return user, nil
}

func (serv UserService) Login(param LoginParam) (string, error) {

	var user model.User
	serv.DB.Where("email = ?", param.Email).First(&user)

	if user.ID == "" {
		s := "email atau password tidak cocok"
		log.Errorf(s)
		return "", errors.New(s)
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(param.Password))

	if err != nil {
		s := "email atau password tidak cocok."
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
