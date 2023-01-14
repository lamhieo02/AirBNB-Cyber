package locationusecase

import (
	"context"
	locationmodel "go01-airbnb/internal/location/model"
	"go01-airbnb/pkg/common"
)

type LocationRepository interface {
	Create(context.Context, *locationmodel.Location) error
	Delete(context.Context, map[string]any) error
	ListDataWithCondition(context.Context, map[string]any, *common.Paging) ([]locationmodel.Location, error)
	Update(context.Context, map[string]any, *locationmodel.Location) error
	FindDataWithCondition(context.Context, map[string]any) (*locationmodel.Location, error)
}

type LocationUseCase struct {
	locationRepo LocationRepository
}

func NewLocationUseCase(locationRepo LocationRepository) *LocationUseCase {
	return &LocationUseCase{locationRepo: locationRepo}
}

func (uc *LocationUseCase) CreateLocation(ctx context.Context, location *locationmodel.Location) error {
	if err := uc.locationRepo.Create(ctx, location); err != nil {
		return common.ErrCannotCreateEntity(locationmodel.EntityName, err)
	}
	return nil
}

func (uc *LocationUseCase) DeleteLocation(ctx context.Context, id int) error {
	if err := uc.locationRepo.Delete(ctx, map[string]any{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(locationmodel.EntityName, err)
	}
	return nil
}

func (uc *LocationUseCase) GetAllLocation(ctx context.Context, paging *common.Paging) ([]locationmodel.Location, error) {
	data, err := uc.locationRepo.ListDataWithCondition(ctx, nil, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(locationmodel.EntityName, err)
	}
	return data, nil
}

func (uc *LocationUseCase) UpdateLocation(ctx context.Context, id int, location *locationmodel.Location) error {
	if err := uc.locationRepo.Update(ctx, map[string]any{"id": id}, location); err != nil {
		return common.ErrCannotUpdateEntity(locationmodel.EntityName, err)
	}
	return nil
}

func (uc *LocationUseCase) GetLocationById(ctx context.Context, id int) (*locationmodel.Location, error) {
	data, err := uc.locationRepo.FindDataWithCondition(ctx, map[string]any{"id": id})
	if err != nil {
		return nil, common.ErrCannotUpdateEntity(locationmodel.EntityName, err)
	}
	return data, nil
}
