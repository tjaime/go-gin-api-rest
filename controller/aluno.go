package controller

import "github.com/gin-gonic/gin"

func ListarAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"nome": "Thiago",
	})
}
