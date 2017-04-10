package model

// type User struct {
// 	ID       string
// 	Nama     string
// 	Email    string
// 	Password string
// 	Phone    string
// 	Address  string
// }

// type Usaha struct {
// 	ID   string
// 	Nama string
// }

// type HakAkses struct {
// 	ID      string
// 	UsahaId string
// 	Nama    string
// }

// type UserUsahaHakAkses struct {
// 	ID         string
// 	UserId     string
// 	UsahaId    string
// 	HakAksesId string
// }

// // =======================

type AkunSide string

const (
	ACTIVA  AkunSide = "ACTIVA"
	PASSIVA AkunSide = "PASSIVA"
)

type AkunChildType string

const (
	SUB_AKUN        AkunChildType = "SUB_AKUN"
	INVENTORY       AkunChildType = "INVENTORY"
	WORK_IN_PROCESS AkunChildType = "WORK_IN_PROCESS"
)

type Akun struct {
	ID          string
	UsahaId     string
	Name        string
	Code        string
	Level       int
	Side        AkunSide
	ChildType   AkunChildType
	CurrentCode int
	ChildCount  int
	ParentId    string
}
