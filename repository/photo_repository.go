package repository

import . "github.com/214alphadev/inventory-bl/entities"

type IItemPhotoRepository interface {
	Save(photo ItemPhoto) error
	DeleteAllFor(item ItemID) error
	Get(item ItemID) (*ItemPhoto, error)
}
