package dto

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type TransactionDTO struct {
	Id           string  `json:"id" validate:"required, uuid4"`
	AccountId    string  `json:"accountId" validate:"required, uuid4"`
	Amount       float64 `json:"amount" validate:"required, numeric"`
	PixKeyIdTo   string  `json:"pixKeyIdTo" validate:"required"`
	PixKeyKindTo string  `json:"pixKeyKindTo" validate:"required"`
	Description  string  `json:"description" validate:"-"`
	Status       string  `json:"status" validate:"required"`
	Error        string  `json:"error"`
}

func (t *TransactionDTO) isValid() error {
	v := validator.New()
	err := v.Struct(t)
	if err != nil {
		fmt.Errorf("error during TransactionDTO validation: %s", err.Error())
		return err
	}
	return nil
}

func (t *TransactionDTO) ParseJson(data []byte) error {
	err := json.Unmarshal(data, t)
	if err != nil {
		return err
	}
	err = t.isValid()
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionDTO) ToJson() ([]byte, error) {
	err := t.isValid()
	if err != nil {
		return nil, err
	}

	response, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func NewTransaction() *TransactionDTO {
	return &TransactionDTO{}
}
