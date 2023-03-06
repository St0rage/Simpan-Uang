package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/St0rage/Simpan-Uang/model/domain"
	"github.com/St0rage/Simpan-Uang/model/web"
	"github.com/St0rage/Simpan-Uang/repository"
	"github.com/St0rage/Simpan-Uang/utils"
)

type WishlistService interface {
	GetWishlist(userId string) []web.WishlistResponse
	GetWishlistById(wishlistId string) web.WishlistResponse
	CreateNewWishlist(userId string, wishlist *web.WishlistCreateUpdateRequest) (map[string]string, error)
	UpdateWishlist(wishlistId string, wishlistUpdate *web.WishlistCreateUpdateRequest) (map[string]string, error)
	GetWishlistUser(wishlistId string) (string, error)
	GetWishlistTarget(wishlistId string) float32
	DeleteWishlist(userId string, wishlistId string)
	GetAllWishlistTotal(userId string) float32
}

type wishlistService struct {
	wishlistRepo          repository.WishlistRepository
	wishlistTransService  WishlistTransactionService
	piggyBankService      PiggyBankService
	piggyBankTransService PiggyBankTransactionService
}

func (w *wishlistService) GetWishlist(userId string) []web.WishlistResponse {
	wishlist := w.wishlistRepo.GetAll(userId)
	wishlistResponse := make([]web.WishlistResponse, len(wishlist))

	for i, value := range wishlist {
		wishlistResponse[i].Id = value.Id
		wishlistResponse[i].UserId = value.UserId
		wishlistResponse[i].WishlistName = value.WishlistName
		wishlistResponse[i].WishlistTarget = value.WishlistTarget
		wishlistResponse[i].Total = w.wishlistTransService.GetWishlistTotal(value.Id)

		// Progress
		rawProgress := (wishlistResponse[i].Total / value.WishlistTarget) * 100
		progress, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", rawProgress), 32)

		wishlistResponse[i].Progress = float32(progress)
	}

	return wishlistResponse
}

func (w *wishlistService) CreateNewWishlist(userId string, wishlistRequest *web.WishlistCreateUpdateRequest) (map[string]string, error) {
	var newWishlist domain.Wishlist
	exist := w.wishlistRepo.CheckWishlistName(wishlistRequest.WishlistName, userId)
	if exist {
		return map[string]string{
			"wishlist_name": "Nama wishlist sudah digunakan",
		}, errors.New("error")
	}

	wishlistTarget := float32(wishlistRequest.WishlistTarget.(float64))

	newWishlist.Id = utils.GenerateId()
	newWishlist.UserId = userId
	newWishlist.WishlistName = wishlistRequest.WishlistName
	newWishlist.WishlistTarget = wishlistTarget

	w.wishlistRepo.CreateNewWishlist(newWishlist)

	return nil, nil
}

func (w *wishlistService) GetWishlistById(wishlistId string) web.WishlistResponse {
	var wishlistResponse web.WishlistResponse

	wishlist := w.wishlistRepo.FindById(wishlistId)

	wishlistResponse.Id = wishlist.Id
	wishlistResponse.UserId = wishlist.UserId
	wishlistResponse.WishlistName = wishlist.WishlistName
	wishlistResponse.WishlistTarget = wishlist.WishlistTarget
	wishlistResponse.Total = w.wishlistTransService.GetWishlistTotal(wishlist.Id)

	// Progress
	rawProgress := (wishlistResponse.Total / wishlist.WishlistTarget) * 100
	progress, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", rawProgress), 32)

	wishlistResponse.Progress = float32(progress)

	return wishlistResponse
}

func (w *wishlistService) UpdateWishlist(wishlistId string, wishlistUpdate *web.WishlistCreateUpdateRequest) (map[string]string, error) {
	wishlist := w.wishlistRepo.FindById(wishlistId)
	total := w.wishlistTransService.GetWishlistTotal(wishlistId)

	if wishlistUpdate.WishlistName != wishlist.WishlistName {
		exist := w.wishlistRepo.CheckWishlistName(wishlistUpdate.WishlistName, wishlist.UserId)
		if exist {
			return map[string]string{
				"wishlist_name": "nama wishlist sudah digunakan",
			}, errors.New("error")
		}
	}

	wishlistTarget := float32(wishlistUpdate.WishlistTarget.(float64))

	if wishlistTarget <= total {
		return map[string]string{
			"wishlist_target": "target tidak boleh kurang atau sama dengan total wishlist saat ini",
		}, errors.New("error")
	}

	wishlist.WishlistName = wishlistUpdate.WishlistName
	wishlist.WishlistTarget = wishlistTarget

	w.wishlistRepo.Update(&wishlist)

	return nil, nil
}

func (w *wishlistService) GetWishlistUser(wishlistId string) (string, error) {
	return w.wishlistRepo.CheckWishlistUser(wishlistId)
}

func (w *wishlistService) GetWishlistTarget(wishlistId string) float32 {
	return w.wishlistRepo.GetTarget(wishlistId)
}

func (w *wishlistService) DeleteWishlist(userId string, wishlistId string) {
	total := w.wishlistTransService.GetWishlistTotal(wishlistId)
	mainPiggyBankId := w.piggyBankService.GetMainPiggyBank(userId)

	if total > 0 {
		var transferRequest web.TransferTransactionRequest
		transferRequest.Amount = total
		w.piggyBankTransService.TransferTransaction(userId, mainPiggyBankId, &transferRequest)
	}

	w.wishlistRepo.Delete(wishlistId)
}

func (w *wishlistService) GetAllWishlistTotal(userId string) float32 {
	wishlists := w.wishlistRepo.GetAll(userId)
	wishlistsTotal := make([]float32, len(wishlists))
	var total float32

	for i, wishlist := range wishlists {
		wishlistsTotal[i] = w.wishlistTransService.GetWishlistTotal(wishlist.Id)
	}

	for _, wishlistTotal := range wishlistsTotal {
		total += wishlistTotal
	}

	return total
}

func NewWishlistService(wishlistRepo repository.WishlistRepository, wishlistTransService WishlistTransactionService, piggyBankService PiggyBankService, piggyBankTransService PiggyBankTransactionService) WishlistService {
	return &wishlistService{
		wishlistRepo:          wishlistRepo,
		wishlistTransService:  wishlistTransService,
		piggyBankService:      piggyBankService,
		piggyBankTransService: piggyBankTransService,
	}
}
