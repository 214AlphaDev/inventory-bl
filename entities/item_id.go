package entities

import "github.com/satori/go.uuid"

type ItemID struct {
	initialized bool
	id          *uuid.UUID
}

func (wi ItemID) Initialized() bool {
	return wi.initialized
}

func (wi ItemID) String() string {
	return wi.id.String()
}

func NewItemID(itemID string) (ItemID, error) {

	id, err := uuid.FromString(itemID)
	if err != nil {
		return ItemID{}, err
	}

	return ItemID{
		id:          &id,
		initialized: true,
	}, nil

}
