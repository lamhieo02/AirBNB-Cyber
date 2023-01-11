package placeamenitiesmodel

import (
	"go01-airbnb/pkg/common"
	"gorm.io/gorm"
	"time"
)

const EntityName = "PlaceAmenities"

// Usecase liet ke cac amenities trong place cu the
type PlaceAmenities struct {
	AmenityId     int       `json:"-" gorm:"column:amenity_id"`
	PlaceId       int       `json:"-" gorm:"column:place_id"`
	FakeAmenityId string    `json:"amenity_id" gorm:"column:-"`
	FakePlaceId   string    `json:"place_id" gorm:"column:-"`
	CreatedAt     time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"column:updated_at"`
	//gorm.DeletedAt: soft delete
	DeletedAt gorm.DeletedAt        `json:"-" gorm:"column:deleted_at"`
	Amenity   *common.SimpleAmenity `json:"amenity,omitempty" gorm:"preload:false"`
}

func (PlaceAmenities) TableName() string {
	return "place_amenities"
}
