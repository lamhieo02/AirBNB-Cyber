package placelikeusecase

import (
	"context"
	placemodel "go01-airbnb/internal/place/model"
	placelikemodel "go01-airbnb/internal/placelike/model"
	"go01-airbnb/pkg/common"
)

type PlaceLikeRepo interface {
	Create(context.Context, *placelikemodel.Like) error
	Delete(context.Context, int, int) error
	GetPlacesLikedByUser(context.Context, int, ...string) ([]common.SimplePlace, error)
}

type placeLikeUseCase struct {
	placeLikeRepo PlaceLikeRepo
}

func NewUserLikePlaceUseCase(placeLikeRepo PlaceLikeRepo) *placeLikeUseCase {
	return &placeLikeUseCase{placeLikeRepo: placeLikeRepo}
}

func (uc *placeLikeUseCase) LikePlace(ctx context.Context, data *placelikemodel.Like) error {
	if err := uc.placeLikeRepo.Create(ctx, data); err != nil {
		return placelikemodel.ErrCannotLikePlace(err)
	}
	return nil
}

func (uc *placeLikeUseCase) UnLikePlace(ctx context.Context, userId int, placeId int) error {
	if err := uc.placeLikeRepo.Delete(ctx, userId, placeId); err != nil {
		return placelikemodel.ErrCannotUnLikePlace(err)
	}

	return nil
}

func (uc *placeLikeUseCase) PlacesLikedByUser(ctx context.Context, userId int) ([]common.SimplePlace, error) {
	listPlace, err := uc.placeLikeRepo.GetPlacesLikedByUser(ctx, userId, "Place")
	if err != nil {
		return nil, common.ErrCannotListEntity(placemodel.EntityName, err)
	}
	return listPlace, nil
}
