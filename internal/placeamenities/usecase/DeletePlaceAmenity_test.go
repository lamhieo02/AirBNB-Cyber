package placeamenitiesusecase

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"go01-airbnb/internal/placeamenities/usecase/mock"
	usermodel "go01-airbnb/internal/user/model"
	"go01-airbnb/pkg/common"
	"testing"
)

func TestPlaceAmenitiesUseCase_DeletePlaceAmenity(t *testing.T) {
	// Cấu hình các dependencies
	ctrl := gomock.NewController(t)
	placeAmtRepoMock := mock.NewMockPlaceAmenitiesRepo(ctrl)
	checkPlOwnerMock := mock.NewMockCheckPlaceOwner(ctrl)
	ctx := context.Background()
	usecase := NewPlaceAmenitiesUseCase(placeAmtRepoMock, checkPlOwnerMock)
	// Chuẩn bị data để test
	placeId := 1
	amentityId := 2
	user := usermodel.User{}
	user.Id = 3
	// Định nghĩa ra các test case cụ thể = Convey
	Convey("Delete a amenity from place", t, func() {
		Convey("user is owner of amenity", func() {
			checkPlOwnerMock.EXPECT().
				CheckOwner(placeId, user.Id).
				Return(nil)
			Convey("delete amenity success", func() {
				placeAmtRepoMock.EXPECT().
					Delete(ctx, map[string]any{"place_id": placeId, "amenity_id": amentityId}).
					Return(nil)
				err := usecase.DeletePlaceAmenity(ctx, placeId, amentityId, &user)
				So(err, ShouldBeNil)
			})
			Convey("delete amenity fail", func() {
				placeAmtRepoMock.EXPECT().
					Delete(ctx, map[string]any{"place_id": placeId, "amenity_id": amentityId}).
					Return(common.ErrCannotDeleteEntity("placeAmenities", errors.New("placeAmenities")))
				err := usecase.DeletePlaceAmenity(ctx, placeId, amentityId, &user)
				So(err, ShouldNotBeNil)
				//So(err, ShouldNotResemble, common.ErrCannotDeleteEntity("PlaceAmenities", err))
			})
		})
		Convey("user is not owner of amenity", func() {
			checkPlOwnerMock.EXPECT().
				CheckOwner(placeId, user.Id).
				Return(errors.New("userId not match with ownerId"))
			//placeAmtRepoMock.EXPECT().
			//	Delete(ctx, map[string]any{"place_id": placeId, "amenity_id": amentityId}).
			//	Return(common.ErrForbidden(errors.New("")))
			err := usecase.DeletePlaceAmenity(ctx, placeId, amentityId, &user)
			So(err, ShouldNotBeNil)
		})
	})
}
