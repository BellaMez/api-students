package api

import "fmt"

type StudentRequest struct {
	Name   string `json:"name"`
	CPF    int    `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active *bool  `json:"registration"` // using bool as a pointer to force true/false input
}

func errParamRequired(param, typ string) error {
	return fmt.Errorf("param '%s'of type '%s' is required", param, typ)
}

func (s *StudentRequest) Validate() error {
	if s.Name == "" {
		return errParamRequired("name", "string")
	}
	if s.CPF == 0 {
		return errParamRequired("cpf", "string")
	}
	if s.Email == "" {
		return errParamRequired("email", "string")
	}
	if s.Age == 0 {
		return errParamRequired("age", "string")
	}
	if s.Active == nil {
		return errParamRequired("registration", "string")
	}
	return nil

}
