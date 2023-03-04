package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/St0rage/Simpan-Uang/model/domain"
	"github.com/St0rage/Simpan-Uang/model/web"
	"github.com/St0rage/Simpan-Uang/repository"
	"github.com/St0rage/Simpan-Uang/utils"
)

type WishlistTransactionService interface {
	DepositWishlist(wishlistId string, wishlistTarget float32, depositRequest *web.DepositTransactionRequest) error
	WithdrawWishlist(wishlistId string, withdrawRequest *web.WithdrawTransactionRequest) error
	GetWishlistTransaction(wishlistId string, page int) []domain.WishlistTransaction
	GetWishlistTotal(wishlistId string) float32
}

type wishlistTransactionService struct {
	wishlistTransRepo repository.WishlistTransactionRepository
}

func (wishlistTransService *wishlistTransactionService) DepositWishlist(wishlistId string, wishlistTarget float32, depositRequest *web.DepositTransactionRequest) error {
	total := wishlistTransService.GetWishlistTotal(wishlistId)

	if depositRequest.Amount < 500 {
		return errors.New("minimal deposit Rp 500")
	} else if total == wishlistTarget {
		return errors.New("wishlist sudah mencapai target")
	} else if depositRequest.Amount > wishlistTarget {
		return errors.New("jumlah deposit melebihi target")
	} else if (depositRequest.Amount + total) > wishlistTarget {
		return errors.New("jumlah deposit melebihi target")
	}
	var wishlistTransaction domain.WishlistTransaction

	wishlistTransaction.Id = utils.GenerateId()
	wishlistTransaction.WishlistId = wishlistId
	wishlistTransaction.TransactionName = "Tambah Saldo"
	wishlistTransaction.Amount = depositRequest.Amount
	wishlistTransaction.Status = true
	wishlistTransaction.Date = time.Now().Unix()

	wishlistTransService.wishlistTransRepo.Save(&wishlistTransaction)

	return nil
}

func (wishlistTransService *wishlistTransactionService) WithdrawWishlist(wishlistId string, withdrawRequest *web.WithdrawTransactionRequest) error {
	total := wishlistTransService.GetWishlistTotal(wishlistId)

	if total == 0 {
		return errors.New("withdraw gagal saldo tidak mencukupi")
	} else if withdrawRequest.Amount < 500 {
		return errors.New("minimal withdraw Rp 500")
	} else if withdrawRequest.Amount > total {
		return errors.New("Gagal, penarikan tidak boleh lebih dari Rp " + strconv.Itoa(int(total)))
	}

	var wishlistTransaction domain.WishlistTransaction

	wishlistTransaction.Id = utils.GenerateId()
	wishlistTransaction.WishlistId = wishlistId
	wishlistTransaction.TransactionName = withdrawRequest.TransactionName
	wishlistTransaction.Amount = 0 - withdrawRequest.Amount
	wishlistTransaction.Status = false
	wishlistTransaction.Date = time.Now().Unix()

	wishlistTransService.wishlistTransRepo.Save(&wishlistTransaction)

	return nil
}

func (wishlistTransService *wishlistTransactionService) GetWishlistTotal(wishlistId string) float32 {
	var total float32
	wishlistTransaction := wishlistTransService.wishlistTransRepo.GetAmount(wishlistId)

	for _, transcation := range wishlistTransaction {
		total += transcation.Amount
	}

	return total
}

func (wishlistTransService *wishlistTransactionService) GetWishlistTransaction(wishlistId string, page int) []domain.WishlistTransaction {
	return wishlistTransService.wishlistTransRepo.GetAll(wishlistId, page)
}

func NewWishlistTransactionService(wishlistTransRepo repository.WishlistTransactionRepository) WishlistTransactionService {
	return &wishlistTransactionService{
		wishlistTransRepo: wishlistTransRepo,
	}
}
