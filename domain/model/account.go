package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Account struct {
	Id        string    `json:"id" gorm:"type:uuid;primary_key" valid:"required"`
	OwnerName string    `json:"ownerName" gorm:"column:ownerName;type:varchar(255);not null" valid:"notnull"`
	Bank      *Bank     `valid:"-"`
	BankId    string    `json:"bankId" gorm:"column:bankId;type:uuid;not null" valid:"notnull"`
	Number    string    `json:"number" gorm:"column:number;type:varchar(20);not null" valid:"notnull"`
	PixKeys   []*PixKey `gorm:"ForeignKey:AccountId" valid:"-"`
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
		BankId:    bank.Id,
		Number:    number,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := account.isValid()
	if err != nil {
		return nil, err
	}

	return &account, nil
}
