package handlers

import (
	"fmt"
	"net/http"

	"github.com/fymorGod/go-alura-apirest/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteAlunoHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	aluno := schemas.Aluno{}
	if err := db.First(&aluno, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("aluno with id: %s not found", id))
		return
	}

	if err := db.Delete(&aluno).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting aluno with id: %s", id))
		return
	}
	sendSuccess(ctx, "delete-aluno", aluno)
}
