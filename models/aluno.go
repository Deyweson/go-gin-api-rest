package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	Nome string `json:"nome"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}

var Alunos = []Aluno{
	{Nome: "aluno1", CPF: "123", RG: "123"},
	{Nome: "aluno1", CPF: "123", RG: "123"},
	{Nome: "aluno1", CPF: "123", RG: "123"},
	{Nome: "aluno1", CPF: "123", RG: "123"},
}
