package bookingrepository

import (
	"context"
	"go.uber.org/zap"
	bookingmodel "go01-airbnb/internal/booking/model"
	placemodel "go01-airbnb/internal/place/model"
	"go01-airbnb/pkg/common"
	"gorm.io/gorm"
)

type bookingRepository struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

func NewBookingRepository(db *gorm.DB, logger *zap.SugaredLogger) *bookingRepository {
	return &bookingRepository{db: db, logger: logger}
}

func (r *bookingRepository) Create(ctx context.Context, booking *bookingmodel.Booking) error {
	db := r.db.Begin()
	if err := db.Table(booking.TableName()).Create(booking).Error; err != nil {
		db.Rollback()
		r.logger.Desugar().Error("Error when create booking", zap.Any("Booking", booking), zap.Error(err))
		return common.ErrorDB(err)
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrorDB(err)
	}
	return nil
}

func (r *bookingRepository) ListDataWithCondition(ctx context.Context, condition map[string]any, paging *common.Paging) ([]bookingmodel.Booking, error) {
	var data []bookingmodel.Booking

	db := r.db.Table(bookingmodel.Booking{}.TableName()).Where(condition)
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

func (r *bookingRepository) FindDataWithCondition(ctx context.Context, condition map[string]any) (*bookingmodel.Booking, error) {
	var data bookingmodel.Booking

	db := r.db.Table(bookingmodel.Booking{}.TableName()).Where(condition)

	if err := db.First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrEntityNotFound(placemodel.EntityName, err)
		}
		return nil, common.ErrorDB(err)
	}
	return &data, nil
}

func (r *bookingRepository) Update(ctx context.Context, condition map[string]any, booking *bookingmodel.Booking) error {
	if err := r.db.Table(bookingmodel.Booking{}.TableName()).Where(condition).Updates(booking).Error; err != nil {
		return common.ErrorDB(err)
	}
	return nil
}

func (r *bookingRepository) Delete(ctx context.Context, condition map[string]any) error {
	if err := r.db.Table(bookingmodel.Booking{}.TableName()).Where(condition).Delete(&bookingmodel.Booking{}).Error; err != nil {
		return common.ErrorDB(err)
	}
	return nil
}
