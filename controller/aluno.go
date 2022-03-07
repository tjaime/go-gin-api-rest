package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tjaime/go-gin-api-rest/db"
	"github.com/tjaime/go-gin-api-rest/model"
)

func ListarAlunos(c *gin.Context) {
	c.JSON(200, model.Alunos)
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
