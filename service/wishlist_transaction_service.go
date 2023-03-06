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
	DepositWishlist(wishlistId string, wishlistTarget float32, depositRequest *web.DepositTransactionRequest) (map[string]string, error)
	WithdrawWishlist(wishlistId string, withdrawRequest *web.WithdrawTransactionRequest) (map[string]string, error)
	GetWishlistTransaction(wishlistId string, page int) []domain.WishlistTransaction
	GetWishlistTotal(wishlistId string) float32
	DeleteTransaction(wishlistTransId string, wishlistId string) error
}

type wishlistTransactionService struct {
	wishlistTransRepo repository.WishlistTransactionRepository
}

func (wishlistTransService *wishlistTransactionService) DepositWishlist(wishlistId string, wishlistTarget float32, depositRequest *web.DepositTransactionRequest) (map[string]string, error) {
	total := wishlistTransService.GetWishlistTotal(wishlistId)

	amount := float32(depositRequest.Amount.(float64))

	if total == wishlistTarget {
		return map[string]string{
			"amount": "wishlist sudah mencapai target",
		}, errors.New("error")
	} else if amount > wishlistTarget {
		return map[string]string{
			"amount": "jumlah deposit melebihi target",
		}, errors.New("error")
	} else if (amount + total) > wishlistTarget {
		return map[string]string{
			"amount": "jumlah deposit melebihi target",
		}, errors.New("error")
	}
	var wishlistTransaction domain.WishlistTransaction

	wishlistTransaction.Id = utils.GenerateId()
	wishlistTransaction.WishlistId = wishlistId
	wishlistTransaction.TransactionName = "Tambah Saldo"
	wishlistTransaction.Amount = amount
	wishlistTransaction.Status = true
	wishlistTransaction.Date = time.Now().Unix()

	wishlistTransService.wishlistTransRepo.Save(&wishlistTransaction)

	return nil, nil
}

func (wishlistTransService *wishlistTransactionService) WithdrawWishlist(wishlistId string, withdrawRequest *web.WithdrawTransactionRequest) (map[string]string, error) {
	total := wishlistTransService.GetWishlistTotal(wishlistId)

	amount := float32(withdrawRequest.Amount.(float64))

	if total == 0 {
		return map[string]string{
			"amount": "withdraw gagal saldo tidak mencukupi",
		}, errors.New("error")
	} else if amount > total {
		return map[string]string{
			"amount": "Gagal, penarikan tidak boleh lebih dari Rp " + strconv.Itoa(int(total)),
		}, errors.New("error")
	}

	var wishlistTransaction domain.WishlistTransaction

	wishlistTransaction.Id = utils.GenerateId()
	wishlistTransaction.WishlistId = wishlistId
	wishlistTransaction.TransactionName = withdrawRequest.TransactionName
	wishlistTransaction.Amount = 0 - amount
	wishlistTransaction.Status = false
	wishlistTransaction.Date = time.Now().Unix()

	wishlistTransService.wishlistTransRepo.Save(&wishlistTransaction)

	return nil, nil
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
	transactions := wishlistTransService.wishlistTransRepo.GetAll(wishlistId, page)
	if transactions != nil {
		return transactions
	}

	return []domain.WishlistTransaction{}
}

func (wishlistTransService *wishlistTransactionService) DeleteTransaction(wishlistTransId string, wishlistId string) error {
	lastTransactionId := wishlistTransService.wishlistTransRepo.FindLastTransaction(wishlistId)

	if lastTransactionId != wishlistTransId {
		return errors.New("hanya transaksi terakhir yang bisa dihapus")
	} else {
		wishlistTransService.wishlistTransRepo.Delete(wishlistTransId)
		return nil
	}
}

func NewWishlistTransactionService(wishlistTransRepo repository.WishlistTransactionRepository) WishlistTransactionService {
	return &wishlistTransactionService{
		wishlistTransRepo: wishlistTransRepo,
	}
}
