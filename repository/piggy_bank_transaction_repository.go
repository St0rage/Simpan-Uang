package repository

import (
	"github.com/St0rage/Simpan-Uang/model/domain"
	"github.com/St0rage/Simpan-Uang/utils"
	"github.com/jmoiron/sqlx"
)

type PiggyBankTransactionRepository interface {
	Save(piggyBankTransaction *domain.PiggyBankTransaction)
	FindAllTransactions(piggyBankId string, page int) []domain.PiggyBankTransaction
	FindAmount(piggyBankId string) []domain.PiggyBankTransaction
}

type piggyBankTransactionRepository struct {
	db *sqlx.DB
}

func (piggyBankTransRepo *piggyBankTransactionRepository) Save(piggyBankTransaction *domain.PiggyBankTransaction) {
	_, err := piggyBankTransRepo.db.NamedExec(utils.INSERT_PIGGY_BANK_TRANSACTION, &piggyBankTransaction)
	if err != nil {
		panic(err)
	}
}

func (piggyBankTransRepo *piggyBankTransactionRepository) FindAllTransactions(piggyBankId string, page int) []domain.PiggyBankTransaction {
	var piggyBankTransactions []domain.PiggyBankTransaction

	limit := 10
	offset := limit * (page - 1)
	err := piggyBankTransRepo.db.Select(&piggyBankTransactions, utils.SELECT_PIGGY_BANK_TRANSACTION, piggyBankId, limit, offset)
	if err != nil {
		panic(err)
	}

	return piggyBankTransactions
}

func (piggyBankTransRepo *piggyBankTransactionRepository) FindAmount(piggyBankId string) []domain.PiggyBankTransaction {
	var piggyBankTransactions []domain.PiggyBankTransaction
	err := piggyBankTransRepo.db.Select(&piggyBankTransactions, utils.SELECT_PIGGY_BANK_AMOUNT, piggyBankId)
	if err != nil {
		panic(err)
	}

	return piggyBankTransactions
}

func NewPiggyBankTransactionRepository(db *sqlx.DB) PiggyBankTransactionRepository {
	return &piggyBankTransactionRepository{
		db: db,
	}
}
