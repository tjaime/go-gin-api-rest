package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tjaime/go-gin-api-rest/controller"
)

func SetupRotasTest() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestVerificaSaudacaoStatusOk(t *testing.T) {
	r := SetupRotasTest()
	r.GET("/:nome", controller.Saudacao)

	req, _ := http.NewRequest("GET", "/thiago", nil)
	rsp := httptest.NewRecorder()

	r.ServeHTTP(rsp, req)
	assert.Equal(t, http.StatusOK, rsp.Code, "Status code.")

	mockRsp := `{"API diz":"Olá, thiago. De boas?"}`
	rspBody, _ := ioutil.ReadAll(rsp.Body)
	assert.Equal(t, mockRsp, string(rspBody))

	fmt.Println(string(rspBody))
}
