package value_objects

import "github.com/satori/go.uuid"

type MemberID struct {
	initialized bool
	id          *uuid.UUID
}

func (mid MemberID) Initialized() bool {
	return mid.initialized
}

func (mid MemberID) String() string {
	return mid.id.String()
}

func NewMemberID(memberID string) (MemberID, error) {

	id, err := uuid.FromString(memberID)
	if err != nil {
		return MemberID{}, err
	}

	return MemberID{
		initialized: true,
		id:          &id,
	}, nil
}
