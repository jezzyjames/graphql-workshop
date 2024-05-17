package repository

import (
	"database/sql"
	"grapql-api/pkg/data/models"

	"github.com/graphql-go/graphql"
)

type TransactionRepo struct {
	DB *sql.DB
}

func NewTransactionRepo(db *sql.DB) *TransactionRepo {
	return &TransactionRepo{DB: db}
}

func (r *TransactionRepo) InsertTransaction(params graphql.ResolveParams) (interface{}, error) {
	t := params.Args["transaction"].(map[string]interface{})

	transaction := &models.Transaction{
		BankAccountID: t["bankAccountId"].(string),
		Amount:        t["amount"].(int),
		Status:        t["status"].(int),
	}

	result, err := r.DB.Exec("INSERT INTO bank_transaction (bank_account_id, amount, status) VALUES ($1, $2, $3)", transaction.BankAccountID, transaction.Amount, transaction.Status)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}
