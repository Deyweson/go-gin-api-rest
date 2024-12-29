package routes

import (
	"github.com/deyweson/go-gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()

	r.GET("/alunos", controllers.ExibeAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.RegistrarAluno)
	r.GET("alunos/:id", controllers.BuscarAluno)
	r.DELETE("alunos/:id", controllers.DeletarAluno)
	r.PATCH("alunos/:id", controllers.EditarAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscarAlunoPorCPF)

	r.Run(":8000")
}
