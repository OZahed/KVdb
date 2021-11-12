package dto

import (
	"github.com/go-playground/validator/v10"
)

type SetDto struct {
	Key   string      `json:"key" validate:"required"`
	Value interface{} `json:"value" validate:"required"`
	Exp   int         `json:"expire" validate:"required,gt=0"`
}

func (s *SetDto) Validate() error {
	return validator.New().Struct(s)
}
