// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/sqlc-dev/pqtype"
)

type Expense struct {
	ID            int64         `json:"id"`
	TeamID        sql.NullInt64 `json:"team_id"`
	Goal          string        `json:"goal"`
	Amount        string        `json:"amount"`
	Currency      string        `json:"currency"`
	SharingMethod string        `json:"sharing_method"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

type ExpenseDetail struct {
	ID           int64          `json:"id"`
	ExpenseID    sql.NullInt64  `json:"expense_id"`
	MemberID     sql.NullInt64  `json:"member_id"`
	ActualAmount sql.NullString `json:"actual_amount"`
	SharedAmount sql.NullString `json:"shared_amount"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	UserID       int64     `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
	Username     string    `json:"username"`
}

type Team struct {
	ID          int64                 `json:"id"`
	Owner       int64                 `json:"owner"`
	TeamName    string                `json:"team_name"`
	Currency    string                `json:"currency"`
	TeamMembers pqtype.NullRawMessage `json:"team_members"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
}

type TeamMember struct {
	ID           int64         `json:"id"`
	TeamID       sql.NullInt64 `json:"team_id"`
	MemberName   string        `json:"member_name"`
	LinkedUserID sql.NullInt64 `json:"linked_user_id"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}

type User struct {
	ID               int64     `json:"id"`
	Username         string    `json:"username"`
	Email            string    `json:"email"`
	Password         string    `json:"password"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	IsEmailVerified  bool      `json:"is_email_verified"`
	PasswordChangeAt time.Time `json:"password_change_at"`
}

type VerifyEmail struct {
	ID         int64     `json:"id"`
	UserID     int64     `json:"user_id"`
	Email      string    `json:"email"`
	SecretCode string    `json:"secret_code"`
	IsUsed     bool      `json:"is_used"`
	CreatedAt  time.Time `json:"created_at"`
	ExpiredAt  time.Time `json:"expired_at"`
}
