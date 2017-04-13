package model

type User struct {
	ID       string
	Nama     string
	Email    string
	Password string
	Telepon  string
	Alamat   string
}

type Usaha struct {
	ID          string
	Name        string
	Description string
}

type UserUsaha struct {
	ID      string
	UserId  string
	UsahaId string
	Usaha   Usaha
}

type HakAkses struct {
	ID      string
	UsahaId string
	Name    string
}

type UserHakAkses struct {
	ID         string
	UserId     string
	HakAksesId string
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

type Akun struct {
	ID          string
	UsahaId     string
	Name        string
	Code        string
	Level       int
	Side        string
	ChildType   string
	CurrentCode int
	ChildCount  int
	ParentId    string
}
