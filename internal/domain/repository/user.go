package repository

import (
	"github.com/wisaitas/grpc-poc/pkg/db/postgres"
	"github.com/wisaitas/grpc-poc/pkg/db/postgres/entity"
	"gorm.io/gorm"
)

// UserRepository สืบทอด method ทั้งหมดจาก BaseRepository[entity.User]
type UserRepository interface {
	postgres.BaseRepository[entity.User]
}

type userRepository struct {
	postgres.BaseRepository[entity.User]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		BaseRepository: postgres.NewBaseRepository[entity.User](db),
	}
}
