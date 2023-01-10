package placelikerepository

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
func (r *placeLikeRepository) GetPlacesLikedByUser(ctx context.Context, userId int, keys ...string) ([]common.SimplePlace, error) {
	var result []placelikemodel.Like

	//db := r.db.Model(&placelikemodel.Like{})
	db := r.db.Table(placelikemodel.Like{}.TableName()).Where("user_id = ?", userId)

	//Preload more keys
	for _, k := range keys {
		db = db.Preload(k)
	}

	//if err := db.Count(&paging.Total).Error; err != nil {
	//	return nil, common.ErrorDB(err)
	//}

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	places := make([]common.SimplePlace, len(result))
	for i := range result {
		places[i] = *result[i].Place
		places[i].CreatedAt = result[i].CreatedAt
	}
	return places, nil
}
