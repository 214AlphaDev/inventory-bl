package entities

import (
	"github.com/214alphadev/inventory-bl/utils"
	. "github.com/214alphadev/inventory-bl/value_objects"
)

type Vote struct {
	initialized bool
	itemID      ItemID
	memberID    MemberID
}

func (v Vote) Initialized() bool {
	return v.initialized
}

func (v Vote) MemberID() MemberID {
	return v.memberID
}

func (v Vote) ItemID() ItemID {
	return v.itemID
}

func NewVote(memberID MemberID, itemID ItemID) (Vote, error) {

	if err := utils.Initialized(memberID, itemID); err != nil {
		return Vote{}, err
	}

	return Vote{
		initialized: true,
		itemID:      itemID,
		memberID:    memberID,
	}, nil

}
