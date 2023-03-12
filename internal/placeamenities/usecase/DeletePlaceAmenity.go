package placeamenitiesusecase

import (
	"context"
	placeamenitiesmodel "go01-airbnb/internal/placeamenities/model"
	"go01-airbnb/pkg/common"
)

func (uc *placeAmenitiesUseCase) DeletePlaceAmenity(ctx context.Context, placeId int, amenityId int, user common.Requester) error {
	// business: user must be the owner of that place to have the permission to delete the place's amenities
	// check if requester.id == place.ownerId => need to find place.OwnerId && requester.id
	if err := uc.checkPlaceOwner.CheckOwner(placeId, user.GetUserId()); err != nil {
		return common.ErrForbidden(err)
	}

	if err := uc.placeAmenitiesRepo.Delete(ctx, map[string]any{"place_id": placeId, "amenity_id": amenityId}); err != nil {
		return common.ErrCannotDeleteEntity(placeamenitiesmodel.EntityName, err)
	}
	return nil
}
