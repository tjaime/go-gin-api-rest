package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tjaime/go-gin-api-rest/controller"
)

func HandleRequest() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.NoRoute(controller.Page404)
	r.GET("/alunos", controller.ListarAlunos)
	r.POST("/alunos", controller.InserirAluno)
	r.GET("/:nome", controller.Saudacao)
	r.GET("/alunos/:id", controller.FindAlunoById)
	r.DELETE("/alunos/:id", controller.DeleleAlunoById)
	r.PATCH("/alunos/:id", controller.EditarAluno)
	r.GET("/alunos/cpf/:cpf", controller.FindAlunoByCpf)
	r.Run()
}
