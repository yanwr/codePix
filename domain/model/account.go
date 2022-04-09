package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Account struct {
	Id        string    `json:"id" valid:"required"`
	OwnerName string    `json:"ownerName" valid:"notnull"`
	Bank      *Bank     `valid:"-"`
	Number    string    `json:"number" valid:"-"`
	PixKeys   []*PixKey `valid:"-"`
	CreatedAt time.Time `json:"createdAt" valid:"required"`
	UpdatedAt time.Time `json:"updatedAt" valid:"required"`
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)
	if err != nil {
		return err
	}
	return nil
}

func NewAccount(bank *Bank, ownerName string, number string) (*Account, error) {
	account := Account{
		Id:        uuid.NewV4().String(),
		OwnerName: ownerName,
		Bank:      bank,
		Number:    number,
		CreatedAt: time.Now(),
	}

	err := account.isValid()
	if err != nil {
		return nil, err
	}

	return &account, nil
}
