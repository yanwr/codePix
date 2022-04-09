package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Bank struct {
	Id        string     `json:"id" valid:"required"`
	Code      string     `json:"code" valid:"notnull"`
	Name      string     `json:"name" valid:"notnull"`
	Accounts  []*Account `valid:"-"`
	CreatedAt time.Time  `json:"createdAt" valid:"required"`
	UpdatedAt time.Time  `json:"updatedAt" valid:"required"`
}

// this is a method in GO, and it's associate to Struct Bank
func (bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)
	if err != nil {
		return err
	}
	return nil
}

// this is only a function, and it isn't associate to Struct Bank
func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Id:        uuid.NewV4().String(),
		Code:      code,
		Name:      name,
		CreatedAt: time.Now(),
	}

	err := bank.isValid()
	if err != nil {
		return nil, err
	}

	return &bank, nil
}
