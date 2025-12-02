package postgres

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

// txKey is a context key for storing the transaction
type txKey struct{}

type BaseRepository[T any] interface {
	Create(ctx context.Context, entity *T) error
	FindByID(ctx context.Context, id any) (*T, error)
	FindOne(ctx context.Context, query interface{}, args ...interface{}) (*T, error)
	FindAll(ctx context.Context) ([]T, error)
	Update(ctx context.Context, entity *T) error
	Delete(ctx context.Context, id any) error
	WithTx(tx *gorm.DB) BaseRepository[T]
	// DoTransaction executes the given function within a transaction
	Transaction(ctx context.Context, fn func(ctx context.Context) error) error
}

type baseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{
		db: db,
	}
}

// getDB extracts transaction from context or returns default db
func (r *baseRepository[T]) getDB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txKey{}).(*gorm.DB); ok {
		return tx
	}
	return r.db.WithContext(ctx)
}

func (r *baseRepository[T]) Create(ctx context.Context, entity *T) error {
	return r.getDB(ctx).Create(entity).Error
}

func (r *baseRepository[T]) FindByID(ctx context.Context, id any) (*T, error) {
	var entity T
	if err := r.getDB(ctx).First(&entity, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *baseRepository[T]) FindOne(ctx context.Context, query interface{}, args ...interface{}) (*T, error) {
	var entity T
	if err := r.getDB(ctx).Where(query, args...).First(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *baseRepository[T]) FindAll(ctx context.Context) ([]T, error) {
	var entities []T
	if err := r.getDB(ctx).Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *baseRepository[T]) Update(ctx context.Context, entity *T) error {
	return r.getDB(ctx).Save(entity).Error
}

func (r *baseRepository[T]) Delete(ctx context.Context, id any) error {
	var entity T
	result := r.getDB(ctx).Delete(&entity, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return nil
}

func (r *baseRepository[T]) WithTx(tx *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{
		db: tx,
	}
}

func (r *baseRepository[T]) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		c := context.WithValue(ctx, txKey{}, tx)
		return fn(c)
	})
}
