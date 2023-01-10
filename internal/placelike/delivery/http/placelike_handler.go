package placelikehttp

import (
	"context"
	"github.com/gin-gonic/gin"
	placelikemodel "go01-airbnb/internal/placelike/model"
	"go01-airbnb/pkg/common"
	"go01-airbnb/pkg/utils"
	"net/http"
)

type PlaceLikeUseCase interface {
	LikePlace(context.Context, *placelikemodel.Like) error
	UnLikePlace(context.Context, int, int) error
	PlacesLikedByUser(context.Context, int) ([]common.SimplePlace, error)
}

type placeLikeHandler struct {
	placeLikeUseCase PlaceLikeUseCase
	hasher           *utils.Hasher
}

func NewUserLikePlaceHandler(placeLikeUseCase PlaceLikeUseCase, hasher *utils.Hasher) *placeLikeHandler {
	return &placeLikeHandler{placeLikeUseCase: placeLikeUseCase, hasher: hasher}
}

func (hdl *placeLikeHandler) UserLikePlace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := hdl.hasher.Decode(ctx.Param("id"))
		requester := ctx.MustGet("User").(common.Requester)

		data := placelikemodel.Like{
			PlaceId: id,
			UserId:  requester.GetUserId(),
		}
		if err := hdl.placeLikeUseCase.LikePlace(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(true))
	}
}

func (hdl *placeLikeHandler) UserUnLikePlace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := hdl.hasher.Decode(ctx.Param("id"))
		requester := ctx.MustGet("User").(common.Requester)
		if err := hdl.placeLikeUseCase.UnLikePlace(ctx.Request.Context(), requester.GetUserId(), id); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(true))
	}
}

func (hdl *placeLikeHandler) GetPlacesLikedByUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requester := ctx.MustGet("User").(common.Requester)

		result, err := hdl.placeLikeUseCase.PlacesLikedByUser(ctx.Request.Context(), requester.GetUserId())
		if err != nil {
			panic(err)
		}
		for i := range result {
			result[i].FakeId = hdl.hasher.Encode(result[i].Id, common.DBTypePlace)
		}
		//data[i].Owner.FakeId = hdl.hasher.Encode(v.Owner.Id, common.DBTypeUser)
		ctx.JSON(http.StatusOK, result)
	}
}
