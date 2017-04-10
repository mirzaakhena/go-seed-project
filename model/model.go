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
