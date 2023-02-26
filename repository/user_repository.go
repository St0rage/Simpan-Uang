package repository

import (
	"github.com/St0rage/Simpan-Uang/model/domain"
	"github.com/St0rage/Simpan-Uang/utils"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Save(user *domain.User) error
	FindById(userId string) (domain.User, error)
	FindByEmail(userEmail string) (domain.User, error)
	Update(user *domain.User) error
	UpdatePassword(user *domain.User) error
}

type userRepository struct {
	db *sqlx.DB
}

func (userRepo *userRepository) Save(user *domain.User) error {
	_, err := userRepo.db.NamedExec(utils.INSERT_USER, &user)
	if err != nil {
		return err
	}

	return nil
}

func (userRepo *userRepository) FindById(userId string) (domain.User, error) {
	var user domain.User
	err := userRepo.db.Get(&user, utils.SELECT_USER_ID, userId)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (userRepo *userRepository) FindByEmail(userEmail string) (domain.User, error) {
	var user domain.User
	err := userRepo.db.Get(&user, utils.SELECT_USER_EMAIL, userEmail)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (userRepo *userRepository) Update(user *domain.User) error {
	_, err := userRepo.db.NamedExec(utils.UPDATE_USER, &user)
	if err != nil {
		return err
	}

	return nil
}

func (userRepo *userRepository) UpdatePassword(user *domain.User) error {
	_, err := userRepo.db.NamedExec(utils.UPDATE_USER_PASSWORD, &user)
	if err != nil {
		return err
	}

	return nil
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
