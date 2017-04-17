package model

import (
	"time"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

type Usaha struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UserUsaha struct {
	ID      string `json:"id"`
	UserId  string `json:"user_id"`
	UsahaId string `json:"usaha_id"`
	Usaha   *Usaha `json:"-"`
}

type HakAkses struct {
	ID      string `json:"id"`
	UsahaId string `json:"usaha_id"`
	Name    string `json:"name"`
}

type UserHakAkses struct {
	ID         string    `json:"id"`
	UserId     string    `json:"user_id"`
	HakAksesId string    `json:"hak_akses_id"`
	HakAkses   *HakAkses `json:"-"`
}

// =======================

const (
	ACTIVA  string = "ACTIVA"
	PASSIVA string = "PASSIVA"
)

const (
	SUB_AKUN        string = "SUB_AKUN"
	INVENTORY       string = "INVENTORY"
	WORK_IN_PROCESS string = "WORK_IN_PROCESS"
)

const (
	DEBET  string = "DEBET"
	CREDIT string = "CREDIT"
)

type Akun struct {
	ID          string `json:"id"`
	UsahaId     string `json:"-"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Level       int    `json:"level"`
	Side        string `json:"side"`
	ChildType   string `json:"child_type"`
	CurrentCode int    `json:"-"`
	ChildCount  int    `json:"-"`
	ParentId    string `json:"parent_id"`
	Parent      *Akun  `json:"-"`
	Deleted     bool   `json:"-"`
}

type AkunBalance struct {
	ID            string    `json:"id"`
	UsahaId       string    `json:"-"`
	Date          time.Time `json:"date"`
	Amount        float64   `json:"amount"`
	Balance       float64   `json:"balance"`
	AkunId        string    `json:"akun_id"`
	AkunDirection string    `json:"akun_direction"`
}

type SubAkun struct {
	ID       string `json:"id"`
	UsahaId  string `json:"-"`
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
	Parent   *Akun  `json:"-"`
}

type Jurnal struct {
	ID          string    `json:"id"`
	UsahaId     string    `json:"-"`
	Date        time.Time `json:"date"`
	UserId      string    `json:"user_id"`
	Description string    `json:"description"`
}

type JurnalAkunBalance struct {
	ID            string `json:"id"`
	UsahaId       string `json:"-"`
	JurnalId      string `json:"jurnal_id"`
	AkunBalanceId string `json:"akun_balance_id"`
}
