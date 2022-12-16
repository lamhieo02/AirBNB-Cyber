package placehttp

import (
	"context"
	"github.com/gin-gonic/gin"
	placemodel "go01-airbnb/internal/place/model"
	"go01-airbnb/pkg/common"
	"go01-airbnb/pkg/utils"
	"net/http"
)

type PlaceUseCase interface {
	CreatePlace(context.Context, *placemodel.Place) error
	GetPlaces(context.Context, *common.Paging, *placemodel.Filter) ([]placemodel.Place, error)
	GetPlaceById(context.Context, int) (*placemodel.Place, error)
	UpdatePlace(context.Context, common.Requester, int, *placemodel.Place) error
	DeletePlaceById(context.Context, common.Requester, int) error
}
type placeHandler struct {
	placeUseCase PlaceUseCase
	hasher       *utils.Hasher
}

func NewPlaceHandler(placeUseCase PlaceUseCase, hasher *utils.Hasher) *placeHandler {
	return &placeHandler{placeUseCase, hasher}
}

func (hdl *placeHandler) CreatePlace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var place placemodel.Place

		if err := ctx.ShouldBind(&place); err != nil {
			panic(common.ErrBadRequest(err))
		}
		if err := hdl.placeUseCase.CreatePlace(ctx.Request.Context(), &place); err != nil {
			panic(err)
		}
		// encode id before send data to user
		place.FakeId = hdl.hasher.Encode(place.Id, 1)
		ctx.JSON(http.StatusOK, common.Response(place))
	}
}

func (hdl *placeHandler) GetPlaces() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//paging
		var paging common.Paging

		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrBadRequest(err))
		}

		paging.Fulfill()
		//filter
		var filter placemodel.Filter
		if err := ctx.ShouldBind(&filter); err != nil {
			panic(common.ErrBadRequest(err))
		}
		data, err := hdl.placeUseCase.GetPlaces(ctx.Request.Context(), &paging, &filter)

		if err != nil {
			//fmt.Println(err.(*common.AppError).RootErr())
			panic(err)
		}
		for i, v := range data {
			data[i].FakeId = hdl.hasher.Encode(v.Id, 1)
			data[i].Owner.FakeId = hdl.hasher.Encode(v.Owner.Id, 2)
		}
		ctx.JSON(http.StatusOK, common.ResponseWithPaging(data, paging))
	}
}

func (hdl *placeHandler) GetPlaceById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := hdl.hasher.Decode(ctx.Param("id"))

		data, err := hdl.placeUseCase.GetPlaceById(ctx.Request.Context(), id)

		if err != nil {
			panic(err)
		}
		data.FakeId = hdl.hasher.Encode(data.Id, 1)
		ctx.JSON(http.StatusOK, common.Response(data))
	}
}
func (hdl *placeHandler) UpdatePlace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get info request
		requester := ctx.MustGet("User").(common.Requester)

		id := hdl.hasher.Decode(ctx.Param("id"))

		var place placemodel.Place

		if err := ctx.ShouldBind(&place); err != nil {
			panic(common.ErrBadRequest(err))
		}
		if err := hdl.placeUseCase.UpdatePlace(ctx.Request.Context(), requester, id, &place); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.Response(place))
	}
}
func (hdl *placeHandler) DeletePlace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get info request
		requester := ctx.MustGet("User").(common.Requester)

		id := hdl.hasher.Decode(ctx.Param("id"))

		if err := hdl.placeUseCase.DeletePlaceById(ctx.Request.Context(), requester, id); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.Response(true))
	}
}
