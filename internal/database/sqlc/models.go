// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Computation struct {
	ID              int64            `json:"id"`
	UserID          int64            `json:"user_id"`
	ComputationType string           `json:"computation_type"`
	InputData       string           `json:"input_data"`
	ResultData      string           `json:"result_data"`
	CreatedAt       pgtype.Timestamp `json:"created_at"`
	UpdatedAt       pgtype.Timestamp `json:"updated_at"`
}

type Credit struct {
	ID           int64            `json:"id"`
	UserID       int64            `json:"user_id"`
	TotalCredits pgtype.Int4      `json:"total_credits"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
}

type Token struct {
	ID        int64            `json:"id"`
	UserID    int64            `json:"user_id"`
	Token     string           `json:"token"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	ExpiresAt pgtype.Timestamp `json:"expires_at"`
}

type User struct {
	ID                int64              `json:"id"`
	FirstName         string             `json:"first_name"`
	LastName          string             `json:"last_name"`
	Email             string             `json:"email"`
	Phone             string             `json:"phone"`
	Address           pgtype.Text        `json:"address"`
	PasswordHash      string             `json:"password_hash"`
	PasswordUpdatedAt pgtype.Timestamptz `json:"password_updated_at"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
}

type UssdTransaction struct {
	ID            int64            `json:"id"`
	UserID        int64            `json:"user_id"`
	TransactionID string           `json:"transaction_id"`
	Amount        pgtype.Numeric   `json:"amount"`
	Status        string           `json:"status"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
}
