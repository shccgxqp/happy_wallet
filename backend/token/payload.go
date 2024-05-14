package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token id invalid")
	)

type Payload struct {
	ID uuid.UUID `json:"id"`
	Username string `json:"username"`
	IssuedAT time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expires_at"`
}

func NewPayload(username string, duration time.Duration)(*Payload, error){
	tokenID, err := uuid.NewRandom()
	if err!= nil {
		return nil, err
	}

	payload := &Payload{
	ID: tokenID,
	Username: username,
	IssuedAT: time.Now(),
	ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error{
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}