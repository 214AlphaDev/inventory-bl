package entities

import "fmt"

type ItemStory struct {
	initialized bool
	story       *string
}

func (ws ItemStory) Initialized() bool {
	return ws.initialized
}

func (ws ItemStory) String() string {
	return *ws.story
}

func (ws ItemStory) IsNil() bool {
	return ws.story == nil
}

func NewItemStory(story *string) (ItemStory, error) {

	switch story {
	case nil:
		return ItemStory{initialized: true}, nil
	default:

		if len(*story) < 20 {
			return ItemStory{}, fmt.Errorf("item story is too short")
		}

		return ItemStory{initialized: true, story: &*story}, nil

	}

}
