package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tjaime/go-gin-api-rest/db"
	"github.com/tjaime/go-gin-api-rest/model"
)

func ListarAlunos(c *gin.Context) {
	var alunos []model.Aluno
	db.DB.Find(&alunos)
	c.JSON(http.StatusOK, alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz": "Olá, " + nome + ". De boas?",
	})
}

func InserirAluno(c *gin.Context) {
	var aluno model.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := model.ValidaModel(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func FindAlunoById(c *gin.Context) {
	var aluno model.Aluno
	id := c.Params.ByName("id")
	db.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "Aluno não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func DeleleAlunoById(c *gin.Context) {
	var aluno model.Aluno
	id := c.Params.ByName("id")
	db.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{
		"Mensage": "Aluno excluído com sucesso!",
	})
}

func EditarAluno(c *gin.Context) {
	var aluno model.Aluno
	id := c.Params.ByName("id")
	db.DB.First(&aluno, id)
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := model.ValidaModel(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)

}

func FindAlunoByCpf(c *gin.Context) {
	var aluno model.Aluno
	cpf := c.Params.ByName("cpf")
	db.DB.Where(&model.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "Aluno não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, aluno)
}
