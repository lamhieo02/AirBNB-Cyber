package reviewusecase

import (
	"context"
	placemodel "go01-airbnb/internal/place/model"
	reviewmodel "go01-airbnb/internal/review/model"
	"go01-airbnb/pkg/common"
)

type ReviewRepository interface {
	Create(context.Context, *reviewmodel.Review) error
	Delete(context.Context, map[string]any) error
	FindDataWithCondition(context.Context, map[string]any) (*reviewmodel.Review, error)
	ListDataWithCondition(context.Context, map[string]any, *common.Paging) ([]reviewmodel.Review, error)
	ListAllReviewByPlace(context.Context, int) ([]common.SimpleReview, error)
}

type ReviewUseCase struct {
	reviewRepo ReviewRepository
}

func NewReviewUseCase(reviewRepo ReviewRepository) *ReviewUseCase {
	return &ReviewUseCase{reviewRepo: reviewRepo}
}

func (uc *ReviewUseCase) CreateReview(ctx context.Context, review *reviewmodel.Review) error {
	if err := uc.reviewRepo.Create(ctx, review); err != nil {
		return common.ErrCannotCreateEntity(reviewmodel.EntityName, err)
	}
	return nil
}

func (uc *ReviewUseCase) DeleteReview(ctx context.Context, id int) error {
	_, err := uc.reviewRepo.FindDataWithCondition(ctx, map[string]any{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(reviewmodel.EntityName, err)
	}

	if err := uc.reviewRepo.Delete(ctx, map[string]any{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(placemodel.EntityName, err)
	}

	return nil
}

func (uc *ReviewUseCase) GetAllReview(ctx context.Context, paging *common.Paging) ([]reviewmodel.Review, error) {
	result, err := uc.reviewRepo.ListDataWithCondition(ctx, nil, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(reviewmodel.EntityName, err)
	}
	return result, nil
}

func (uc *ReviewUseCase) GetReviewById(ctx context.Context, id int) (*reviewmodel.Review, error) {
	data, err := uc.reviewRepo.FindDataWithCondition(ctx, map[string]any{"id": id})
	if err != nil {
		return nil, common.ErrCannotListEntity(reviewmodel.EntityName, err)
	}
	return data, nil
}

func (uc *ReviewUseCase) GetAllReviewByPlaceId(ctx context.Context, placeId int) ([]common.SimpleReview, error) {
	data, err := uc.reviewRepo.ListAllReviewByPlace(ctx, placeId)
	if err != nil {
		return nil, common.ErrCannotListEntity(reviewmodel.EntityName, err)
	}
	return data, nil
}
