package amenityrepository

import (
	"context"
	"go.uber.org/zap"
	amenitymodel "go01-airbnb/internal/amenity/model"
	"go01-airbnb/pkg/common"
	"gorm.io/gorm"
)

type amenityRepository struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

func NewAmenityRepository(db *gorm.DB, logger *zap.SugaredLogger) *amenityRepository {
	return &amenityRepository{db: db, logger: logger}
}

func (r *amenityRepository) Create(ctx context.Context, amenity *amenitymodel.Amenity) error {
	db := r.db.Begin()
	if err := db.Table(amenity.TableName()).Create(amenity).Error; err != nil {
		db.Rollback()
		r.logger.Desugar().Error("Error when create amenity", zap.Any("Place", amenity), zap.Error(err))
		return common.ErrorDB(err)
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrorDB(err)
	}
	return nil
}
func (r *amenityRepository) ListAmenities(ctx context.Context, paging *common.Paging) ([]amenitymodel.Amenity, error) {
	var data []amenitymodel.Amenity

	db := r.db.Table(amenitymodel.Amenity{}.TableName())

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

func (r *amenityRepository) FindDataWithCondition(ctx context.Context, condition map[string]any) (*amenitymodel.Amenity, error) {
	var data amenitymodel.Amenity

	db := r.db.Table(amenitymodel.Amenity{}.TableName()).Where(condition)

	if err := db.First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrEntityNotFound(amenitymodel.EntityName, err)
		}
		return nil, common.ErrorDB(err)
	}
	return &data, nil
}
func (r *amenityRepository) Delete(ctx context.Context, condition map[string]any) error {
	if err := r.db.Table(amenitymodel.Amenity{}.TableName()).Where(condition).Delete(&amenitymodel.Amenity{}).Error; err != nil {
		return common.ErrorDB(err)
	}
	return nil
}
func (r *amenityRepository) Update(ctx context.Context, condition map[string]any, amenity *amenitymodel.Amenity) error {
	if err := r.db.Table(amenitymodel.Amenity{}.TableName()).Where(condition).Updates(amenity).Error; err != nil {
		return common.ErrorDB(err)
	}
	return nil
}
