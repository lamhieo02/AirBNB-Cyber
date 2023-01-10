package amenityhttp

import (
	"context"
	"github.com/gin-gonic/gin"
	amenitymodel "go01-airbnb/internal/amenity/model"
	"go01-airbnb/pkg/common"
	"go01-airbnb/pkg/utils"
	"net/http"
)

type AmenityUseCase interface {
	CreateAmenity(context.Context, *amenitymodel.Amenity) error
	GetAmenities(context.Context, *common.Paging) ([]amenitymodel.Amenity, error)
	DeleteAmenity(context.Context, int) error
	UpdateAmenity(context.Context, int, *amenitymodel.Amenity) error
	GetAmenityById(context.Context, int) (*amenitymodel.Amenity, error)
}

type amenityHandler struct {
	amenityUseCase AmenityUseCase
	hasher         *utils.Hasher
}

func NewAmenityHandler(useCase AmenityUseCase, hasher *utils.Hasher) *amenityHandler {
	return &amenityHandler{amenityUseCase: useCase, hasher: hasher}
}

func (hdl *amenityHandler) CreateAmenity() gin.HandlerFunc {
	// chỉ có host với admin mới có quyền tạo amenity
	return func(ctx *gin.Context) {
		var amenity amenitymodel.Amenity
		if err := ctx.ShouldBind(&amenity); err != nil {
			panic(common.ErrBadRequest(err))
		}

		if err := hdl.amenityUseCase.CreateAmenity(ctx.Request.Context(), &amenity); err != nil {
			panic(err)
		}
		amenity.FakeId = hdl.hasher.Encode(amenity.Id, common.DBTypeAmenity)
		ctx.JSON(http.StatusOK, common.Response(amenity))
	}
}
func (hdl *amenityHandler) GetAmenities() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var paging common.Paging
		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrBadRequest(err))
		}
		paging.Fulfill()

		data, err := hdl.amenityUseCase.GetAmenities(ctx.Request.Context(), &paging)

		if err != nil {
			panic(err)
		}
		for i := range data {
			data[i].FakeId = hdl.hasher.Encode(data[i].Id, common.DBTypeAmenity)
		}

		ctx.JSON(http.StatusOK, common.ResponseWithPaging(data, paging))
	}
}

func (hdl *amenityHandler) DeleteAmenity() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := hdl.hasher.Decode(ctx.Param("id"))

		if err := hdl.amenityUseCase.DeleteAmenity(ctx.Request.Context(), id); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.Response(true))
	}
}

func (hdl *amenityHandler) UpdateAmenity() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var amenity *amenitymodel.Amenity

		if err := ctx.ShouldBind(&amenity); err != nil {
			panic(common.ErrBadRequest(err))
		}
		id := hdl.hasher.Decode(ctx.Param("id"))

		if err := hdl.amenityUseCase.UpdateAmenity(ctx.Request.Context(), id, amenity); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.Response(true))
	}
}

func (hdl *amenityHandler) GetAmenityById() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id := hdl.hasher.Decode(ctx.Param("id"))

		data, err := hdl.amenityUseCase.GetAmenityById(ctx.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		data.FakeId = hdl.hasher.Encode(data.Id, common.DBTypeAmenity)

		ctx.JSON(http.StatusOK, common.Response(data))
	}
}
