package handlers

import (
	"net/http"

	"github.com/fymorGod/go-alura-apirest/schemas"
	"github.com/gin-gonic/gin"
)

func ShowAlunoHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	aluno := schemas.Aluno{}
	if err := db.First(&aluno, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "aluno not found")
		return
	}
	sendSuccess(ctx, "show-aluno", aluno)
}
