package entities

import "errors"

type ItemDescription struct {
	initialized bool
	description *string
}

func (wd ItemDescription) Initialized() bool {
	return wd.initialized
}

func (wd ItemDescription) String() string {
	return *wd.description
}

func NewItemDescription(description string) (ItemDescription, error) {

	if len(description) < 30 {
		return ItemDescription{}, errors.New("item description is too short")
	}

	return ItemDescription{
		initialized: true,
		description: &description,
	}, nil

}
