package repository

import (
	"context"
	placemodel "go01-airbnb/internal/place/model"
	"go01-airbnb/pkg/common"
	"gorm.io/gorm"
)

type placeRepository struct {
	db *gorm.DB
}

// Constructor
func NewPlaceRepository(db *gorm.DB) *placeRepository {
	return &placeRepository{db}
}

// Create place
func (r *placeRepository) Create(ctx context.Context, place *placemodel.Place) error {
	db := r.db.Begin()
	if err := db.Table(placemodel.Place{}.TableName()).Create(place).Error; err != nil {
		db.Rollback()
		return common.ErrorDB(err)
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrorDB(err)
	}
	return nil
}

// Get list place
func (r *placeRepository) ListDataWithCondition(ctx context.Context, paging *common.Paging, filter *placemodel.Filter,
	keys ...string) ([]placemodel.Place, error) {

	var data []placemodel.Place

	//db := r.db.Table(placemodel.Place{}.TableName())
	db := r.db.Model(&placemodel.Place{})
	if v := filter.OwnerId; v > 0 {
		db = db.Where("owner_id=?", v)
	}
	if v := filter.CityId; v > 0 {
		db = db.Where("city_id=?", v)
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	for k := range keys {
		db = db.Preload(keys[k])
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

// Get place by id
func (r *placeRepository) FindDataWithCondition(ctx context.Context, condition map[string]any) (*placemodel.Place, error) {
	var data placemodel.Place

	if err := r.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrEntityNotFound(placemodel.EntityName, err)
		}
		return nil, common.ErrorDB(err)
	}
	return &data, nil
}
func (r *placeRepository) Update(ctx context.Context, condition map[string]any, place *placemodel.Place) error {
	if err := r.db.Table(placemodel.Place{}.TableName()).Where(condition).Updates(place).Error; err != nil {
		return common.ErrorDB(err)
	}
	return nil
}
func (r *placeRepository) Delete(ctx context.Context, condition map[string]any) error {
	if err := r.db.Table(placemodel.Place{}.TableName()).Where(condition).Delete(&placemodel.Place{}).Error; err != nil {
		return common.ErrorDB(err)
	}
	return nil
}
