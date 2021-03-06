package main

import (
	"bitbucket.org/mirzaakhena/miranc-go/model"
	"bitbucket.org/mirzaakhena/miranc-go/service"
	"bytes"
	"encoding/json"
	"github.com/op/go-logging"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var log = logging.MustGetLogger("miranc")

func TestRegister(t *testing.T) {

	filename := "barang.db"

	os.Remove("./" + filename)

	r, _ := MainEngine(filename)

	//================= REGISTER

	{
		postBody := service.RegisterParam{
			Name:     "akhena",
			Email:    "aaa@gmail.com",
			Password: "123",
		}

		body, _ := json.Marshal(postBody)

		req, _ := http.NewRequest("POST", "/register", bytes.NewReader(body))

		resp := httptest.NewRecorder()

		r.ServeHTTP(resp, req)

		user := &model.User{}

		json.Unmarshal(resp.Body.Bytes(), user)

		log.Debug(resp.Body.String())

		assert.Equal(t, "akhena", user.Name)
		assert.Equal(t, "aaa@gmail.com", user.Email)
		assert.NotEmpty(t, user.Password)

		req.Body.Close()
	}

	//================= LOGIN

	var token string

	{
		postBody := service.LoginParam{
			Email:    "aaa@gmail.com",
			Password: "123",
		}

		body, _ := json.Marshal(postBody)

		req, _ := http.NewRequest("POST", "/login", bytes.NewReader(body))

		resp := httptest.NewRecorder()

		r.ServeHTTP(resp, req)

		token = resp.Header().Get("token")
		assert.NotEmpty(t, token)

		req.Body.Close()
	}

	//================= CREATE USAHA

	var usahaId string

	{
		postBody := service.CreateUsahaParam{
			Name:        "Laura",
			Description: "laundry kita",
		}

		body, _ := json.Marshal(postBody)

		req, _ := http.NewRequest("POST", "/usaha", bytes.NewReader(body))

		req.Header.Add("Authorization", "Bearer "+token)

		resp := httptest.NewRecorder()

		r.ServeHTTP(resp, req)

		usaha := &model.Usaha{}

		json.Unmarshal(resp.Body.Bytes(), usaha)

		log.Debug(resp.Body.String())

		assert.Equal(t, "Laura", usaha.Name)
		assert.Equal(t, "laundry kita", usaha.Description)

		usahaId = usaha.ID

		req.Body.Close()
	}

	//================= CREATE MULTIPLE AKUN

	var akunId string

	{
		// postBody := service.CreateAkunParam{
		// 	Name:       "Kas",
		// 	Side:       "ACTIVA",
		// 	ParentCode: "",
		// 	ChildType:  "SUB_AKUN",
		// }

		postBody := []service.CreateAkunParam{
			service.CreateAkunParam{
				Name:       "Harta",
				Side:       "ACTIVA",
				ParentCode: "",
				ChildType:  "AKUN",
			},
			service.CreateAkunParam{
				Name:       "Harta Lancar",
				Side:       "ACTIVA",
				ParentCode: "1",
				ChildType:  "AKUN",
			},
			service.CreateAkunParam{
				Name:       "Kas",
				Side:       "ACTIVA",
				ParentCode: "1.1",
				ChildType:  "SUB_AKUN",
			},
			service.CreateAkunParam{
				Name:       "Bank",
				Side:       "ACTIVA",
				ParentCode: "1.1",
				ChildType:  "SUB_AKUN",
			},
			service.CreateAkunParam{
				Name:       "Piutang",
				Side:       "ACTIVA",
				ParentCode: "1.1",
				ChildType:  "SUB_AKUN",
			},
			service.CreateAkunParam{
				Name:       "Utang",
				Side:       "PASSIVA",
				ParentCode: "",
				ChildType:  "SUB_AKUN",
			},
			service.CreateAkunParam{
				Name:       "Modal",
				Side:       "PASSIVA",
				ParentCode: "",
				ChildType:  "AKUN",
			},
			service.CreateAkunParam{
				Name:       "Modal Mirza",
				Side:       "PASSIVA",
				ParentCode: "3",
				ChildType:  "AKUN",
			},
			service.CreateAkunParam{
				Name:       "Pendapatan",
				Side:       "PASSIVA",
				ParentCode: "",
				ChildType:  "AKUN",
			},
			service.CreateAkunParam{
				Name:       "Biaya",
				Side:       "ACTIVA",
				ParentCode: "",
				ChildType:  "SUB_AKUN",
			},
		}

		body, _ := json.Marshal(postBody)

		req, _ := http.NewRequest("POST", "/usaha/"+usahaId+"/akun", bytes.NewReader(body))

		req.Header.Add("Authorization", "Bearer "+token)

		resp := httptest.NewRecorder()

		r.ServeHTTP(resp, req)

		akun := &model.Akun{}

		json.Unmarshal(resp.Body.Bytes(), akun)

		log.Debug(resp.Body.String())

		assert.Equal(t, "Kas", akun.Name)

		akunId = akun.ID

		req.Body.Close()
	}

	//================= CREATE SINGLE AKUN

	// var akunId string

	{
		postBody := service.CreateAkunParam{
			Name:       "Barang Dagangan",
			Side:       "ACTIVA",
			ParentCode: "1.1",
			ChildType:  "SUB_AKUN",
		}

		body, _ := json.Marshal(postBody)

		req, _ := http.NewRequest("POST", "/usaha/"+usahaId+"/akun", bytes.NewReader(body))

		req.Header.Add("Authorization", "Bearer "+token)

		resp := httptest.NewRecorder()

		r.ServeHTTP(resp, req)

		akun := &model.Akun{}

		json.Unmarshal(resp.Body.Bytes(), akun)

		log.Debug(resp.Body.String())

		assert.Equal(t, "Barang Dagangan", akun.Name)

		akunId = akun.ID

		req.Body.Close()
	}

	//================= CREATE SUB_AKUN

	{
		postBody := service.CreateSubAkunParam{
			Name: "Kas Besar",
		}

		body, _ := json.Marshal(postBody)

		req, _ := http.NewRequest("POST", "/usaha/"+usahaId+"/akun/"+akunId, bytes.NewReader(body))

		req.Header.Add("Authorization", "Bearer "+token)

		resp := httptest.NewRecorder()

		r.ServeHTTP(resp, req)

		akun := &model.SubAkun{}

		json.Unmarshal(resp.Body.Bytes(), akun)

		log.Debug(resp.Body.String())

		assert.Equal(t, "Kas Besar", akun.Name)

		req.Body.Close()
	}

	//================= CREATE JURNAL

	{
		postBody := service.CreateJurnalParam{
			"bikin jurnal",
			[]interface{}{
				service.AkunIO{
					service.BaseAkun{"aaaa"},
					"123",
					20000,
				},
				service.AkunIO{
					service.BaseAkun{"xxxx"},
					"456",
					35000,
				},
			},
		}

		body, _ := json.Marshal(postBody)

		req, _ := http.NewRequest("POST", "/usaha/"+usahaId+"/jurnal", bytes.NewReader(body))

		req.Header.Add("Authorization", "Bearer "+token)

		resp := httptest.NewRecorder()

		r.ServeHTTP(resp, req)

		akun := &model.Jurnal{}

		json.Unmarshal(resp.Body.Bytes(), akun)

		log.Debug(resp.Body.String())

		assert.Equal(t, "Kas Besar", akun.Description)

		req.Body.Close()
	}
}
