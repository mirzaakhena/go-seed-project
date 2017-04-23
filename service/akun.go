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

func (serv AkunService) Create(usahaId string, param CreateAkunParam) (*model.Akun, error) {

	// TODO apakah user terdaftar pada usaha ini?
	// TODO apakah user diijinkan utk membuat akun?

	if param.ParentCode == "" {
		var count int

		if param.Side != "ACTIVA" && param.Side != "PASSIVA" {
			s := "sisi harus PASSIVA atau ACTIVA"
			log.Errorf(s)
			return nil, errors.New(s)
		}

		// TODO apakah dalam usaha dan level ini ada yang namanya sama?

		serv.DB.Model(&model.Akun{}).Where("parent_id = ?", "").Count(&count)

		log.Debugf("membuat akun parent dengan nama %s", param.Name)

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

		serv.DB.Create(akun)

		log.Infof("akun parent %s berhasil dibuat", param.Name)

		return akun, nil

	} else {
		var parentAkun model.Akun

		tx := serv.DB.Begin()

		tx.Where("code = ?", param.ParentCode).First(&parentAkun)

		// TODO apakah parent benar ada?
		// TODO apakah dalam usaha dan parent_level + 1 ini ada nama yang sama?

		nextChildIndex := parentAkun.CurrentCode + 1

		log.Debugf("membuat akun anak dengan nama %s dibawah parent %s", param.Name, param.ParentCode)

		user := &model.Akun{
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

		err := tx.Create(user).Error

		if err != nil {
			tx.Rollback()
			log.Errorf("error ketika buat akun anak. Rollback! ", err.Error())
			return nil, err
		}

		log.Debugf("update count akun parent jadi ", parentAkun.ChildCount+1)

		parentAkun.CurrentCode = parentAkun.CurrentCode + 1
		parentAkun.ChildCount = parentAkun.ChildCount + 1
		err = tx.Save(&parentAkun).Error
		if err != nil {
			tx.Rollback()
			log.Errorf("error ketika update akun parent setelah buat akun anak. ", err.Error())
			return nil, err
		}

		tx.Commit()

		log.Infof("akun anak %s berhasil dibuat", param.Name)

		return user, nil
	}

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
