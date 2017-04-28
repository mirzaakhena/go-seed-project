package service

import (
	"bitbucket.org/mirzaakhena/miranc-go/model"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"strconv"
)

type AkunService struct {
	DB *gorm.DB
}

type CreateAkunParam struct {
	Name       string `json:"name" binding:"required"`
	ParentCode string `json:"parent_code"`
	Side       string `json:"side"`
	ChildType  string `json:"child_type" binding:"required"`
}

func (serv AkunService) Create(usahaId string, params []CreateAkunParam) ([]model.Akun, error) {

	// TODO apakah user terdaftar pada usaha ini?
	// TODO apakah user diijinkan utk membuat akun?

	if len(params) == 0 {
		return nil, errors.New("empty arrays")
	}

	akuns := []model.Akun{}

	tx := serv.DB.Begin()

	for _, param := range params {

		if param.ParentCode == "" {
			var count int

			if param.Side != "ACTIVA" && param.Side != "PASSIVA" {
				tx.Rollback()
				s := "sisi harus PASSIVA atau ACTIVA"
				log.Errorf(s)
				return nil, errors.New(s)
			}

			// TODO apakah dalam usaha dan level ini ada yang namanya sama?

			tx.Model(&model.Akun{}).Where("usaha_id=? AND parent_id = ?", usahaId, "").Count(&count)

			akun := &model.Akun{
				ID:          uuid.NewV4().String(),
				Name:        param.Name,
				UsahaId:     usahaId,
				ChildType:   param.ChildType,
				Code:        strconv.Itoa(count + 1),
				Side:        param.Side,
				Level:       1,
				ParentId:    "",
				CurrentCode: 0,
				ChildCount:  0,
				Deleted:     false,
			}

			akuns = append(akuns, *akun)

			tx.Create(akun)

			log.Infof("akun parent %s berhasil dibuat", param.Name)

		} else {

			var parentAkun model.Akun

			tx.Where("code = ? AND usaha_id = ?", param.ParentCode, usahaId).First(&parentAkun)

			// TODO apakah parent benar ada?
			// TODO apakah dalam usaha dan parent_level + 1 ini ada nama yang sama?

			nextChildIndex := parentAkun.CurrentCode + 1

			akun := &model.Akun{
				ID:          uuid.NewV4().String(),
				Name:        param.Name,
				UsahaId:     usahaId,
				ChildType:   param.ChildType,
				Code:        param.ParentCode + "." + strconv.Itoa(nextChildIndex),
				Side:        parentAkun.Side,
				Level:       parentAkun.Level + 1,
				ParentId:    parentAkun.ID,
				CurrentCode: 0,
				ChildCount:  0,
				Deleted:     false,
			}

			akuns = append(akuns, *akun)

			err := tx.Create(akun).Error

			if err != nil {
				tx.Rollback()
				log.Errorf("error ketika buat akun anak. Rollback! %s", err.Error())
				return nil, err
			}

			parentAkun.CurrentCode = parentAkun.CurrentCode + 1
			parentAkun.ChildCount = parentAkun.ChildCount + 1
			err = tx.Save(&parentAkun).Error
			if err != nil {
				tx.Rollback()
				log.Errorf("error ketika update akun parent setelah buat akun anak. %s", err.Error())
				return nil, err
			}

			log.Infof("akun anak %s berhasil dibuat", param.Name)

		}
	}

	tx.Commit()

	return akuns, nil

}

func (serv AkunService) GetAll(usahaId string) []model.Akun {
	var listOfAkun []model.Akun
	serv.DB.Where("usaha_id = ? AND deleted = 0", usahaId).Find(&listOfAkun)
	return listOfAkun
}

func (serv AkunService) GetById(usahaId string, akunId string) model.Akun {
	var akun model.Akun
	serv.DB.Where("usaha_id = ? AND id = ? AND deleted = 0", usahaId, akunId).First(&akun)
	return akun
}
