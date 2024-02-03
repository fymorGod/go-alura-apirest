package handlers

import (
	"net/http"

	"github.com/fymorGod/go-alura-apirest/schemas"
	"github.com/gin-gonic/gin"
)

func ListAlunosHandler(ctx *gin.Context) {
	alunos := []schemas.Aluno{}
	if err := db.Find(&alunos).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error getting alunos")
		return
	}
	sendSuccess(ctx, "list-alunos", alunos)
}
