package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tjaime/go-gin-api-rest/controller"
	"github.com/tjaime/go-gin-api-rest/db"
	"github.com/tjaime/go-gin-api-rest/model"
)

func SetupRotasTest() *gin.Engine {
	rotas := gin.Default()
	gin.SetMode(gin.ReleaseMode)
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
}

func TestListarAlunosHandler(t *testing.T) {
	db.ConnectDb()

	// Mock
	InsertlunoMock()
	defer DeleteAlunoMock()

	r := SetupRotasTest()
	r.GET("/alunos", controller.ListarAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	rsp := httptest.NewRecorder()
	r.ServeHTTP(rsp, req)

	assert.Equal(t, http.StatusOK, rsp.Code)
}

func TestFindAlunoByCpf(t *testing.T) {
	db.ConnectDb()

	// Mock
	InsertlunoMock()
	defer DeleteAlunoMock()

	r := SetupRotasTest()
	r.GET("/alunos/cpf/:cpf", controller.FindAlunoByCpf)
	req, _ := http.NewRequest("GET", "/alunos/cpf/31317916832", nil)
	rsp := httptest.NewRecorder()

	r.ServeHTTP(rsp, req)

	assert.Equal(t, http.StatusOK, rsp.Code)
}

var ID int

func InsertlunoMock() {
	aluno := model.Aluno{Nome: "Thiago", CPF: "31317916832", RG: "333831214"}
	db.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeleteAlunoMock() {
	var aluno model.Aluno
	db.DB.Delete(&aluno, ID)
}
