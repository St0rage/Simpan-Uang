package repository

import (
	"github.com/St0rage/Simpan-Uang/model/domain"
	"github.com/St0rage/Simpan-Uang/utils"
	"github.com/jmoiron/sqlx"
)

type PiggyBankTransactionRepository interface {
	Save(piggyBankTransaction *domain.PiggyBankTransaction)
	FindAll(piggyBankId string, page int) []domain.PiggyBankTransaction
	Delete(piggyBankTransId string)
	FindAmount(piggyBankId string) []domain.PiggyBankTransaction
	FindLastTransaction(piggyBankId string) string
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


func (piggyBankTransRepo *piggyBankTransactionRepository) Delete(piggyBankTransId string) {
	piggyBankTransRepo.db.MustExec(utils.DELETE_PIGGY_BANK_TRANSACTION, piggyBankTransId)
}

func (piggyBankTransRepo *piggyBankTransactionRepository) FindAmount(piggyBankId string) []domain.PiggyBankTransaction {
	var piggyBankTransactions []domain.PiggyBankTransaction
	err := piggyBankTransRepo.db.Select(&piggyBankTransactions, utils.SELECT_PIGGY_BANK_AMOUNT, piggyBankId)
	if err != nil {
		panic(err)
	}

	return piggyBankTransactions
}


func (piggyBankTransRepo *piggyBankTransactionRepository) FindLastTransaction(piggyBankId string) string {
	var piggyBankTransId string
	err := piggyBankTransRepo.db.Get(&piggyBankTransId, utils.SELECT_LAST_TRANSACTION, piggyBankId)
	if err != nil {
		panic(err)
	}

	return piggyBankTransId
}

func NewPiggyBankTransactionRepository(db *sqlx.DB) PiggyBankTransactionRepository {
	return &piggyBankTransactionRepository{
		db: db,
	}
}
