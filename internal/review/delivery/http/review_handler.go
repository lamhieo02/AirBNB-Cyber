package reviewhttp

import (
	"context"
	"github.com/gin-gonic/gin"
	reviewmodel "go01-airbnb/internal/review/model"
	"go01-airbnb/pkg/common"
	"go01-airbnb/pkg/utils"
	"net/http"
)

type ReviewUseCase interface {
	CreateReview(context.Context, *reviewmodel.Review) error
	DeleteReview(context.Context, int) error
	GetAllReview(context.Context, *common.Paging) ([]reviewmodel.Review, error)
	GetReviewById(context.Context, int) (*reviewmodel.Review, error)
	GetAllReviewByPlaceId(context.Context, int) ([]common.SimpleReview, error)
}

type reviewHandler struct {
	reviewUseCase ReviewUseCase
	hasher        *utils.Hasher
}

func NewReviewHandler(reviewUseCase ReviewUseCase, hasher *utils.Hasher) *reviewHandler {
	return &reviewHandler{reviewUseCase: reviewUseCase, hasher: hasher}
}

func (hdl *reviewHandler) CreateReview() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var review reviewmodel.Review
		if err := ctx.ShouldBind(&review); err != nil {
			panic(common.ErrBadRequest(err))
		}
		bookingId := hdl.hasher.Decode(ctx.Param("booking_id"))
		review.BookingId = bookingId
		if err := hdl.reviewUseCase.CreateReview(ctx, &review); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.Response(true))
	}
}

func (hdl *reviewHandler) DeleteReview() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := hdl.hasher.Decode(ctx.Param("id"))
		if err := hdl.reviewUseCase.DeleteReview(ctx, id); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.Response(true))
	}
}

func (hdl *reviewHandler) GetAllReview() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging common.Paging
		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrBadRequest(err))
		}
		paging.Fulfill()
		result, err := hdl.reviewUseCase.GetAllReview(ctx, &paging)
		if err != nil {
			panic(err)
		}
		for i := range result {
			result[i].FakeId = hdl.hasher.Encode(result[i].Id, common.DBTypeReview)
		}
		ctx.JSON(http.StatusOK, common.ResponseWithPaging(result, paging))
	}
}

func (hdl *reviewHandler) GetReviewById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := hdl.hasher.Decode(ctx.Param("id"))
		result, err := hdl.reviewUseCase.GetReviewById(ctx, id)
		if err != nil {
			panic(err)
		}
		result.FakeId = hdl.hasher.Encode(result.Id, common.DBTypeReview)

		ctx.JSON(http.StatusOK, common.Response(result))
	}
}

func (hdl *reviewHandler) GetAllReviewByPlaceId() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		placeId := hdl.hasher.Decode(ctx.Param("place_id"))
		result, err := hdl.reviewUseCase.GetAllReviewByPlaceId(ctx, placeId)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.Response(result))
	}
}
