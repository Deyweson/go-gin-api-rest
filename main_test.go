package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/deyweson/go-gin-api-rest/controllers"
	"github.com/deyweson/go-gin-api-rest/db"
	"github.com/deyweson/go-gin-api-rest/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriarAlunoMock() {
	aluno := models.Aluno{
		Nome: "Aluno teste",
		CPF:  "12345678900",
		RG:   "123456789",
	}
	db.DB.Create(&aluno)
	ID = int(aluno.ID)
}
func DeletaAlunoMock() {
	var aluno models.Aluno
	db.DB.Delete(&aluno, ID)

}

func TestSaudacao(t *testing.T) {
	r := SetupRotasDeTeste()

	r.GET("/:nome", controllers.Saudacao)
	// Cria a request
	req, _ := http.NewRequest("GET", "/deyve", nil)
	// Cria o response
	res := httptest.NewRecorder()

	// Faz a request e armazena na response a resposta
	r.ServeHTTP(res, req)
	// if res.Code != http.StatusOK {
	// 	t.Fatalf("Status Error: valor recebido foi %d e o esperado era %d", res.Code, http.StatusOK)
	// }

	assert.Equal(t, http.StatusOK, res.Code)

	responseMock := `{"API diz":"Olá deyve, tudo bem?"}`
	responseBody, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, responseMock, string(responseBody))
}

func TestListarAlunos(t *testing.T) {
	db.ConectarDB()
	CriarAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRotasDeTeste()
	r.GET("/alunos", controllers.ExibeAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

}

func TestBuscarAlunoPorCPF(t *testing.T) {
	db.ConectarDB()
	CriarAlunoMock()
	defer DeletaAlunoMock()

	r := SetupRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscarAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestBuscarAlunoPorID(t *testing.T) {
	db.ConectarDB()
	CriarAlunoMock()
	defer DeletaAlunoMock()

	r := SetupRotasDeTeste()
	r.GET("/alunos/:id", controllers.BuscarAluno)
	reqPath := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", reqPath, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	var aluno models.Aluno
	json.Unmarshal(res.Body.Bytes(), &aluno)

	assert.Equal(t, "Aluno teste", aluno.Nome)
	assert.Equal(t, "12345678900", aluno.CPF)
	assert.Equal(t, "123456789", aluno.RG)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestDeletarAluno(t *testing.T) {
	db.ConectarDB()
	CriarAlunoMock()

	r := SetupRotasDeTeste()
	r.DELETE("alunos/:id", controllers.DeletarAluno)
	reqPath := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", reqPath, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestAtualizarAluno(t *testing.T) {
	db.ConectarDB()
	CriarAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRotasDeTeste()
	r.PATCH("/alunos/:id", controllers.EditarAluno)
	reqPath := "/alunos/" + strconv.Itoa(ID)

	aluno := models.Aluno{
		Nome: "Aluno teste",
		CPF:  "12345678999",
		RG:   "999999999",
	}
	alunoJson, _ := json.Marshal(aluno)

	req, _ := http.NewRequest("PATCH", reqPath, bytes.NewBuffer(alunoJson))
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)
	var alunoAtualizado models.Aluno
	json.Unmarshal(res.Body.Bytes(), &alunoAtualizado)

	assert.Equal(t, aluno.Nome, alunoAtualizado.Nome)
	assert.Equal(t, aluno.CPF, alunoAtualizado.CPF)
	assert.Equal(t, aluno.RG, alunoAtualizado.RG)
	assert.Equal(t, http.StatusOK, res.Code)

}
