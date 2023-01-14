package locationrepository

import (
	"context"
	"go.uber.org/zap"
	locationmodel "go01-airbnb/internal/location/model"
	"go01-airbnb/pkg/common"
	"gorm.io/gorm"
)

type LocationRepository struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

func NewLocationRepository(db *gorm.DB, logger *zap.SugaredLogger) *LocationRepository {
	return &LocationRepository{db: db, logger: logger}
}

func (r *LocationRepository) Create(ctx context.Context, location *locationmodel.Location) error {
	db := r.db.Begin()
	if err := db.Table(location.TableName()).Create(location).Error; err != nil {
		db.Rollback()
		r.logger.Desugar().Error("Error when create location", zap.Any("Location", location), zap.Error(err))
		return common.ErrorDB(err)
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrorDB(err)
	}
	return nil
}

func (r *LocationRepository) Delete(ctx context.Context, condition map[string]any) error {
	if err := r.db.Table(locationmodel.Location{}.TableName()).Where(condition).Delete(&locationmodel.Location{}).Error; err != nil {
		return common.ErrorDB(err)
	}
	return nil
}
func (r *LocationRepository) Update(ctx context.Context, condition map[string]any, location *locationmodel.Location) error {
	if err := r.db.Table(locationmodel.Location{}.TableName()).Where(condition).Updates(location).Error; err != nil {
		return common.ErrorDB(err)
	}
	return nil
}

func (r *LocationRepository) ListDataWithCondition(ctx context.Context, condition map[string]any, paging *common.Paging) ([]locationmodel.Location, error) {
	var data []locationmodel.Location

	db := r.db.Table(locationmodel.Location{}.TableName()).Where(condition)

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

func (r *LocationRepository) FindDataWithCondition(ctx context.Context,
	condition map[string]any) (*locationmodel.Location, error) {
	var data locationmodel.Location

	db := r.db.Table(locationmodel.Location{}.TableName()).Where(condition)
	if err := db.First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrEntityNotFound(locationmodel.EntityName, err)
		}
		return nil, common.ErrorDB(err)
	}
	return &data, nil
}
