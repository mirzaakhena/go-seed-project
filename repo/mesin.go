package repo

import (
	"bitbucket.org/mirzaakhena/go-seed-project/model"
	"github.com/jinzhu/gorm"
)

type Barang struct {
	DB *gorm.DB
}

func (repo Barang) Save(model model.Barang) {
	repo.DB.Create(&model)
}

func (repo Barang) GetAll() []model.Barang {
	listOfModel := []model.Barang{}
	repo.DB.Find(&listOfModel)
	return listOfModel
}
