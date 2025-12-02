package entity

import "github.com/wisaitas/grpc-poc/pkg/db/postgres"

type User struct {
	postgres.BaseEntity
	FirstName string `gorm:"column:first_name;type:varchar(255);not null"`
	LastName  string `gorm:"column:last_name;type:varchar(255);not null"`
	Email     string `gorm:"column:email;type:varchar(255);not null"`
	Password  string `gorm:"column:password;type:varchar(255);not null"`
}
