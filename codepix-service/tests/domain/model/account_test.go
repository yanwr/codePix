package model_test

import (
	"codePix/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestShouldCreateANewAccount(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, err := model.NewBank(code, name)

	accountNumber := "4586"
	ownerName := "Juliano Silveira da Costa"
	account, err := model.NewAccount(bank, ownerName, accountNumber)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(account.Id))
	require.Equal(t, account.Number, accountNumber)
	require.Equal(t, account.BankId, bank.Id)

	_, err = model.NewAccount(bank, "", ownerName)
	require.Nil(t, err)
}
