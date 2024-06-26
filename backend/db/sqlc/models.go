// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Expense struct {
	ID            int64            `json:"id"`
	TeamID        pgtype.Int8      `json:"team_id"`
	Goal          string           `json:"goal"`
	Amount        pgtype.Numeric   `json:"amount"`
	Currency      string           `json:"currency"`
	SharingMethod string           `json:"sharing_method"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
	UpdatedAt     pgtype.Timestamp `json:"updated_at"`
}

type ExpenseDetail struct {
	ID           int64            `json:"id"`
	ExpenseID    pgtype.Int8      `json:"expense_id"`
	MemberID     pgtype.Int8      `json:"member_id"`
	ActualAmount pgtype.Numeric   `json:"actual_amount"`
	SharedAmount pgtype.Numeric   `json:"shared_amount"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
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
	ID          int64            `json:"id"`
	Owner       int64            `json:"owner"`
	TeamName    string           `json:"team_name"`
	Currency    string           `json:"currency"`
	TeamMembers []byte           `json:"team_members"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
}

type TeamMember struct {
	ID           int64            `json:"id"`
	TeamID       pgtype.Int8      `json:"team_id"`
	MemberName   string           `json:"member_name"`
	LinkedUserID pgtype.Int8      `json:"linked_user_id"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
}

type User struct {
	ID               int64            `json:"id"`
	Username         string           `json:"username"`
	Email            string           `json:"email"`
	Password         string           `json:"password"`
	CreatedAt        pgtype.Timestamp `json:"created_at"`
	UpdatedAt        pgtype.Timestamp `json:"updated_at"`
	IsEmailVerified  bool             `json:"is_email_verified"`
	PasswordChangeAt time.Time        `json:"password_change_at"`
}

type VerifyEmail struct {
	ID         int64            `json:"id"`
	UserID     int64            `json:"user_id"`
	Email      string           `json:"email"`
	SecretCode string           `json:"secret_code"`
	IsUsed     bool             `json:"is_used"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
	ExpiredAt  pgtype.Timestamp `json:"expired_at"`
}
