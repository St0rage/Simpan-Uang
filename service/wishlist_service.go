package service

import (
	"errors"

	"github.com/St0rage/Simpan-Uang/model/domain"
	"github.com/St0rage/Simpan-Uang/model/web"
	"github.com/St0rage/Simpan-Uang/repository"
	"github.com/St0rage/Simpan-Uang/utils"
)

type WishlistService interface {
	GetWishlist(userId string) ([]web.WishlistResponse, error)
	GetWishlistById(wishlistId string) web.WishlistResponse
	CreateNewWishlist(userId string, wishlist *web.WishlistRequest) error
	UpdateWishlist(wishlistId string, wishlistUpdate *web.WishlistUpdateRequest) error
	GetWishlistUser(wishlistId string) (string, error)
	GetWishlistTarget(wishlistId string) float32
}

type wishlistService struct {
	wishlistRepo         repository.WishlistRepository
	wishlistTransService WishlistTransactionService
}

func (w *wishlistService) GetWishlist(userId string) ([]web.WishlistResponse, error) {
	wishlist, err := w.wishlistRepo.GetAll(userId)
	wishlistResponse := make([]web.WishlistResponse, len(wishlist))
	if err != nil {
		return wishlistResponse, err
	}
	for i, value := range wishlist {
		wishlistResponse[i].Id = value.Id
		wishlistResponse[i].UserId = value.UserId
		wishlistResponse[i].WishlistName = value.WishlistName
		wishlistResponse[i].WishlistTarget = value.WishlistTarget
		wishlistResponse[i].Progress = value.Progress
		wishlistResponse[i].Total = w.wishlistTransService.GetWishlistTotal(value.Id)
	}
	return wishlistResponse, nil
}

func (w *wishlistService) CreateNewWishlist(userId string, wishlist *web.WishlistRequest) error {
	var newWishlist domain.Wishlist
	exist := w.wishlistRepo.CheckWishlistName(wishlist.WishlistName, userId)
	if exist {
		return errors.New("nama wishlist sudah digunakan")
	}
	newWishlist.Id = utils.GenerateId()
	newWishlist.UserId = userId
	newWishlist.WishlistName = wishlist.WishlistName
	newWishlist.WishlistTarget = wishlist.WishlistTarget
	err := w.wishlistRepo.CreateNewWishlist(newWishlist)
	if err != nil {
		return err
	}

	return nil
}

func (w *wishlistService) GetWishlistById(wishlistId string) web.WishlistResponse {
	var WishlistResponse web.WishlistResponse

	wishlist := w.wishlistRepo.FindById(wishlistId)

	WishlistResponse.Id = wishlist.Id
	WishlistResponse.UserId = wishlist.UserId
	WishlistResponse.WishlistName = wishlist.WishlistName
	WishlistResponse.WishlistTarget = wishlist.WishlistTarget
	WishlistResponse.Progress = wishlist.Progress
	WishlistResponse.Total = w.wishlistTransService.GetWishlistTotal(wishlist.Id)

	return WishlistResponse
}

func (w *wishlistService) UpdateWishlist(wishlistId string, wishlistUpdate *web.WishlistUpdateRequest) error {
	wishlist := w.wishlistRepo.FindById(wishlistId)
	total := w.wishlistTransService.GetWishlistTotal(wishlistId)

	exist := w.wishlistRepo.CheckWishlistName(wishlistUpdate.WishlistName, wishlist.UserId)
	if exist {
		return errors.New("nama wishlist sudah digunakan")
	} else if wishlistUpdate.WishlistTarget <= total {
		return errors.New("target tidak boleh kurang atau sama dengan total wishlist saat ini")
	}

	wishlist.WishlistName = wishlistUpdate.WishlistName
	wishlist.WishlistTarget = wishlistUpdate.WishlistTarget

	w.wishlistRepo.Update(&wishlist)

	return nil
}

func (w *wishlistService) GetWishlistUser(wishlistId string) (string, error) {
	return w.wishlistRepo.CheckWishlistUser(wishlistId)
}

func (w *wishlistService) GetWishlistTarget(wishlistId string) float32 {
	return w.wishlistRepo.GetTarget(wishlistId)
}

func NewWishlistService(wishlistRepo repository.WishlistRepository, wishlistTransService WishlistTransactionService) WishlistService {
	return &wishlistService{
		wishlistRepo:         wishlistRepo,
		wishlistTransService: wishlistTransService,
	}
}
