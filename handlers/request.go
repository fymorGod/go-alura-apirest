package handlers

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// Create Aluno
type CreateAlunoRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Cpf   string `json:"cpf"`
	Rg    string `json:"rg"`
}

func (r *CreateAlunoRequest) Validate() error {
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}
	if r.Cpf == "" {
		return errParamIsRequired("cpf", "string")
	}
	if r.Rg == "" {
		return errParamIsRequired("rg", "string")
	}

	return nil
}

type UpdateAlunoRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Cpf   string `json:"cpf"`
	Rg    string `json:"rg"`
}

func (r *UpdateAlunoRequest) Validate() error {
	if r.Name != "" || r.Email != "" || r.Cpf != "" || r.Rg != "" {
		return nil
	}

	return fmt.Errorf("at least one valid field must be provided")
}
