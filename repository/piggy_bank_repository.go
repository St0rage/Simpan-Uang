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
	CheckMainPiggyBank(userId string) bool
	CheckPiggyBankName(piggyBankName string, userId string) bool
	CheckPiggyBankUser(piggyBankId string) (string, error)
}

type piggyBankRepository struct {
	db *sqlx.DB
}

func (piggyBankRepo *piggyBankRepository) Save(piggyBank *domain.PiggyBank) {
	_, err := piggyBankRepo.db.NamedExec(utils.INSERT_PIGGY_BANK, &piggyBank)
	if err != nil {
		panic(err)
	}
}

func (piggyBankRepo *piggyBankRepository) FindAllByUserId(userId string) []domain.PiggyBank {
	var piggyBanks []domain.PiggyBank
	err := piggyBankRepo.db.Select(&piggyBanks, utils.SELECT_PIGGY_BANK, userId)
	if err != nil {
		panic(err)
	}

	return piggyBanks
}

func (piggyBankRepo *piggyBankRepository) FindById(piggyBankId string) domain.PiggyBank {
	var piggyBank domain.PiggyBank
	err := piggyBankRepo.db.Get(&piggyBank, utils.SELECT_PIGGY_BANK_ID, piggyBankId)
	if err != nil {
		panic(err)
	}

	return piggyBank
}

func (piggyBankRepository *piggyBankRepository) Update(piggyBank *domain.PiggyBank) {
	_, err := piggyBankRepository.db.NamedExec(utils.UPDATE_PIGGY_BANK, piggyBank)
	if err != nil {
		panic(err)
	}
}

func (piggyBankRepo *piggyBankRepository) CheckMainPiggyBank(userId string) bool {
	var exist int
	err := piggyBankRepo.db.Get(&exist, utils.CHECK_MAIN_PIGGY_BANK, userId)
	if err != nil {
		panic(err)
	}

	if exist == 1 {
		return true
	} else {
		return false
	}
}

func (piggyBankRepo *piggyBankRepository) CheckPiggyBankName(piggyBankName string, userId string) bool {
	var exist int
	err := piggyBankRepo.db.Get(&exist, utils.CHECK_PIGGY_BANK_NAME, piggyBankName, userId)
	if err != nil {
		panic(err)
	}

	if exist == 1 {
		return true
	} else {
		return false
	}
}

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