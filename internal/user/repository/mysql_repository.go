package userrepository

import (
	"context"
	"errors"
	usermodel "go01-airbnb/internal/user/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// Constructor
func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, data *usermodel.UserCreate) error {
	db := r.db.Begin()

	if err := db.Table(usermodel.User{}.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return err
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}

func (r *userRepository) FindDataWithCondition(ctx context.Context, condition map[string]any) (*usermodel.User, error) {
	var user usermodel.User

	if err := r.db.Table(usermodel.User{}.TableName()).Where(condition).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("record not found")
		}
		return nil, err
	}
	return &user, nil
}
