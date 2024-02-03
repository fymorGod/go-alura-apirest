package handlers

import (
	"net/http"

	"github.com/fymorGod/go-alura-apirest/schemas"
	"github.com/gin-gonic/gin"
)

func CreateAlunosHandler(ctx *gin.Context) {
	request := CreateAlunoRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	aluno := schemas.Aluno{
		Name:  request.Name,
		Email: request.Email,
		Cpf:   request.Cpf,
		Rg:    request.Rg,
	}
	if err := db.Create(&aluno).Error; err != nil {
		logger.Errorf("error creating request: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating aluno on database")
		return
	}
	sendSuccess(ctx, "create-aluno", aluno)
}
