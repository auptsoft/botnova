package models

import "time"

// User is the non-sensitive user profile model safe for regular app flows.
type User struct {
	Id        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// UserAuth contains auth-only fields and should not be exposed in API responses.
type UserAuth struct {
	UserID       string
	Email        string
	PasswordHash string
}
