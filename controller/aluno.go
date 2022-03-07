package controller

import "github.com/gin-gonic/gin"

func ListarAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"nome": "Thiago",
	})
}

func Saudacao(c *gin.Context) {

	nome := c.Params.ByName("nome")

	c.JSON(200, gin.H{
		"API diz": "Olá, " + nome + ". De boas?",
	})
}
