package service

import (
	"github.com/St0rage/Simpan-Uang/model/domain"
	"github.com/St0rage/Simpan-Uang/model/web"
	"github.com/St0rage/Simpan-Uang/repository"
	"github.com/St0rage/Simpan-Uang/utils"
)

type WishlistService interface {
	GetWishlist(userId string) ([]domain.Wishlist, error)
	CreateNewWishlist(userId string, wishlist *web.WishlistRequest) error
}

type wishlistService struct {
	wishlistRepo repository.WishlistRepository
}

func (wishlistService *wishlistService) GetWishlist(userId string) ([]domain.Wishlist, error) {
	// var wishlistResponse []web.WishlistResponse
	wishlist, err := wishlistService.wishlistRepo.GetAll(userId)
	if err != nil {
		return wishlist, err
	}
	// for i, value := range wishlist {
	// 	wishlistResponse[i].Id = value.Id
	// 	wishlistResponse[i].UserId = value.UserId
	// 	wishlistResponse[i].WishlistName = value.WishlistName
	// 	wishlistResponse[i].WishlistTarget = value.WishlistTarget
	// 	wishlistResponse[i].Progress = value.Progress
	// 	wishlistResponse[i].Total = 0
	// }
	return wishlist, nil
}

func (wishlistService *wishlistService) CreateNewWishlist(userId string, wishlist *web.WishlistRequest) error {
	var newWishlist domain.Wishlist
	newWishlist.Id = utils.GenerateId()
	newWishlist.UserId = userId
	newWishlist.WishlistName = wishlist.WishlistName
	newWishlist.WishlistTarget = wishlist.WishlistTarget
	err := wishlistService.wishlistRepo.CreateNewWishlist(newWishlist)
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
