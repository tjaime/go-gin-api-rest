package controller

import (
	"github.com/gin-gonic/gin"
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
