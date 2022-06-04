package models

import (
	"gorm.io/gorm"
	"gopkg.in/validator.v2"
)

type Aluno struct {
	gorm.Model
	Nome string  `json:"nome" validate:"min=2,max=40,regexp=^[a-zA-Z]*$"`
	CPF  string  `json:"cpf" validate:"len=11, regexp=^[0-9]*$"`
	RG   string  `json:"rg" validate:"len=9, regexp=^[0-9]*$"`
	// Email string `validate:"regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$"`
	// Senha string `json:"senha" validate:"min=6,max=12"`
	// Idade uint8  `validate:"nonzero"`
}

func ValidateFields(aluno *Aluno) error {
	err := validator.Validate(aluno)
	if err != nil {
		return err
	}
	return nil
}
