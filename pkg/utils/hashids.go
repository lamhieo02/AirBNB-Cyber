package utils

import (
	"github.com/speps/go-hashids/v2"
)

type Hasher struct {
	HashID *hashids.HashID
}

func NewHashIds(salt string, minLength int) *Hasher {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLength

	h, _ := hashids.NewWithData(hd)

	return &Hasher{h}
}

func (h *Hasher) Encode(id, dbType int) string {
	encode, _ := h.HashID.Encode([]int{id, dbType})
	return encode
}

func (h *Hasher) Decode(encode string) int {

	decode, err := h.HashID.DecodeWithError(encode)

	if err != nil {
		return 0
	}

	return decode[0]
}
