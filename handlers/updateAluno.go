package handlers

import (
	"net/http"

	"github.com/fymorGod/go-alura-apirest/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateAlunoHandler(ctx *gin.Context) {
	request := CreateAlunoRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	aluno := schemas.Aluno{}
	if request.Name != "" {
		aluno.Name = request.Name
	}
	if request.Email != "" {
		aluno.Email = request.Email
	}
	if request.Cpf != "" {
		aluno.Cpf = request.Cpf
	}
	if request.Rg != "" {
		aluno.Rg = request.Rg
	}

	if err := db.Save(&aluno).Error; err != nil {
		logger.Errorf("error updating aluno: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating aluno")
		return
	}
	sendSuccess(ctx, "update-aluno", aluno)
}
