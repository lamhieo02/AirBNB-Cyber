package placeamenitieshttp

import (
	"context"
	"github.com/gin-gonic/gin"
	placeamenitiesmodel "go01-airbnb/internal/placeamenities/model"
	"go01-airbnb/pkg/common"
	"go01-airbnb/pkg/utils"
	"net/http"
)

type PlaceAmenitiesUseCase interface {
	CreatePlaceAmenity(context.Context, *placeamenitiesmodel.PlaceAmenities) error
	DeletePlaceAmenity(context.Context, int, int, common.Requester) error
	GetPlaceAmenities(context.Context, *common.Paging) ([]placeamenitiesmodel.PlaceAmenities, error)
	GetAmenitiesByPlaceId(context.Context, int) ([]placeamenitiesmodel.PlaceAmenities, error)
}

type placeAmenitiesHandler struct {
	placeAmenitiesUseCase PlaceAmenitiesUseCase
	hasher                *utils.Hasher
}

func NewPlaceAmenitiesHandler(placeAmenitiesUseCase PlaceAmenitiesUseCase, hasher *utils.Hasher) *placeAmenitiesHandler {
	return &placeAmenitiesHandler{placeAmenitiesUseCase: placeAmenitiesUseCase, hasher: hasher}
}

func (hdl *placeAmenitiesHandler) CreatePlaceAmenities() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var placeAmenity placeamenitiesmodel.PlaceAmenities

		placeId := hdl.hasher.Decode(ctx.Param("pid"))
		amenityId := hdl.hasher.Decode(ctx.Param("aid"))

		placeAmenity.PlaceId = placeId
		placeAmenity.AmenityId = amenityId

		if err := hdl.placeAmenitiesUseCase.CreatePlaceAmenity(ctx.Request.Context(), &placeAmenity); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(true))
	}
}

func (hdl *placeAmenitiesHandler) DeletePlaceAmenities() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("User").(common.Requester)
		placeId := hdl.hasher.Decode(ctx.Param("pid"))
		amenityId := hdl.hasher.Decode(ctx.Param("aid"))

		if err := hdl.placeAmenitiesUseCase.DeletePlaceAmenity(ctx.Request.Context(), placeId, amenityId, user); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(true))
	}
}

func (hdl *placeAmenitiesHandler) GetPlaceAmenities() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var paging common.Paging
		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrBadRequest(err))
		}
		paging.Fulfill()

		result, err := hdl.placeAmenitiesUseCase.GetPlaceAmenities(ctx.Request.Context(), &paging)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].FakeAmenityId = hdl.hasher.Encode(common.DBTypePlaceAmenities, result[i].AmenityId)
			result[i].FakePlaceId = hdl.hasher.Encode(common.DBTypePlaceAmenities, result[i].PlaceId)
		}
		ctx.JSON(http.StatusOK, common.ResponseWithPaging(result, paging))
	}
}

func (hdl *placeAmenitiesHandler) GetAmenitiesByPlaceId() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		placeId := hdl.hasher.Decode(ctx.Param("place_id"))

		result, err := hdl.placeAmenitiesUseCase.GetAmenitiesByPlaceId(ctx.Request.Context(), placeId)
		if err != nil {
			panic(err)
		}
		var data = make([]common.SimpleAmenity, len(result))

		for i := range result {
			data[i] = *result[i].Amenity
			data[i].FakeId = hdl.hasher.Encode(data[i].Id, common.DBTypeAmenity)
		}
		ctx.JSON(http.StatusOK, common.Response(data))
	}
}
