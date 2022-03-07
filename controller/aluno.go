package controller

import (
	"log"
	"net/http"
	"strconv"

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
	db.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func FindAlunoById(c *gin.Context) {
	var aluno model.Aluno
	id := c.Params.ByName("id")
	idConvertido, err := strconv.Atoi(id)
	if err != nil {
		log.Panic("Erro ao converter o id do Aluno")
	}
	db.DB.First(&aluno, idConvertido)

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
