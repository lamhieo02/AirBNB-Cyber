package placeamenitiesrepository

import (
	"context"
	"go.uber.org/zap"
	placeamenitiesmodel "go01-airbnb/internal/placeamenities/model"
	"go01-airbnb/pkg/common"
	"gorm.io/gorm"
)

type placeAmenitiesRepository struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

func NewPlaceAmenitiesRepo(db *gorm.DB, logger *zap.SugaredLogger) *placeAmenitiesRepository {
	return &placeAmenitiesRepository{db: db, logger: logger}
}

func (r *placeAmenitiesRepository) Create(ctx context.Context, placeAmenity *placeamenitiesmodel.PlaceAmenities) error {
	db := r.db.Begin()
	if err := db.Table(placeAmenity.TableName()).Create(placeAmenity).Error; err != nil {
		db.Rollback()
		r.logger.Desugar().Error("Error when create place", zap.Any("PlaceAmenity", placeAmenity), zap.Error(err))
		return common.ErrorDB(err)
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrorDB(err)
	}
	return nil
}

func (r *placeAmenitiesRepository) Delete(ctx context.Context, condition map[string]any) error {

	if err := r.db.Table(placeamenitiesmodel.PlaceAmenities{}.TableName()).Where(condition).Delete(nil).Error; err != nil {
		return common.ErrorDB(err)
	}
	return nil
}

func (r *placeAmenitiesRepository) ListDataWithCondition(ctx context.Context,
	condition map[string]any, paging *common.Paging, keys ...string) ([]placeamenitiesmodel.PlaceAmenities, error) {
	var data []placeamenitiesmodel.PlaceAmenities

	//db := r.db.Table(placemodel.Place{}.TableName())
	db := r.db.Model(&placeamenitiesmodel.PlaceAmenities{}).Where(condition)

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	for k := range keys {
		db = db.Preload(keys[k])
	}

	if v := paging.Cursor; v != 0 {
		db = db.Where("place_id > ?", v)
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	if len(data) > 0 {
		paging.NextCursor = data[len(data)-1].PlaceId
	}
	return data, nil
}
