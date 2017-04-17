package service

import (
	"bitbucket.org/mirzaakhena/miranc-go/model"
	// "errors"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	// "strconv"
	"time"
)

type JurnalService struct {
	DB *gorm.DB
}

type BaseAkun struct {
	Type string
}

type AkunInputOutput struct {
	BaseAkun
	ID     string
	Amount float64
}

type AkunOutput struct {
	BaseAkun
	ID     string
	Amount float64
}

type CreateJurnalParam struct {
	Description string        `json:"description" binding:"required"`
	Akuns       []interface{} `json:"akuns"`
}

func (serv JurnalService) CreateJurnal(usahaId string, userId string, param CreateJurnalParam) (*model.Jurnal, error) {

	jurnal := &model.Jurnal{
		ID:          uuid.NewV4().String(),
		UsahaId:     usahaId,
		UserId:      userId,
		Description: param.Description,
		Date:        time.Now(),
	}

	serv.DB.Create(jurnal)

	log.Infof("jurnal %s berhasil dibuat", param.Description)

	return jurnal, nil

}

func (serv JurnalService) GetAllJurnal(usahaId string) []model.Jurnal {
	var listOfJurnal []model.Jurnal
	serv.DB.Where("usaha_id = ? AND deleted = 0", usahaId).Find(&listOfJurnal)
	return listOfJurnal
}
