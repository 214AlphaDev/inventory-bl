package repository

import (
	. "github.com/214alphadev/inventory-bl/entities"
)

type IItemRepository interface {
	Save(item Item) error
	Get(itemID ItemID) (*Item, error)
	Query(from *ItemID, next uint32) ([]Item, error)
	VotesOf(itemID ItemID) (uint32, error)
	Delete(itemID ItemID) error
}
