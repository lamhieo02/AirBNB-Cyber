package locationhttp

import (
	"context"
	"github.com/gin-gonic/gin"
	locationmodel "go01-airbnb/internal/location/model"
	"go01-airbnb/pkg/common"
	"go01-airbnb/pkg/utils"
	"net/http"
)

type LocationUseCase interface {
	CreateLocation(context.Context, *locationmodel.Location) error
	DeleteLocation(context.Context, int) error
	GetAllLocation(context.Context, *common.Paging) ([]locationmodel.Location, error)
	UpdateLocation(context.Context, int, *locationmodel.Location) error
	GetLocationById(context.Context, int) (*locationmodel.Location, error)
}

type LocationHandler struct {
	locationUseCase LocationUseCase
	hasher          *utils.Hasher
}

func NewLocationHandler(locationUseCase LocationUseCase, hasher *utils.Hasher) *LocationHandler {
	return &LocationHandler{locationUseCase: locationUseCase, hasher: hasher}
}

func (hdl *LocationHandler) CreateLocation() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var location locationmodel.Location
		if err := ctx.ShouldBind(&location); err != nil {
			panic(common.ErrBadRequest(err))
		}
		if err := hdl.locationUseCase.CreateLocation(ctx, &location); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.Response(true))
	}
}

func (hdl *LocationHandler) DeleteLocation() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id := hdl.hasher.Decode(ctx.Param("id"))
		if err := hdl.locationUseCase.DeleteLocation(ctx, id); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.Response(true))
	}
}

func (hdl *LocationHandler) GetAllLocation() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var paging common.Paging
		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrBadRequest(err))
		}
		paging.Fulfill()
		result, err := hdl.locationUseCase.GetAllLocation(ctx, &paging)
		if err != nil {
			panic(err)
		}
		for i := range result {
			result[i].FakeId = hdl.hasher.Encode(result[i].Id, common.DBTypeLocation)
		}
		ctx.JSON(http.StatusOK, common.ResponseWithPaging(result, paging))
	}
}

func (hdl *LocationHandler) UpdateLocation() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id := hdl.hasher.Decode(ctx.Param("id"))
		var location locationmodel.Location
		if err := ctx.ShouldBind(&location); err != nil {
			panic(common.ErrBadRequest(err))
		}
		if err := hdl.locationUseCase.UpdateLocation(ctx, id, &location); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.Response(true))
	}
}

func (hdl *LocationHandler) GetLocationById() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id := hdl.hasher.Decode(ctx.Param("id"))
		data, err := hdl.locationUseCase.GetLocationById(ctx, id)
		if err != nil {
			panic(err)
		}
		data.FakeId = hdl.hasher.Encode(data.Id, common.DBTypeLocation)
		ctx.JSON(http.StatusOK, common.Response(data))
	}
}
