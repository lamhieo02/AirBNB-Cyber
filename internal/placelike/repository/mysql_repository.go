package userplacerepository

import (
	"context"
	"go.uber.org/zap"
	placelikemodel "go01-airbnb/internal/placelike/model"
	"go01-airbnb/pkg/common"
	"gorm.io/gorm"
)

type placeLikeRepository struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

func NewPlaceLikeRepository(db *gorm.DB, logger *zap.SugaredLogger) *placeLikeRepository {
	return &placeLikeRepository{db, logger}
}
func (r *placeLikeRepository) Create(ctx context.Context, placeLike *placelikemodel.Like) error {
	db := r.db.Begin()
	if err := db.Table(placeLike.TableName()).Create(placeLike).Error; err != nil {
		db.Rollback()
		r.logger.Desugar().Error("Error when like place", zap.Any("PlaceLike", placeLike), zap.Error(err))
		return common.ErrorDB(err)
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrorDB(err)
	}
	return nil
}
func (r *placeLikeRepository) Delete(ctx context.Context, userId int, placeId int) error {
	if err := r.db.Table(placelikemodel.Like{}.TableName()).Where("place_id = ? and user_id = ?", placeId, userId).Delete(nil).Error; err != nil {
		return common.ErrorDB(err)
	}
	return nil
}
