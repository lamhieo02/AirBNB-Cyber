package placeusecase

import (
	"context"
	"errors"
	placemodel "go01-airbnb/internal/place/model"
	"go01-airbnb/pkg/common"
	"testing"
)

type mockPlaceRepository struct {
}

// Giả lập các hàm trong place repository
func (m mockPlaceRepository) Create(ctx context.Context, place *placemodel.Place) error {
	if place.Name == "Phuong" {
		return common.ErrorDB(errors.New("something went wrong in db"))
	}
	place.Id = 101
	return nil
}

func (m mockPlaceRepository) Update(ctx context.Context, condition map[string]any, place *placemodel.Place) error {
	return nil
}

func (m mockPlaceRepository) Delete(ctx context.Context, condition map[string]any) error {
	return nil
}

func (m mockPlaceRepository) ListDataWithCondition(ctx context.Context,
	paging *common.Paging, place *placemodel.Filter, keys ...string) ([]placemodel.Place, error) {
	return nil, nil
}

func (m mockPlaceRepository) FindDataWithCondition(ctx context.Context, condition map[string]any, keys ...string) (*placemodel.Place, error) {
	return nil, nil
}

func TestPlaceUseCase_CreatePlace(t *testing.T) {
	placeUC := NewPlaceUseCase(mockPlaceRepository{})
	data := []struct {
		Input    placemodel.Place
		Expected error
	}{
		{Input: placemodel.Place{Name: "", Address: "An Nhon"}, Expected: placemodel.ErrNameIsEmpty},
		{Input: placemodel.Place{Name: "lamheo", Address: ""}, Expected: placemodel.ErrAddressIsEmpty},
		{Input: placemodel.Place{Name: "Phuong", Address: "An Nhon"}, Expected: errors.New("something went wrong in db")},
		{Input: placemodel.Place{Name: "Gia Phuong", Address: "An Nhon"}, Expected: nil},
	}

	for _, item := range data {
		err := placeUC.CreatePlace(context.Background(), &item.Input)
		if err != nil && err.Error() != item.Expected.Error() {
			t.Errorf("create plac - Input: %v, Expected: %v, Output: %v", &item.Input, item.Expected, err)
		}
	}
	// Test trường hợp thành công
	dataTest := placemodel.Place{Name: "Gia Phuong", Address: "An Nhon"}
	err := placeUC.CreatePlace(context.Background(), &dataTest)
	if err != nil {
		t.Errorf("create plac - Input: %v, Expected: %v, Output: %v", &dataTest, nil, err)
	}

}
