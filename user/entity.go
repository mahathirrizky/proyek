package user

import "time"

type UserTable struct {
	IdUser            int `gorm:"primaryKey;autoIncrement"`
	Nama           string
	Username       string
	Password   		 string
	Role           string
	Token          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
