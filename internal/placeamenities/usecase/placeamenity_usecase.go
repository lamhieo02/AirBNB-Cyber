package placeamenitiesusecase

import (
	"context"
	placeamenitiesmodel "go01-airbnb/internal/placeamenities/model"
	"go01-airbnb/pkg/common"
)

type PlaceAmenitiesRepo interface {
	Create(context.Context, *placeamenitiesmodel.PlaceAmenities) error
	Delete(context.Context, map[string]any) error
	ListDataWithCondition(context.Context, map[string]any, *common.Paging, ...string) ([]placeamenitiesmodel.PlaceAmenities, error)
}

type placeAmenitiesUseCase struct {
	placeAmenitiesRepo PlaceAmenitiesRepo
}

func NewPlaceAmenitiesUseCase(placeAmenitiesRepo PlaceAmenitiesRepo) *placeAmenitiesUseCase {
	return &placeAmenitiesUseCase{placeAmenitiesRepo: placeAmenitiesRepo}
}

func (uc *placeAmenitiesUseCase) CreatePlaceAmenity(ctx context.Context, placeAmenities *placeamenitiesmodel.PlaceAmenities) error {
	if err := uc.placeAmenitiesRepo.Create(ctx, placeAmenities); err != nil {
		return common.ErrCannotCreateEntity(placeamenitiesmodel.EntityName, err)
	}
	return nil
}

func (uc *placeAmenitiesUseCase) DeletePlaceAmenity(ctx context.Context, pid int, aid int) error {
	// business: user must be the owner of that place to have the right to delete the place's amenities
	// check if requester.id == place.ownerId => need to find place.OwnerId

	if err := uc.placeAmenitiesRepo.Delete(ctx, map[string]any{"place_id": pid, "amenity_id": aid}); err != nil {
		return common.ErrCannotDeleteEntity(placeamenitiesmodel.EntityName, err)
	}
	return nil
}

func (uc *placeAmenitiesUseCase) GetPlaceAmenities(ctx context.Context, paging *common.Paging) ([]placeamenitiesmodel.PlaceAmenities, error) {
	result, err := uc.placeAmenitiesRepo.ListDataWithCondition(ctx, nil, paging, "Amenity")
	if err != nil {
		return nil, common.ErrCannotListEntity(placeamenitiesmodel.EntityName, err)
	}
	return result, nil
}

func (uc *placeAmenitiesUseCase) GetAmenitiesByPlaceId(ctx context.Context, placeId int) ([]placeamenitiesmodel.PlaceAmenities, error) {

	result, err := uc.placeAmenitiesRepo.ListDataWithCondition(ctx, map[string]any{"place_id": placeId}, &common.Paging{Limit: 50}, "Amenity")
	if err != nil {
		return nil, common.ErrCannotListEntity(placeamenitiesmodel.EntityName, err)
	}
	return result, nil
}
