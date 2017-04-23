package service

import (
	"bitbucket.org/mirzaakhena/miranc-go/model"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	// "strconv"
)

type SubAkunService struct {
	DB *gorm.DB
}

type CreateSubAkunWithParentCodeParam struct {
	Name       string `json:"name" binding:"required"`
	ParentCode string `json:"parent_code" binding:"required"`
}

type CreateSubAkunParam struct {
	Name string `json:"name" binding:"required"`
}

func (serv SubAkunService) CreateSubAkunWithParentCode(usahaId string, param CreateSubAkunWithParentCodeParam) error {

	// TODO apakah boleh membuat subakun
	var akun model.Akun
	serv.DB.Where("usaha_id = ? AND code = ? AND deleted = 0", usahaId, param.ParentCode).First(&akun)

	if akun.ID == "" {
		m := "akun with code " + param.ParentCode + " not Found"
		log.Error(m)
		return errors.New(m)
	}

	serv.DB.Create(&model.SubAkun{
		ID:       uuid.NewV4().String(),
		UsahaId:  usahaId,
		Name:     param.Name,
		ParentId: akun.ID,
	})

	return nil

}

func (serv SubAkunService) CreateSubAkun(usahaId string, akunId string, param CreateSubAkunParam) (*model.SubAkun, error) {

	// TODO apakah boleh membuat subakun
	var akun model.Akun
	serv.DB.Where("usaha_id = ? AND id = ? AND deleted = 0", usahaId, akunId).First(&akun)

	if akun.ID == "" {
		m := "akun with id " + akunId + " not Found"
		log.Error(m)
		return nil, errors.New(m)
	}

	subAkun := &model.SubAkun{
		ID:       uuid.NewV4().String(),
		UsahaId:  usahaId,
		Name:     param.Name,
		ParentId: akun.ID,
	}

	serv.DB.Create(subAkun)

	return subAkun, nil

}

func (serv SubAkunService) GetAllSubAkun(usahaId string, akunId string) []model.SubAkun {
	var listOfSubAkun []model.SubAkun
	serv.DB.Where("usaha_id = ? AND parent_id = ?", usahaId, akunId).Find(&listOfSubAkun)
	return listOfSubAkun
}

func (serv SubAkunService) GetById(usahaId string, subAkunId string) *model.SubAkun {
	var subAkun model.SubAkun
	serv.DB.Preload("Parent").Where("usaha_id = ? AND id = ?", usahaId, subAkunId).First(&subAkun)
	return &subAkun
}

func (serv SubAkunService) GetSubAkunByName(usahaId string, parentCode string, name string) *model.SubAkun {
	var listOfSubAkun []model.SubAkun
	serv.DB.Preload("Parent").Where("usaha_id = ? AND name = ?", usahaId, name).Find(&listOfSubAkun)

	for _, obj := range listOfSubAkun {
		if obj.Parent.Code == parentCode {
			return &obj
		}
	}

	return nil
}
