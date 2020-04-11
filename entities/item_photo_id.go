package entities

import "github.com/satori/go.uuid"

type ItemPhotoID struct {
	initialized bool
	id          *uuid.UUID
}

func (id ItemPhotoID) Initialized() bool {
	return id.initialized
}

func (id ItemPhotoID) String() string {
	return id.id.String()
}

func NewItemPhotoID(id string) (ItemPhotoID, error) {

	i, err := uuid.FromString(id)
	if err != nil {
		return ItemPhotoID{}, err
	}

	return ItemPhotoID{
		initialized: true,
		id:          &i,
	}, nil

}
