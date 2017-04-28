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
	DB             *gorm.DB
	AkunService    *AkunService
	SubAkunService *SubAkunService
}

type BaseAkun struct {
	Type string
}

type AkunIO struct {
	BaseAkun
	ID     string
	Amount float64
}

type InventoryInput struct {
	BaseAkun
	ID       string
	Price    float64
	Quantity float64
}

type InventoryOutput struct {
	BaseAkun
	ID    string
	Price float64
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

	for _, baseAkun := range param.Akuns {
		serv.getAmountOfAkun(jurnal, baseAkun)

	}

	log.Infof("jurnal %s berhasil dibuat", param.Description)

	return jurnal, nil

}

func (serv JurnalService) GetAllJurnal(usahaId string) []model.Jurnal {
	var listOfJurnal []model.Jurnal
	serv.DB.Where("usaha_id = ? AND deleted = 0", usahaId).Find(&listOfJurnal)
	return listOfJurnal
}

func (serv JurnalService) getAmountOfAkun(jurnal *model.Jurnal, baseAkun interface{}) (code string, direction string, amount float64) {
	switch baseAkun.(type) {
	case AkunIO:

		x, _ := baseAkun.(AkunIO)

		if x.Type == "akun-input" {
			akun := serv.AkunService.GetById(jurnal.UsahaId, x.ID)
			{
				code = akun.Code
				direction = serv.getDirection(akun.Side, x.Amount)
				amount = x.Amount
			}
			return

		} else if x.Type == "akun-output" {
			akun := serv.AkunService.GetById(jurnal.UsahaId, x.ID)
			{
				code = akun.Code
				direction = serv.getDirection(akun.Side, -x.Amount)
				amount = -x.Amount
			}
			return

		} else if x.Type == "subakun-input" {
			subakun := serv.SubAkunService.GetById(jurnal.UsahaId, x.ID)

			// get the last balance
			var subAkunBal model.SubAkunBalance
			serv.DB.Preload("Jurnal").Where("usaha_id = ? AND sub_akun_id = ? AND jurnal.date <= ?", jurnal.UsahaId, x.ID, jurnal.Date).Order("jurnal.date DESC").First(&subAkunBal)
			lastBal := subAkunBal.Balance

			akunBalance := &model.AkunBalance{
				ID:            uuid.NewV4().String(),
				UsahaId:       jurnal.UsahaId,
				AkunDirection: serv.getDirection(subakun.Parent.Side, x.Amount),
				AkunId:        subakun.ParentId,
				Date:          jurnal.Date,
				Amount:        x.Amount,
				Balance:       lastBal + x.Amount,
			}

			serv.DB.Create(akunBalance)

			{
				code = subakun.Parent.Code
				direction = serv.getDirection(subakun.Parent.Side, x.Amount)
				amount = x.Amount
			}
			return

		} else if x.Type == "subakun-output" {
			subakun := serv.SubAkunService.GetById(jurnal.UsahaId, x.ID)

			// get the last balance
			var subAkunBal model.SubAkunBalance
			serv.DB.Preload("Jurnal").Where("usaha_id = ? AND sub_akun_id = ? AND jurnal.date <= ?", jurnal.UsahaId, x.ID, jurnal.Date).Order("jurnal.date DESC").First(&subAkunBal)
			lastBal := subAkunBal.Balance

			akunBalance := &model.AkunBalance{
				ID:            uuid.NewV4().String(),
				UsahaId:       jurnal.UsahaId,
				AkunDirection: serv.getDirection(subakun.Parent.Side, x.Amount),
				AkunId:        subakun.ParentId,
				Date:          jurnal.Date,
				Amount:        x.Amount,
				Balance:       lastBal + -x.Amount,
			}

			serv.DB.Create(akunBalance)

			{
				code = subakun.Parent.Code
				direction = serv.getDirection(subakun.Parent.Side, -x.Amount)
				amount = -x.Amount
			}
			return
		}

	case InventoryInput:
		x, _ := baseAkun.(InventoryInput)
		log.Debug(x)
	case InventoryOutput:
		x, _ := baseAkun.(InventoryOutput)
		log.Debug(x)
	}

	return "", "", 0
}

func (serv JurnalService) getDirection(side string, amount float64) string {
	if (side == "ACTIVA" && amount > 0) || (side == "PASSIVA" && amount < 0) {
		return "DEBET"
	}
	if (side == "PASSIVA" && amount > 0) || (side == "ACTIVA" && amount < 0) {
		return "CREDIT"
	}
	return ""
}
