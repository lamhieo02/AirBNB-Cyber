package userplacehttp

import (
	"context"
	"github.com/gin-gonic/gin"
	placelikemodel "go01-airbnb/internal/placelike/model"
	"go01-airbnb/pkg/common"
	"go01-airbnb/pkg/utils"
	"net/http"
)

type UserPlaceUseCase interface {
	LikePlace(context.Context, *placelikemodel.Like) error
	UnLikePlace(context.Context, int, int) error
}

type userLikePlaceHandler struct {
	userLikeResUseCase UserPlaceUseCase
	hasher             *utils.Hasher
}

func NewUserLikePlaceUseCase(userLikeResUseCase UserPlaceUseCase, hasher *utils.Hasher) *userLikePlaceHandler {
	return &userLikePlaceHandler{userLikeResUseCase: userLikeResUseCase, hasher: hasher}
}

func (hdl *userLikePlaceHandler) UserLikePlace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := hdl.hasher.Decode(ctx.Param("id"))
		requester := ctx.MustGet("User").(common.Requester)

		data := placelikemodel.Like{
			PlaceId: id,
			UserId:  requester.GetUserId(),
		}
		if err := hdl.userLikeResUseCase.LikePlace(ctx, &data); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(true))
	}
}

func (hdl *userLikePlaceHandler) UserUnLikePlace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := hdl.hasher.Decode(ctx.Param("id"))
		requester := ctx.MustGet("User").(common.Requester)
		if err := hdl.userLikeResUseCase.UnLikePlace(ctx, requester.GetUserId(), id); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.Response(true))
	}
}
