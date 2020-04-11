package entities

import (
	"encoding/base64"
	"errors"
	"github.com/214alphadev/inventory-bl/utils"
)

type ItemPhoto struct {
	id          *ItemPhotoID
	photo       *[]byte
	item        *ItemID
	initialized bool
}

func (ip ItemPhoto) Initialized() bool {
	return ip.initialized
}

func (ip ItemPhoto) Item() ItemID {
	return *ip.item
}

func (ip ItemPhoto) String() string {
	return base64.StdEncoding.EncodeToString(*ip.photo)
}

func (ip ItemPhoto) ID() ItemPhotoID {
	return *ip.id
}

func NewItemPhoto(id ItemPhotoID, item ItemID, photo []byte) (ItemPhoto, error) {

	if err := utils.Initialized(id, item); err != nil {
		return ItemPhoto{}, err
	}

	if photo == nil || len(photo) == 0 {
		return ItemPhoto{}, errors.New("invalid photo")
	}

	return ItemPhoto{
		id:          &id,
		item:        &item,
		initialized: true,
		photo:       &photo,
	}, nil

}
