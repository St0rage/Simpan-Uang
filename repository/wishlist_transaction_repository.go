package repository

import (
	"github.com/St0rage/Simpan-Uang/model/domain"
	"github.com/St0rage/Simpan-Uang/utils"
	"github.com/jmoiron/sqlx"
)

type WishlistTransactionRepository interface {
	Save(wishlistTransaction *domain.WishlistTransaction)
	GetAll(wishlistId string, page int) []domain.WishlistTransaction
	GetAmount(wishlistId string) []domain.WishlistTransaction
	Delete(wishlistTransId string)
	FindLastTransaction(wishlistId string) string
}

type wishlistTransactionRepository struct {
	db *sqlx.DB
}

func (wishlistTransRepo *wishlistTransactionRepository) Save(wishlistTransaction *domain.WishlistTransaction) {
	_, err := wishlistTransRepo.db.NamedExec(utils.INSERT_WISHLIST_TRANSACTION, &wishlistTransaction)
	if err != nil {
		panic(err)
	}
}

func (wishlistTransRepo *wishlistTransactionRepository) GetAll(wishlistId string, page int) []domain.WishlistTransaction {
	var wishlistTransaction []domain.WishlistTransaction

	limit := 10
	offset := limit * (page - 1)
	err := wishlistTransRepo.db.Select(&wishlistTransaction, utils.SELECT_WISHLIST_TRANSACTION, wishlistId, limit, offset)
	if err != nil {
		panic(err)
	}

	return wishlistTransaction
}

func (wishlistTransRepo *wishlistTransactionRepository) GetAmount(wishlistId string) []domain.WishlistTransaction {
	var wishlistTransaction []domain.WishlistTransaction
	err := wishlistTransRepo.db.Select(&wishlistTransaction, utils.SELECT_WISHLIST_AMOUNT, wishlistId)
	if err != nil {
		panic(err)
	}

	return wishlistTransaction
}

func (wishlistTransRepo *wishlistTransactionRepository) Delete(wishlistTransId string) {
	wishlistTransRepo.db.MustExec(utils.DELETE_WISHLIST_TRANSACTION, wishlistTransId)
}

func (wishlistTransRepo *wishlistTransactionRepository) FindLastTransaction(wishlistId string) string {
	var wishlistTransId string
	err := wishlistTransRepo.db.Get(&wishlistTransId, utils.SELECT_WISHLIST_LAST_TRANSACTION, wishlistId)
	if err != nil {
		panic(err)
	}

	return wishlistTransId
}

func NewWishlistTransactionRepository(db *sqlx.DB) WishlistTransactionRepository {
	return &wishlistTransactionRepository{
		db: db,
	}
}
