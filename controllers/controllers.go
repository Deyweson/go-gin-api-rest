package controllers

import (
	"net/http"

	"github.com/deyweson/go-gin-api-rest/db"
	"github.com/deyweson/go-gin-api-rest/models"
	"github.com/gin-gonic/gin"
)

func ExibeAlunos(c *gin.Context) {
	var alunos []models.Aluno

	db.DB.Find(&alunos)

	c.JSON(200, alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")

	c.JSON((200), gin.H{
		"API diz": "Olá " + nome + ", tudo bem?",
	})
}

func RegistrarAluno(c *gin.Context) {
	var aluno models.Aluno

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	if err := models.ValidarAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	db.DB.Create(&aluno)

	c.JSON(http.StatusOK, aluno)
}

func BuscarAluno(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno
	db.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func DeletarAluno(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno
	db.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{
		"Message": "Aluno deletado com sucesso",
	})
}

func EditarAluno(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno
	db.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	if err := models.ValidarAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	db.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscarAlunoPorCPF(c *gin.Context) {
	cpf := c.Param("cpf")
	var aluno models.Aluno
	db.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"mensagem": "Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, aluno)

}

func ExibePaginaIndex(c *gin.Context) {

	var alunos []models.Aluno

	db.DB.Find(&alunos)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

func NotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
