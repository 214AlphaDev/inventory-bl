package entities

import (
	"github.com/214alphadev/inventory-bl/utils"
	. "github.com/214alphadev/inventory-bl/value_objects"
)

type Item struct {
	initialized bool
	id          *ItemID
	name        *ItemName
	description *ItemDescription
	story       *ItemStory
	creator     *MemberID
	createdAt   *ItemCreationDate
	category    *Category
}

func (i Item) Initialized() bool {
	return i.initialized
}

func (i Item) ID() ItemID {
	return *i.id
}

func (i Item) Name() ItemName {
	return *i.name
}

func (i Item) Description() ItemDescription {
	return *i.description
}

func (i Item) Story() ItemStory {
	return *i.story
}

func (i Item) Creator() MemberID {
	return *i.creator
}

func (i *Item) ChangeName(name ItemName) error {

	err := utils.Initialized(name)
	if err != nil {
		return err
	}

	i.name = &name
	return nil

}

func (i *Item) ChangeDescription(description ItemDescription) error {

	err := utils.Initialized(description)
	if err != nil {
		return err
	}

	i.description = &description
	return nil

}

func (i *Item) ChangeStory(story ItemStory) error {

	if err := utils.Initialized(story); err != nil {
		return err
	}

	i.story = &story
	return nil

}

func (i Item) CreatedAt() ItemCreationDate {
	return *i.createdAt
}

func (i Item) Category() Category {
	return *i.category
}

func (i *Item) ChangeCategory(category Category) {
	i.category = &category
}

func NewItem(creator MemberID, id ItemID, name ItemName, description ItemDescription, story ItemStory, creationDate ItemCreationDate, category Category) (Item, error) {

	if err := utils.Initialized(creator, id, name, description, story, creationDate); err != nil {
		return Item{}, err
	}

	return Item{
		id:          &id,
		name:        &name,
		description: &description,
		story:       &story,
		initialized: true,
		createdAt:   &creationDate,
		creator:     &creator,
		category:    &category,
	}, nil

}
