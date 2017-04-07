package service

import (
	"bitbucket.org/mirzaakhena/go-seed-project/dto"
	"bitbucket.org/mirzaakhena/go-seed-project/model"
	"bitbucket.org/mirzaakhena/go-seed-project/repo"
)

type Barang struct {
	Repo repo.Barang
}

func (service Barang) AddBarang(params dto.AddBarangParam) {
	model := model.Barang{Nama: params.Nama}
	service.Repo.Save(model)
}

func (service Barang) GetAllBarang() []model.Barang {
	return service.Repo.GetAll()
}