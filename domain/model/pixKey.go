package model

import (
	"errors"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type PixKeyRepository interface {
	register(pixKey *PixKey) (*PixKey, error)
	findByKind(key string, kind string) (*PixKey, error)
	addBank(bank *Bank) error
	addAccount(account *Account) error
	findAccount(accountId string) (*Account, error)
}

const (
	EMAIL    string = "email"
	CPF      string = "CPF"
	ACTIVE   string = "active"
	INACTIVE string = "inactive"
)

type PixKey struct {
	Id        string    `json:"id" valid:"required"`
	Kind      string    `json:"type" valid:"notnull"`
	Key       string    `json:"key" valid:"notnull"`
	AccountId string    `json:"accountId" valid:"notnull"`
	Account   *Account  `valid:"-"`
	Status    string    `json:"status" valid:"notnull"`
	CreatedAt time.Time `json:"createdAt" valid:"required"`
	UpdatedAt time.Time `json:"updatedAt" valid:"required"`
}

func (pixKey *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(pixKey)
	if err != nil {
		return err
	}
	err = pixKey.isKindValid()
	if err != nil {
		return err
	}
	err = pixKey.isStatusValid()
	if err != nil {
		return err
	}
	return nil
}

func (pixKey *PixKey) isKindValid() error {
	if pixKey.Kind != EMAIL && pixKey.Kind != CPF {
		return errors.New("invalid type of PixKey")
	}
	return nil
}

func (pixKey *PixKey) isStatusValid() error {
	if pixKey.Status != ACTIVE && pixKey.Status != INACTIVE {
		return errors.New("invalid status of PixKey")
	}
	return nil
}

func NewPixKey(kind string, account *Account, key string) (*PixKey, error) {
	pixKey := PixKey{
		Id:        uuid.NewV4().String(),
		Kind:      kind,
		Key:       key,
		AccountId: "",
		Account:   account,
		Status:    ACTIVE,
		CreatedAt: time.Now(),
	}

	err := pixKey.isValid()
	if err != nil {
		return nil, err
	}
	return &pixKey, nil
}
