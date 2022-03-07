package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tjaime/go-gin-api-rest/controller"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/alunos", controller.ListarAlunos)
	r.POST("/alunos", controller.InserirAluno)
	r.GET("/:nome", controller.Saudacao)
	r.GET("/alunos/:id", controller.FindAlunoById)
	r.Run()
}
