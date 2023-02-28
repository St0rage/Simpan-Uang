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
	// GetWishlistById(Id string) (web.WishlistIdRequest, error)
	CreateNewWishlist(userId string, wishlist *web.WishlistRequest) error
}

type wishlistService struct {
	wishlistRepo repository.WishlistRepository
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
		wishlistResponse[i].Total = 0
	}
	return wishlistResponse, nil
}

// func (w *wishlistService) GetWishlistById(Id string) (web.WishlistIdRequest, error) {
// 	wishlist, err := w.wishlistRepo.GetById(Id)
// 	if err != nil {
// 		return wishlist, err
// 	}
// 	return wishlist, nil
// }

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

func NewWishlistService(wishlistRepo repository.WishlistRepository) WishlistService {
	return &wishlistService{
		wishlistRepo: wishlistRepo,
	}
}
