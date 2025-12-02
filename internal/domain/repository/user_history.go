package repository

import (
	"github.com/wisaitas/grpc-poc/pkg/db/postgres"
	"github.com/wisaitas/grpc-poc/pkg/db/postgres/entity"
	"gorm.io/gorm"
)

type UserHistoryRepository interface {
	postgres.BaseRepository[entity.UserHistory]
}

type userHistoryRepository struct {
	postgres.BaseRepository[entity.UserHistory]
}

func NewUserHistoryRepository(db *gorm.DB) UserHistoryRepository {
	return &userHistoryRepository{
		BaseRepository: postgres.NewBaseRepository[entity.UserHistory](db),
	}
}
