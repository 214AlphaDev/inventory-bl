package entities

import "errors"

type ItemName struct {
	initialized bool
	name        *string
}

func (wn ItemName) Initialized() bool {
	return wn.initialized
}

func (wn ItemName) String() string {
	return *wn.name
}

func NewItemName(name string) (ItemName, error) {

	if len(name) < 5 {
		return ItemName{}, errors.New("item name is too short")
	}

	return ItemName{
		initialized: true,
		name:        &name,
	}, nil
}
