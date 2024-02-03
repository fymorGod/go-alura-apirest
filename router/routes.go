package router

import (
	"github.com/fymorGod/go-alura-apirest/handlers"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handlers.InitializeHandler()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/alunos", handlers.CreateAlunosHandler)
		v1.GET("/alunos", handlers.ListAlunosHandler)
		v1.GET("/aluno", handlers.ShowAlunoHandler)
		v1.GET("/search/aluno", handlers.SearchAlunoCpfHandler)
		v1.DELETE("/aluno", handlers.DeleteAlunoHandler)
		v1.PUT("/aluno", handlers.UpdateAlunoHandler)
	}
}
