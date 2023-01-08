package amenitymodel

import (
	"errors"
	"go01-airbnb/pkg/common"
	"strings"
)

type Amenity struct {
	common.SqlModel
	Name        string        `json:"name" gorm:"column:name"`
	Description string        `json:"description" gorm:"column:description"`
	Icon        *common.Image `json:"icon,omitempty" gorm:"column:icon"`
}

const EntityName = "amenity"

func (Amenity) TableName() string {
	return "amenities"
}

func (a *Amenity) Validate() error {
	a.Name = strings.TrimSpace(a.Name)
	if a.Name == "" {
		return ErrNameIsEmpty
	}
	return nil
}

var (
	ErrNameIsEmpty = common.NewCustomError(errors.New("name can not be blank"), "name can not be blank")
)
