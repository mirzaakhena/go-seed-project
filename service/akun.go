package service

import (
	"bitbucket.org/mirzaakhena/miranc-go/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"strconv"
)

type AkunService struct {
	DB *gorm.DB
}

type CreateAkunParam struct {
	Name       string
	ParentCode string
	Side       string
	ChildType  string
}

func (serv AkunService) CreateNewAkun(param CreateAkunParam) {

	if param.ParentCode == "" {
		var count int
		serv.DB.Model(&model.Akun{}).Where("parent_id = ?", "").Count(&count)

		serv.DB.Create(&model.Akun{
			ID:          uuid.NewV4().String(),
			Name:        param.Name,
			UsahaId:     "something",
			ChildType:   param.ChildType,
			Side:        param.Side,
			Code:        strconv.Itoa(count + 1),
			Level:       1,
			ParentId:    "",
			CurrentCode: 0,
			ChildCount:  0,
		})

		fmt.Println("input akun parent")
	} else {
		var parentAkun model.Akun
		serv.DB.Where("code = ?", param.ParentCode).First(&parentAkun)

		nextChildIndex := parentAkun.CurrentCode + 1

		serv.DB.Create(&model.Akun{
			ID:          uuid.NewV4().String(),
			Name:        param.Name,
			UsahaId:     "something",
			ChildType:   param.ChildType,
			Side:        parentAkun.Side,
			Code:        param.ParentCode + "." + strconv.Itoa(nextChildIndex),
			Level:       parentAkun.Level + 1,
			ParentId:    parentAkun.ParentId,
			CurrentCode: 0,
			ChildCount:  0,
		})

		parentAkun.CurrentCode = parentAkun.CurrentCode + 1
		parentAkun.ChildCount = parentAkun.ChildCount + 1
		serv.DB.Save(&parentAkun)

		fmt.Println("input anak parentAkun " + parentAkun.Name)
	}

	// serv.DB.Create(&model.Akun{Nama: param.Name})
}
