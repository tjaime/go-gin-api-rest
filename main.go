package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/alunos", ListarAlunos)
	r.Run()
}

func ListarAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"nome": "Thiago",
	})
}
