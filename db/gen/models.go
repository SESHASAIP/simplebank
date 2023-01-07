// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package gen

import (
	"time"
)

type Accounts struct {
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	Balance   int32     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type Entries struct {
	ID        int64 `json:"id"`
	AccountID int32 `json:"account_id"`
	// can be negative or positive
	Amount    int32     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type Teacher struct {
	ID        int64     `json:"id"`
	Subject   string    `json:"subject"`
	Salary    string    `json:"salary"`
	CreatedAt time.Time `json:"created_at"`
}

type Transfers struct {
	ID            int64 `json:"id"`
	FromAccountID int32 `json:"from_account_id"`
	ToAccountID   int32 `json:"to_account_id"`
	// must be positive
	Amount    int32     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
