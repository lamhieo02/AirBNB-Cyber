package placemodel

import "testing"

type testData struct {
	Input    Place
	Expected error
}

func TestPlace_Validate(t *testing.T) {

	data := []testData{
		{Input: Place{Name: "", Address: "Quy Nhon"}, Expected: ErrNameIsEmpty},
		{Input: Place{Name: "Gia Phuong", Address: ""}, Expected: ErrAddressIsEmpty},
		{Input: Place{Name: "Lam", Address: "Binh Dinh"}, Expected: nil},
	}
	for _, item := range data {
		err := item.Input.Validate()
		if err != item.Expected {
			t.Errorf("validate place: Input %v, Expected: %v, Output: %v", item.Input, item.Expected, err)
		}
	}
}
