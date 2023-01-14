package bookinghttp

import (
	"context"
	"github.com/gin-gonic/gin"
	bookingmodel "go01-airbnb/internal/booking/model"
	"go01-airbnb/pkg/common"
	"go01-airbnb/pkg/utils"
	"net/http"
)

type BookingUseCase interface {
	CreateBooking(context.Context, *bookingmodel.Booking) error
	GetAllBooking(context.Context, *common.Paging) ([]bookingmodel.Booking, error)
	GetBookingById(context.Context, int) (*bookingmodel.Booking, error)
	UpdateBooking(context.Context, int, *bookingmodel.Booking) error
	DeleteBooking(context.Context, int) error
}

type bookingHandler struct {
	bookingUseCase BookingUseCase
	hasher         *utils.Hasher
}

func NewBookingHandler(bookingUseCase BookingUseCase, hasher *utils.Hasher) *bookingHandler {
	return &bookingHandler{bookingUseCase: bookingUseCase, hasher: hasher}
}

func (hdl *bookingHandler) CreateBooking() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requester := ctx.MustGet("User").(common.Requester)
		placeId := hdl.hasher.Decode(ctx.Param("place_id"))
		var booking bookingmodel.Booking

		if err := ctx.ShouldBind(&booking); err != nil {
			panic(common.ErrBadRequest(err))
		}

		booking.UserId = requester.GetUserId()
		booking.PlaceId = placeId

		if err := hdl.bookingUseCase.CreateBooking(ctx, &booking); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.Response(true))
	}
}

func (hdl *bookingHandler) GetAllBooking() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var paging common.Paging
		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrBadRequest(err))
		}
		paging.Fulfill()
		bookings, err := hdl.bookingUseCase.GetAllBooking(ctx, &paging)
		if err != nil {
			panic(err)
		}

		for i := range bookings {
			bookings[i].FakeId = hdl.hasher.Encode(bookings[i].Id, common.DBTypeBooking)
		}
		ctx.JSON(http.StatusOK, common.ResponseWithPaging(bookings, paging))
	}
}

func (hdl *bookingHandler) GetBookingById() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id := hdl.hasher.Decode(ctx.Param("id"))
		booking, err := hdl.bookingUseCase.GetBookingById(ctx, id)
		if err != nil {
			panic(err)
		}

		booking.FakeId = hdl.hasher.Encode(booking.Id, common.DBTypeBooking)

		ctx.JSON(http.StatusOK, common.Response(booking))
	}
}

func (hdl *bookingHandler) UpdateBooking() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id := hdl.hasher.Decode(ctx.Param("id"))
		var booking bookingmodel.Booking
		if err := ctx.ShouldBind(&booking); err != nil {
			panic(common.ErrBadRequest(err))
		}
		if err := hdl.bookingUseCase.UpdateBooking(ctx, id, &booking); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(true))
	}
}

func (hdl *bookingHandler) DeleteBooking() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id := hdl.hasher.Decode(ctx.Param("id"))

		if err := hdl.bookingUseCase.DeleteBooking(ctx, id); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(true))
	}
}
