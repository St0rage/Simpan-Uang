package repository

import (
	"github.com/St0rage/Simpan-Uang/model/domain"
	"github.com/St0rage/Simpan-Uang/utils"
	"github.com/jmoiron/sqlx"
)

type WishlistRepository interface {
	GetAll(userId string) ([]domain.Wishlist, error)
	GetById(Id string) ([]domain.Wishlist, error)
	CreateNewWishlist(wishlistName domain.Wishlist) error
	CheckWishlistName(wishlistName string, userId string) bool

}

type wishlistRepository struct {
	db *sqlx.DB
}

func (w *wishlistRepository) GetAll(userId string) ([]domain.Wishlist, error) {
	var wishlists []domain.Wishlist
	err := w.db.Select(&wishlists, utils.SELECT_WISHLIST, userId)
	if err != nil {
		return nil, err
	}
	return wishlists, nil
}

func (w *wishlistRepository) CheckWishlistName(wishlistName string, userId string) bool {
	var exist int
	err := w.db.Get(&exist, utils.CHECK_WISHLIST_NAME, wishlistName, userId)
	if err != nil {
		panic(err)
	}

	if exist == 1 {
		return true
	} else {
		return false
	}
}

func (w *wishlistRepository) GetById(Id string) ([]domain.Wishlist, error) {
	var wishlists []domain.Wishlist
	err := w.db.Select(utils.SELECT_WISHLIST_ID, Id)
	if err != nil {
		return nil, err
	}
	return wishlists, nil
}

func (w *wishlistRepository) CreateNewWishlist(wishlistName domain.Wishlist) error {
	_, err := w.db.NamedExec(utils.INSERT_WISHLIST, wishlistName)
	if err != nil {
		return err
	}
	return nil
}

func NewWishlistRepository(db *sqlx.DB) WishlistRepository {
	return &wishlistRepository{
		db: db,
	}
}
