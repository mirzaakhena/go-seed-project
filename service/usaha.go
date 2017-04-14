package service

import (
	"bitbucket.org/mirzaakhena/miranc-go/model"
	// "errors"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	// "strconv"
)

type UsahaService struct {
	DB *gorm.DB
}

type CreateUsahaParam struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func (serv UsahaService) CreateUsaha(userId string, param CreateUsahaParam) error {

	usahaId := uuid.NewV4().String()
	serv.DB.Create(&model.Usaha{
		ID:          usahaId,
		Name:        param.Name,
		Description: param.Description,
	})

	serv.DB.Create(&model.UserUsaha{
		ID:      uuid.NewV4().String(),
		UsahaId: usahaId,
		UserId:  userId,
	})

	return nil
}

func (serv UsahaService) GetAllUsahaByUser(userId string) []model.Usaha {
	var listOfUserUsaha []model.UserUsaha
	serv.DB.Preload("Usaha").Where("user_id = ?", userId).Find(&listOfUserUsaha)

	listOfUsaha := make([]model.Usaha, 0)
	for _, obj := range listOfUserUsaha {
		listOfUsaha = append(listOfUsaha, *obj.Usaha)
	}

	return listOfUsaha
}
