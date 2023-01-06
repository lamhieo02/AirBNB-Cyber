package placelikemodel

type Filter struct {
	PlaceId int `json:"place_id,omitempty" form:"place_id"`
	UserId  int `json:"user_id,omitempty" form:"user_id"`
}
