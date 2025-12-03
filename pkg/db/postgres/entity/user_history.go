package entity

import (
	"github.com/google/uuid"
	"github.com/wisaitas/grpc-poc/pkg/db/postgres"
)

type UserHistory struct {
	postgres.BaseEntity
	UserID uuid.UUID `gorm:"column:user_id;type:uuid"`
	Action string    `gorm:"column:action;type:varchar(255);not null"`
}

func (UserHistory) TableName() string {
	return "tbl_user_histories"
}
