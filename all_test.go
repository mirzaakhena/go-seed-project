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
	//=================

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

}
