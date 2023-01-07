package placelikemodel

import (
	"go01-airbnb/pkg/common"
	"time"
)

const EntityName = "UserLikePlace"

type Like struct {
	PlaceId   int                 `json:"place_id" gorm:"column:place_id"`
	UserId    int                 `json:"user_id" gorm:"column:user_id"`
	CreatedAt time.Time           `json:"created_at" gorm:"column:created_at"`
	Place     *common.SimplePlace `json:"place" gorm:"preload:false"`
}

func (Like) TableName() string {
	return "place_likes"
}

func ErrCannotLikePlace(err error) *common.AppError {
	return common.NewCustomError(err, "Cannot like this place")
}
func ErrCannotUnLikePlace(err error) *common.AppError {
	return common.NewCustomError(err, "Cannot unlike this place")
}
