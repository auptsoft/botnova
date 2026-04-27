package models

import "time"

type User struct {
	Id        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
