package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
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

func TestFindAlunoById(t *testing.T) {
	db.ConnectDb()

	// Mock data
	InsertlunoMock()
	defer DeleteAlunoMock()

	r := SetupRotasTest()
	r.GET("/alunos/:id", controller.FindAlunoById)

	pathAlunoById := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathAlunoById, nil)
	rsp := httptest.NewRecorder()

	r.ServeHTTP(rsp, req)

	var alunoMock model.Aluno
	json.Unmarshal(rsp.Body.Bytes(), &alunoMock)

	assert.Equal(t, "Thiago", alunoMock.Nome)
	assert.Equal(t, "31317916832", alunoMock.CPF)
	assert.Equal(t, "333831214", alunoMock.RG)

}

func TestEditAluno(t *testing.T) {

	db.ConnectDb()

	InsertlunoMock()
	defer DeleteAlunoMock()

	r := SetupRotasTest()
	r.PATCH("/alunos/:id", controller.EditarAluno)

	aluno := model.Aluno{Nome: "Thiago Jaime", CPF: "71717916832", RG: "443831214"}
	jsonAluno, _ := json.Marshal(aluno)

	pathEdit := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathEdit, bytes.NewBuffer(jsonAluno))
	rsp := httptest.NewRecorder()
	r.ServeHTTP(rsp, req)

	var alunoTarget model.Aluno
	json.Unmarshal(rsp.Body.Bytes(), &alunoTarget)

	assert.Equal(t, "71717916832", alunoTarget.CPF)
	assert.Equal(t, "443831214", alunoTarget.RG)
	assert.Equal(t, "Thiago Jaime", alunoTarget.Nome)

}

func TestDeleteAluno(t *testing.T) {

	db.ConnectDb()
	InsertlunoMock()

	r := SetupRotasTest()
	r.DELETE("/alunos/:id", controller.DeleleAlunoById)

	pathAlunoById := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathAlunoById, nil)
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
