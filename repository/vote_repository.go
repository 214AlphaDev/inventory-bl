package repository

import (
	. "github.com/214alphadev/inventory-bl/entities"
)

type IVoteRepository interface {
	Save(v Vote) error
	DoesExist(v Vote) (bool, error)
	Delete(v Vote) error
}
