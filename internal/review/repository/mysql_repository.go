package reviewrepository

import (
	"context"
	"go.uber.org/zap"
	reviewmodel "go01-airbnb/internal/review/model"
	"go01-airbnb/pkg/common"
	"gorm.io/gorm"
)

type reviewRepository struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

func NewReviewRepository(db *gorm.DB, logger *zap.SugaredLogger) *reviewRepository {
	return &reviewRepository{db, logger}
}

func (r *reviewRepository) Create(ctx context.Context, review *reviewmodel.Review) error {
	db := r.db.Begin()
	if err := db.Table(reviewmodel.Review{}.TableName()).Create(review).Error; err != nil {
		db.Rollback()
		r.logger.Desugar().Error("Error when create review", zap.Any("Review", review), zap.Error(err))
		return common.ErrorDB(err)
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrorDB(err)
	}
	return nil
}

func (r *reviewRepository) FindDataWithCondition(ctx context.Context, condition map[string]any) (*reviewmodel.Review, error) {
	var data reviewmodel.Review

	db := r.db.Table(reviewmodel.Review{}.TableName()).Where(condition)

	if err := db.First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrEntityNotFound(reviewmodel.EntityName, err)
		}
		return nil, common.ErrorDB(err)
	}
	return &data, nil
}

func (r *reviewRepository) Delete(ctx context.Context, condition map[string]any) error {
	if err := r.db.Table(reviewmodel.Review{}.TableName()).Where(condition).Delete(&reviewmodel.Review{}).Error; err != nil {
		return common.ErrorDB(err)
	}
	return nil
}

func (r *reviewRepository) ListDataWithCondition(ctx context.Context, condition map[string]any, paging *common.Paging) ([]reviewmodel.Review, error) {
	var data []reviewmodel.Review

	db := r.db.Table(reviewmodel.Review{}.TableName()).Where(condition)

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	if v := paging.Cursor; v != 0 {
		db = db.Where("id > ?", v)
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	if len(data) > 0 {
		paging.NextCursor = data[len(data)-1].Id
	}
	return data, nil
}

func (r *reviewRepository) ListAllReviewByPlace(ctx context.Context, placeId int) ([]common.SimpleReview, error) {
	var data []common.SimpleReview

	db := r.db.Model(&reviewmodel.Review{}).Select("reviews.rating, reviews.comment, bookings.place_id, users.first_name, users.last_name, users.avatar").Joins("join bookings on reviews.booking_id = bookings.id").Joins("join users on users.id = bookings.user_id")

	db.Where("place_id = ?", placeId).Find(&data)

	return data, nil
}
