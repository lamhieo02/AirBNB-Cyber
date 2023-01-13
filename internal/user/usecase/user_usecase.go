package userusecase

import (
	"context"
	"go01-airbnb/config"
	usermodel "go01-airbnb/internal/user/model"
	"go01-airbnb/pkg/common"
	"go01-airbnb/pkg/utils"
)

type UserRepository interface {
	Create(context.Context, *usermodel.UserRegister) error
	FindDataWithCondition(context.Context, map[string]any) (*usermodel.User, error)
	Update(context.Context, *usermodel.UserUpdate, map[string]any) error
}

type userUseCase struct {
	userRepository UserRepository
	cfg            *config.Config
}

func NewUserUseCase(cfg *config.Config, userRepo UserRepository) *userUseCase {
	return &userUseCase{userRepo, cfg}
}

func (u *userUseCase) Register(ctx context.Context, data *usermodel.UserRegister) error {
	// Check email have been existed ?
	user, _ := u.userRepository.FindDataWithCondition(ctx, map[string]any{"email": data.Email})
	if user != nil {
		return usermodel.ErrEmailExisted
	}
	//Validate data
	if err := data.Validate(); err != nil {
		return err
	}

	// Prepare data before create user
	// Must hash value of password before store in db
	if err := data.PrepareCreate(); err != nil {
		return err
	}

	if err := u.userRepository.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}

func (u *userUseCase) Login(ctx context.Context, data *usermodel.UserLogin) (*utils.Token, error) {
	// Check email have been existed ?
	// B1: find user by email
	user, err := u.userRepository.FindDataWithCondition(ctx, map[string]any{"email": data.Email})
	if err != nil {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}
	// B2: Compare password of user with hashed password in db
	if err := utils.Compare(user.Password, data.Password); err != nil {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	// B3: Generate toke
	token, err := utils.GenerateJWT(utils.TokenPayload{user.Email, user.Role}, u.cfg)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return token, nil
}

func (u *userUseCase) UpdateProfile(ctx context.Context, data *usermodel.UserUpdate, userEmail string) error {
	if err := u.userRepository.Update(ctx, data, map[string]any{"email": userEmail}); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}
	return nil
}
