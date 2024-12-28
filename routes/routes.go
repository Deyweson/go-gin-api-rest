package routes

import (
	"github.com/deyweson/go-gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()

	r.GET("/alunos", controllers.ExibeAlunos)
	r.GET("/:nome", controllers.Saudacao)

	r.Run(":8000")
}
