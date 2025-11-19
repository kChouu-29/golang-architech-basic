package model

import "time"

type Users struct {
	ID           int
	Username     string
	Email        string
	PasswordHard string
	Age          int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
