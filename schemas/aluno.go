package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	Cpf   string `json:"cpf"`
	Rg    string `json:"rg"`
}

type AlunoResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt, omitempty"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Cpf       string    `json:"cpf"`
	Rg        string    `json:"rg"`
}
