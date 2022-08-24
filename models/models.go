package models

import (
	"time"
)

//This is used by gorm to create the tables IF they don't exist (which isn't obvious and is weird.)

//Effective the Organistation or Group
type OU struct {
	ID        string     `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Status    string     `json:"status"`
}

//

type Endpoint struct {
	EndpointID string     `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name       string     `json:"name"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `sql:"index" json:"deleted_at"`
	Status     string     `json:"status"`
	Parent     string     `sql:"type:uuid;()"`
	Hostname   string     `json:"hostname"`
}

type Credential struct {
	ID        string     `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Status    string     `json:"status"`
	Owner     string     `json:"owner"`
}

type User struct {
	ID           string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Emailaddress string `json:"emailaddress"`
	Password     string `json:"password"` //This isn't staying as type string calm down
}
