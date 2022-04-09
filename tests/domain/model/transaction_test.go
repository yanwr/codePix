package model_test

import (
	"codePix/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestShouldCreateANewTransaction(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, err := model.NewBank(code, name)

	accountNumber := "4586"
	ownerName := "Juliano Silveira da Costa"
	account, err := model.NewAccount(bank, ownerName, accountNumber)

	accountNumberDestination := "6625"
	ownerName = "Mariana Nogueira"
	accountDestination, _ := model.NewAccount(bank, accountNumberDestination, ownerName)

	kind := model.EMAIL
	key := "jsilveira@gmail.com"
	pixKey, err := model.NewPixKey(kind, accountDestination, key)

	require.NotEqual(t, account.Id, accountDestination.Id)

	amount := 3.10
	transaction, err := model.NewTransaction(account, amount, pixKey, "Some description")

	require.Nil(t, err)
	require.NotNil(t, uuid.FromStringOrNil(transaction.Id))
	require.Equal(t, transaction.Amount, amount)
	require.Equal(t, transaction.Status, model.TRANSACTION_PENDING)
	require.Equal(t, transaction.Description, "Some description")
	require.Empty(t, transaction.CancelDescription)

	pixKeySameAccount, err := model.NewPixKey(kind, account, key)

	_, err = model.NewTransaction(account, amount, pixKeySameAccount, "Some description")
	require.NotNil(t, err)

	_, err = model.NewTransaction(account, 0, pixKey, "Some description")
	require.NotNil(t, err)

}

func TestShouldChangeStatusTransaction(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, _ := model.NewBank(code, name)

	accountNumber := "4586"
	ownerName := "Juliano Silveira da Costa"
	account, _ := model.NewAccount(bank, ownerName, accountNumber)

	accountNumberDestination := "6625"
	ownerName = "Mariana Nogueira"
	accountDestination, _ := model.NewAccount(bank, accountNumberDestination, ownerName)

	kind := model.EMAIL
	key := "jsilveira@gmail.com"
	pixKey, _ := model.NewPixKey(kind, accountDestination, key)

	amount := 3.10
	transaction, _ := model.NewTransaction(account, amount, pixKey, "Some description")

	err := transaction.Complete()
	require.Nil(t, err)
	require.Equal(t, transaction.Status, model.TRANSACTION_COMPLETED)

	err = transaction.Cancel("Error")
	require.Nil(t, err)
	require.Equal(t, transaction.Status, model.TRANSACTION_CANCELED)
	require.Equal(t, transaction.CancelDescription, "Error")
}
