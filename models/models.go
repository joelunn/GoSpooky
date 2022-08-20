package models

import "time"

//This is used by gorm to create the tables IF they don't exist (which isn't obvious and is weird.)

type test struct {
	ID        string     `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}
