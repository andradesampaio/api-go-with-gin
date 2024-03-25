package router

import (
	"github.com/gin-gonic/gin"
	"api-go-with-gin/internal/controller"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controller.ExibeTodosAlunos)
	r.POST("/alunos", controller.CriaNovoAluno)
	r.GET("/aluno/:id", controller.BuscaAluno)
	r.DELETE("/aluno/:id", controller.DeletaAluno)
	r.PATCH("/aluno/:id", controller.EditaAluno)
	r.GET("/aluno/cpf/:cpf", controller.BuscaAlunoPorCpf)
	r.Run()
}