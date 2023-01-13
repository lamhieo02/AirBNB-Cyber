package bookingusecase

import (
	"context"
	bookingmodel "go01-airbnb/internal/booking/model"
	"go01-airbnb/pkg/common"
)

type BookingRepository interface {
	Create(context.Context, *bookingmodel.Booking) error
	ListDataWithCondition(context.Context, map[string]any, *common.Paging) ([]bookingmodel.Booking, error)
	FindDataWithCondition(context.Context, map[string]any) (*bookingmodel.Booking, error)
	Update(context.Context, map[string]any, *bookingmodel.Booking) error
	Delete(context.Context, map[string]any) error
}

type bookingUseCase struct {
	bookingRepo BookingRepository
}

func NewBookingUseCase(bookingRepo BookingRepository) *bookingUseCase {
	return &bookingUseCase{bookingRepo: bookingRepo}
}

func (uc *bookingUseCase) CreateBooking(ctx context.Context, booking *bookingmodel.Booking) error {

	if err := uc.bookingRepo.Create(ctx, booking); err != nil {
		return common.ErrCannotCreateEntity(bookingmodel.EntityName, err)
	}
	return nil
}

func (uc *bookingUseCase) GetAllBooking(ctx context.Context, paging *common.Paging) ([]bookingmodel.Booking, error) {

	data, err := uc.bookingRepo.ListDataWithCondition(ctx, nil, paging)
	if err != nil {
		return nil, common.ErrCannotCreateEntity(bookingmodel.EntityName, err)
	}
	return data, nil
}

func (uc *bookingUseCase) GetBookingById(ctx context.Context, id int) (*bookingmodel.Booking, error) {
	data, err := uc.bookingRepo.FindDataWithCondition(ctx, map[string]any{"id": id})
	if err != nil {
		return nil, common.ErrCannotListEntity(bookingmodel.EntityName, err)
	}
	return data, nil
}

func (uc *bookingUseCase) UpdateBooking(ctx context.Context, id int, booking *bookingmodel.Booking) error {
	if err := uc.bookingRepo.Update(ctx, map[string]any{"id": id}, booking); err != nil {
		return common.ErrCannotUpdateEntity(bookingmodel.EntityName, err)
	}
	return nil
}

func (uc *bookingUseCase) DeleteBooking(ctx context.Context, id int) error {
	if err := uc.bookingRepo.Delete(ctx, map[string]any{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(bookingmodel.EntityName, err)
	}
	return nil
}
