package db

import (
	"context"
	"database/sql"
	"fmt"
)

// to main aci d properties , for transaction
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// exectx executes a function  within the database transaction
func (store *Store) exectx(ctx context.Context, fn func(*Queries) error) error {

	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx Error: %v, rb Error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// Contains inputparameter for transfer transaction
type TransferTxParams struct {
	FromAcccountId int64 `json:"from_acccount_id"`
	ToAccountId    int64 `json:"to_account_id"`
	Amount         int64 `json:"amount"`
}

//TransferTx will perform trasnsaction from 1 account to another
// Creates Transfer Records , add account entries, update account balance within a single transaction

on section 1.8 time:(8:45)