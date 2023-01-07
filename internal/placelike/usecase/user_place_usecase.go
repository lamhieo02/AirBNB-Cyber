package placelikeusecase

import (
	"context"
	placemodel "go01-airbnb/internal/place/model"
	placelikemodel "go01-airbnb/internal/placelike/model"
	"go01-airbnb/pkg/common"
)

type UserPlaceRepo interface {
	Create(context.Context, *placelikemodel.Like) error
	Delete(context.Context, int, int) error
	GetPlacesLikedByUser(context.Context, int, ...string) ([]common.SimplePlace, error)
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

func (uc *userPlaceUseCase) PlacesLikedByUser(ctx context.Context, userId int) ([]common.SimplePlace, error) {
	listPlace, err := uc.userPlaceRepo.GetPlacesLikedByUser(ctx, userId, "Place")
	if err != nil {
		return nil, common.ErrCannotListEntity(placemodel.EntityName, err)
	}
	return listPlace, nil
}
