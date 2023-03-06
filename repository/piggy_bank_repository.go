package repository

import (
	"github.com/St0rage/Simpan-Uang/model/domain"
	"github.com/St0rage/Simpan-Uang/utils"
	"github.com/jmoiron/sqlx"
)

type PiggyBankRepository interface {
	Save(piggyBank *domain.PiggyBank)
	FindAllByUserId(userId string) []domain.PiggyBank
	FindById(piggyBankId string) domain.PiggyBank
	Update(piggyBank *domain.PiggyBank)
	Delete(piggyBankId string)
	FindMainPiggyBank(userId string) string
	CheckMainPiggyBank(userId string) bool
	CheckPiggyBankName(piggyBankName string, userId string) bool
	CheckPiggyBankUser(piggyBankId string) (string, error)
}

type piggyBankRepository struct {
	db *sqlx.DB
}

func (piggyBankRepo *piggyBankRepository) Save(piggyBank *domain.PiggyBank) {
	_, err := piggyBankRepo.db.NamedExec(utils.INSERT_PIGGY_BANK, &piggyBank)
	utils.PanicIfError(err)
}

func (piggyBankRepo *piggyBankRepository) FindAllByUserId(userId string) []domain.PiggyBank {
	var piggyBanks []domain.PiggyBank
	err := piggyBankRepo.db.Select(&piggyBanks, utils.SELECT_PIGGY_BANK, userId)
	utils.PanicIfError(err)

	return piggyBanks
}

func (piggyBankRepo *piggyBankRepository) FindById(piggyBankId string) domain.PiggyBank {
	var piggyBank domain.PiggyBank
	err := piggyBankRepo.db.Get(&piggyBank, utils.SELECT_PIGGY_BANK_ID, piggyBankId)
	utils.PanicIfError(err)

	return piggyBank
}

func (piggyBankRepo *piggyBankRepository) Update(piggyBank *domain.PiggyBank) {
	_, err := piggyBankRepo.db.NamedExec(utils.UPDATE_PIGGY_BANK, piggyBank)
	utils.PanicIfError(err)
}

func (piggyBankRepo *piggyBankRepository) Delete(piggyBankId string) {
	piggyBankRepo.db.MustExec(utils.DELETE_PIGGY_BANK, piggyBankId)
}

func (piggyBankRepo *piggyBankRepository) FindMainPiggyBank(userId string) string {
	var piggyBankId string
	err := piggyBankRepo.db.Get(&piggyBankId, utils.SELECT_MAIN_PIGGY_BANK, userId)
	utils.PanicIfError(err)

	return piggyBankId
}

func (piggyBankRepo *piggyBankRepository) CheckMainPiggyBank(userId string) bool {
	var exist int
	err := piggyBankRepo.db.Get(&exist, utils.CHECK_MAIN_PIGGY_BANK, userId)
	utils.PanicIfError(err)

	if exist == 1 {
		return true
	} else {
		return false
	}
}

func (piggyBankRepo *piggyBankRepository) CheckPiggyBankName(piggyBankName string, userId string) bool {
	var exist int
	err := piggyBankRepo.db.Get(&exist, utils.CHECK_PIGGY_BANK_NAME, piggyBankName, userId)
	utils.PanicIfError(err)

	if exist == 1 {
		return true
	} else {
		return false
	}
}

// Middleware Authorization
func (piggyBankRepo *piggyBankRepository) CheckPiggyBankUser(piggyBankId string) (string, error) {
	var userId string
	err := piggyBankRepo.db.Get(&userId, utils.SELECT_PIGGY_BANK_USER_ID, piggyBankId)
	if err != nil {
		return "", err
	}

	return userId, nil
}

func NewPiggyBankRepository(db *sqlx.DB) PiggyBankRepository {
	return &piggyBankRepository{
		db: db,
	}
}
