package entities

import (
	"errors"
	"time"
)

type ItemCreationDate struct {
	initialized  bool
	creationDate *int64
}

func (cd ItemCreationDate) Initialized() bool {
	return cd.initialized
}

func (cd ItemCreationDate) Time() time.Time {
	return time.Unix(*cd.creationDate, 0)
}

func NewItemCreationDate(creationDate int64) (ItemCreationDate, error) {

	if creationDate <= 0 {
		return ItemCreationDate{}, errors.New("invalid creation date")
	}

	return ItemCreationDate{
		initialized:  true,
		creationDate: &creationDate,
	}, nil

}
