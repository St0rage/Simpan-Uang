package repository

import (
	"github.com/St0rage/Simpan-Uang/model/domain"
	"github.com/St0rage/Simpan-Uang/utils"
	"github.com/jmoiron/sqlx"
)

type WishlistRepository interface {
	GetAll(userId string) ([]domain.Wishlist, error)
	FindById(wishlistId string) domain.Wishlist
	Update(wishlist *domain.Wishlist)
	CreateNewWishlist(wishlistName domain.Wishlist) error
	CheckWishlistName(wishlistName string, userId string) bool
	CheckWishlistUser(wishlistId string) (string, error)
	GetTarget(wishlistId string) float32
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

func (w *wishlistRepository) FindById(wishlistId string) domain.Wishlist {
	var wishlist domain.Wishlist
	err := w.db.Get(&wishlist, utils.SELECT_WISHLIST_ID, wishlistId)
	if err != nil {
		panic(err)
	}

	return wishlist
}

func (w *wishlistRepository) Update(wishlist *domain.Wishlist) {
	_, err := w.db.NamedExec(utils.UPDATE_WISHLIST, wishlist)
	if err != nil {
		panic(err)
	}
}

func (w *wishlistRepository) CreateNewWishlist(wishlistName domain.Wishlist) error {
	_, err := w.db.NamedExec(utils.INSERT_WISHLIST, wishlistName)
	if err != nil {
		return err
	}
	return nil
}

func (w *wishlistRepository) CheckWishlistUser(wishlistId string) (string, error) {
	var userId string
	err := w.db.Get(&userId, utils.SELECT_WISHLIST_USER_ID, wishlistId)
	if err != nil {
		return "", err
	}

	return userId, nil
}

func (w *wishlistRepository) GetTarget(wishlistId string) float32 {
	var wishlistTarget float32
	err := w.db.Get(&wishlistTarget, utils.SELECT_WISHLIST_TARGET, wishlistId)
	if err != nil {
		panic(err)
	}
	return wishlistTarget
}

func NewWishlistRepository(db *sqlx.DB) WishlistRepository {
	return &wishlistRepository{
		db: db,
	}
}
