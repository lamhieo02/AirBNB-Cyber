package placelikeusecase

import (
	"context"
	placelikemodel "go01-airbnb/internal/placelike/model"
)

type UserPlaceRepo interface {
	Create(context.Context, *placelikemodel.Like) error
	Delete(context.Context, int, int) error
}

type userPlaceUseCase struct {
	userPlaceRepo UserPlaceRepo
}

func NewUserLikePlaceUseCase(userPlaceRepo UserPlaceRepo) *userPlaceUseCase {
	return &userPlaceUseCase{userPlaceRepo: userPlaceRepo}
}

func (uc *userPlaceUseCase) LikePlace(ctx context.Context, data *placelikemodel.Like) error {
	if err := uc.userPlaceRepo.Create(ctx, data); err != nil {
		return placelikemodel.ErrCannotLikePlace(err)
	}

	return nil
}

func (uc *userPlaceUseCase) UnLikePlace(ctx context.Context, userId int, placeId int) error {
	if err := uc.userPlaceRepo.Delete(ctx, userId, placeId); err != nil {
		return placelikemodel.ErrCannotUnLikePlace(err)
	}

	return nil
}
