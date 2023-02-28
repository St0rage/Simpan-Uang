package repository

import (
	"github.com/St0rage/Simpan-Uang/model/domain"
	"github.com/St0rage/Simpan-Uang/utils"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Save(user *domain.User)
	FindById(userId string) domain.User
	FindByEmail(userEmail string) (domain.User, error)
	Update(user *domain.User)
	UpdatePassword(user *domain.User)
	CheckAdmin() bool
	CheckEmail(email string) bool
	IsAdmin(userId string) bool
}

type userRepository struct {
	db *sqlx.DB
}

func (userRepo *userRepository) Save(user *domain.User) {
	userRepo.db.NamedExec(utils.INSERT_USER, &user)
}

func (userRepo *userRepository) FindById(userId string) domain.User {
	var user domain.User
	userRepo.db.Get(&user, utils.SELECT_USER_ID, userId)

	return user
}

func (userRepo *userRepository) FindByEmail(userEmail string) (domain.User, error) {
	var user domain.User
	err := userRepo.db.Get(&user, utils.SELECT_USER_EMAIL, userEmail)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (userRepo *userRepository) Update(user *domain.User) {
	_, err := userRepo.db.NamedExec(utils.UPDATE_USER, &user)
	if err != nil {
		panic(err)
	}
}

func (userRepo *userRepository) UpdatePassword(user *domain.User) {
	_, err := userRepo.db.NamedExec(utils.UPDATE_USER_PASSWORD, &user)
	if err != nil {
		panic(err)
	}
}

func (userRepo *userRepository) CheckAdmin() bool {
	var is_admin int
	userRepo.db.Get(&is_admin, utils.CHECK_ADMIN)

	if is_admin == 0 {
		return false
	} else {
		return true
	}
}

func (userRepo *userRepository) IsAdmin(userId string) bool {
	var is_admin bool
	userRepo.db.Get(&is_admin, utils.IS_ADMIN, userId)

	if is_admin {
		return true
	} else {
		return false
	}
}

func (userRepo *userRepository) CheckEmail(email string) bool {
	var exist int
	userRepo.db.Get(&exist, utils.CHECK_EMAIL, email)

	if exist == 1 {
		return true
	} else {
		return false
	}
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
