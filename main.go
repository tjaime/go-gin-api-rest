package main

import (
	"github.com/tjaime/go-gin-api-rest/model"
	"github.com/tjaime/go-gin-api-rest/routes"
)

func main() {
	model.Alunos = []model.Aluno{
		{Nome: "Thiago Jaime", CPF: "313.179.168-32", RG: "33.383.121-4"},
		{Nome: "Ana Jaime", CPF: "400.179.168-32", RG: "44.383.121-4"},
		{Nome: "Raque Jaime", CPF: "500.179.168-32", RG: "55.383.121-4"},
	}
	routes.HandleRequest()
}
