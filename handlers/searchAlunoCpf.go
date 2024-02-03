package handlers

import (
	"fmt"
	"net/http"

	"github.com/fymorGod/go-alura-apirest/schemas"
	"github.com/gin-gonic/gin"
)

func SearchAlunoCpfHandler(ctx *gin.Context) {
	cpf := ctx.Query("cpf")
	if cpf == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("cpf", "queryParameter").Error())
		return
	}

	aluno := schemas.Aluno{}
	if err := db.Where(&schemas.Aluno{Cpf: cpf}).First(&aluno).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("aluno with cpf: %s not found", cpf))
		return
	}
	sendSuccess(ctx, "search-aluno", aluno)
}
