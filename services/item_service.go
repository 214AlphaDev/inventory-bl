package services

import (
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	. "github.com/214alphadev/inventory-bl/entities"
	. "github.com/214alphadev/inventory-bl/repository"
	"github.com/214alphadev/inventory-bl/utils"
	. "github.com/214alphadev/inventory-bl/value_objects"
	"time"
)

type IItemService interface {
	Create(memberID MemberID, name ItemName, description ItemDescription, story ItemStory, category Category) (ItemID, error)
	Vote(itemID ItemID, memberID MemberID) error
	WithdrawVote(itemID ItemID, memberID MemberID) error
	Items(start *ItemID, next uint32) ([]Item, error)
	GetByID(itemID ItemID) (*Item, error)
	VotesOfItem(itemID ItemID) (uint32, error)
	Update(item Item) error
	Delete(item ItemID) error
	SetPhoto(item ItemID, rawPhoto []byte) error
	VotedOnItem(member MemberID, item ItemID) (bool, error)
	GetPhoto(item ItemID) (*ItemPhoto, error)
}

type itemService struct {
	itemRepository      IItemRepository
	voteRepository      IVoteRepository
	itemPhotoRepository IItemPhotoRepository
}

func (ws *itemService) Create(memberID MemberID, name ItemName, description ItemDescription, story ItemStory, category Category) (ItemID, error) {

	if err := utils.Initialized(memberID, name, description, story); err != nil {
		return ItemID{}, err
	}

	id, err := NewItemID(uuid.NewV4().String())
	if err != nil {
		return ItemID{}, err
	}

	creationDate, err := NewItemCreationDate(time.Now().Unix())
	if err != nil {
		return ItemID{}, err
	}

	item, err := NewItem(memberID, id, name, description, story, creationDate, category)
	if err != nil {
		return ItemID{}, err
	}

	if err := ws.itemRepository.Save(item); err != nil {
		return ItemID{}, err
	}

	return id, nil

}

func (ws *itemService) Vote(itemID ItemID, memberID MemberID) error {

	if err := utils.Initialized(itemID, memberID); err != nil {
		return err
	}

	fetchedItem, err := ws.itemRepository.Get(itemID)
	if err != nil {
		return err
	}
	if fetchedItem == nil {
		return fmt.Errorf("ItemDoesNotExist")
	}

	vote, err := NewVote(memberID, itemID)
	if err != nil {
		return err
	}

	exists, err := ws.voteRepository.DoesExist(vote)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("AlreadyVoted")
	}

	return ws.voteRepository.Save(vote)

}

func (ws *itemService) WithdrawVote(itemID ItemID, memberID MemberID) error {

	if err := utils.Initialized(itemID, memberID); err != nil {
		return err
	}

	fetchedItem, err := ws.itemRepository.Get(itemID)
	if err != nil {
		return err
	}
	if fetchedItem == nil {
		return errors.New("ItemDoesNotExist")
	}

	vote, err := NewVote(memberID, itemID)
	if err != nil {
		return err
	}

	exists, err := ws.voteRepository.DoesExist(vote)
	if err != nil {
		return err
	}

	switch exists {
	case true:
		return ws.voteRepository.Delete(vote)
	default:
		return errors.New("NeverVotedOnItem")
	}

}

func (ws *itemService) Items(start *ItemID, next uint32) ([]Item, error) {

	if err := utils.Initialized(start); err != nil {
		return nil, err
	}

	return ws.itemRepository.Query(start, next)

}

func (ws *itemService) GetByID(itemID ItemID) (*Item, error) {

	w, err := ws.itemRepository.Get(itemID)

	if err != nil {
		return nil, err
	}

	return w, nil

}

func (ws *itemService) VotesOfItem(itemID ItemID) (uint32, error) {

	if err := utils.Initialized(itemID); err != nil {
		return 0, err
	}

	return ws.itemRepository.VotesOf(itemID)

}

func (ws *itemService) Delete(item ItemID) error {

	if err := utils.Initialized(item); err != nil {
		return err
	}

	fetchedItem, err := ws.itemRepository.Get(item)
	if err != nil {
		return err
	}
	if fetchedItem == nil {
		return errors.New("ItemDoesNotExist")
	}

	return ws.itemRepository.Delete(item)

}

func (ws *itemService) Update(item Item) error {

	if err := utils.Initialized(item); err != nil {
		return err
	}

	fetchedItem, err := ws.itemRepository.Get(item.ID())
	if err != nil {
		return err
	}
	if fetchedItem == nil {
		return errors.New("ItemDoesNotExist")
	}

	return ws.itemRepository.Save(item)

}

func (ws *itemService) SetPhoto(item ItemID, rawPhoto []byte) error {

	if err := utils.Initialized(item); err != nil {
		return err
	}

	if err := ws.itemPhotoRepository.DeleteAllFor(item); err != nil {
		return err
	}

	if rawPhoto == nil {
		return nil
	}

	photoID, err := NewItemPhotoID(uuid.NewV4().String())
	if err != nil {
		return err
	}

	photo, err := NewItemPhoto(photoID, item, rawPhoto)
	if err != nil {
		return err
	}

	return ws.itemPhotoRepository.Save(photo)

}

func (ws *itemService) VotedOnItem(member MemberID, item ItemID) (bool, error) {

	vote, err := NewVote(member, item)
	if err != nil {
		return false, err
	}

	return ws.voteRepository.DoesExist(vote)

}

func (ws *itemService) GetPhoto(item ItemID) (*ItemPhoto, error) {
	return ws.itemPhotoRepository.Get(item)
}

func NewItemService(itemRepository IItemRepository, voteRepository IVoteRepository, itemPhotoRepo IItemPhotoRepository) IItemService {
	return &itemService{
		itemRepository:      itemRepository,
		voteRepository:      voteRepository,
		itemPhotoRepository: itemPhotoRepo,
	}
}
