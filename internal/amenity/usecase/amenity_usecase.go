package amenityusecase

import (
	"context"
	amenitymodel "go01-airbnb/internal/amenity/model"
	"go01-airbnb/pkg/common"
)

type AmenityRepository interface {
	Create(context.Context, *amenitymodel.Amenity) error
	ListAmenities(context.Context, *common.Paging) ([]amenitymodel.Amenity, error)
	FindDataWithCondition(context.Context, map[string]any) (*amenitymodel.Amenity, error)
	Delete(context.Context, map[string]any) error
	Update(context.Context, map[string]any, *amenitymodel.Amenity) error
}

type AmenityUseCase struct {
	amenityRepo AmenityRepository
}

func NewAmenityUseCase(amenityRepo AmenityRepository) *AmenityUseCase {
	return &AmenityUseCase{amenityRepo: amenityRepo}
}

func (uc *AmenityUseCase) CreateAmenity(ctx context.Context, amenity *amenitymodel.Amenity) error {
	if err := amenity.Validate(); err != nil {
		return common.ErrBadRequest(err)
	}
	if err := uc.amenityRepo.Create(ctx, amenity); err != nil {
		return common.ErrCannotCreateEntity(amenitymodel.EntityName, err)
	}
	return nil
}

func (uc *AmenityUseCase) GetAmenities(ctx context.Context, paging *common.Paging) ([]amenitymodel.Amenity, error) {
	data, err := uc.amenityRepo.ListAmenities(ctx, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(amenitymodel.EntityName, err)
	}
	return data, nil
}

func (uc *AmenityUseCase) GetAmenityById(ctx context.Context, id int) (*amenitymodel.Amenity, error) {
	data, err := uc.amenityRepo.FindDataWithCondition(ctx, map[string]any{"id": id})
	if err != nil {
		return nil, common.ErrCannotListEntity(amenitymodel.EntityName, err)
	}
	return data, nil
}

func (uc *AmenityUseCase) DeleteAmenity(ctx context.Context, id int) error {
	_, err := uc.amenityRepo.FindDataWithCondition(ctx, map[string]any{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(amenitymodel.EntityName, err)
	}
	if err := uc.amenityRepo.Delete(ctx, map[string]any{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(amenitymodel.Amenity{}.TableName(), err)
	}
	return nil
}
func (uc *AmenityUseCase) UpdateAmenity(ctx context.Context, id int, amenity *amenitymodel.Amenity) error {
	_, err := uc.amenityRepo.FindDataWithCondition(ctx, map[string]any{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(amenitymodel.EntityName, err)
	}
	if err := uc.amenityRepo.Update(ctx, map[string]any{"id": id}, amenity); err != nil {
		return common.ErrCannotUpdateEntity(amenitymodel.Amenity{}.TableName(), err)
	}
	return nil
}
