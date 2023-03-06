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

type PiggyBankTransactionService interface {
	DepositTransaction(piggyBankId string, depositRequest *web.DepositTransactionRequest)
	WithdrawTransaction(piggyBankId string, withdrawRequest *web.WithdrawTransactionRequest) (map[string]string, error)
	TransferTransaction(userId string, mainPiggyBankId string, transferRequest *web.TransferTransactionRequest)
	DeleteTransaction(piggyBankTransId string, piggyBankId string) error
	GetAllTransactions(piggyBankId string, page int) []domain.PiggyBankTransaction
	GetTotalAmount(piggyBankId string) float32
}

type piggyBankTransactionService struct {
	piggyBankTransRepo repository.PiggyBankTransactionRepository
}

func (piggyBankTransService *piggyBankTransactionService) DepositTransaction(piggyBankId string, depositRequest *web.DepositTransactionRequest) {
	var piggyBankTransaction domain.PiggyBankTransaction

	piggyBankTransaction.Id = utils.GenerateId()
	piggyBankTransaction.PiggyBankId = piggyBankId
	piggyBankTransaction.TransactionName = "Tambah Saldo"
	piggyBankTransaction.Amount = float32((depositRequest.Amount.(float64)))
	piggyBankTransaction.Status = true
	piggyBankTransaction.Date = time.Now().Unix()

	piggyBankTransService.piggyBankTransRepo.Save(&piggyBankTransaction)

}

func (piggyBankTransService *piggyBankTransactionService) WithdrawTransaction(piggyBankId string, withdrawRequest *web.WithdrawTransactionRequest) (map[string]string, error) {
	total := piggyBankTransService.GetTotalAmount(piggyBankId)

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

	var piggyBankTransaction domain.PiggyBankTransaction

	piggyBankTransaction.Id = utils.GenerateId()
	piggyBankTransaction.PiggyBankId = piggyBankId
	piggyBankTransaction.TransactionName = withdrawRequest.TransactionName
	piggyBankTransaction.Amount = 0 - amount
	piggyBankTransaction.Status = false
	piggyBankTransaction.Date = time.Now().Unix()

	piggyBankTransService.piggyBankTransRepo.Save(&piggyBankTransaction)

	return nil, nil
}

func (piggyBankTransService *piggyBankTransactionService) TransferTransaction(userId string, mainPiggyBankId string, transferRequest *web.TransferTransactionRequest) {
	var piggyBankTransaction domain.PiggyBankTransaction

	piggyBankTransaction.Id = utils.GenerateId()
	piggyBankTransaction.PiggyBankId = mainPiggyBankId
	piggyBankTransaction.TransactionName = "Pindahan"
	piggyBankTransaction.Amount = transferRequest.Amount
	piggyBankTransaction.Status = true
	piggyBankTransaction.Date = time.Now().Unix()

	piggyBankTransService.piggyBankTransRepo.Save(&piggyBankTransaction)
}

func (piggyBankTransService *piggyBankTransactionService) DeleteTransaction(piggyBankTransId string, piggyBankId string) error {
	lastTransactionId := piggyBankTransService.piggyBankTransRepo.FindLastTransaction(piggyBankId)

	if lastTransactionId != piggyBankTransId {
		return errors.New("hanya transaksi terakhir yang bisa dihapus")
	} else {
		piggyBankTransService.piggyBankTransRepo.Delete(piggyBankTransId)
		return nil
	}
}

func (piggyBankTransService *piggyBankTransactionService) GetAllTransactions(piggyBankId string, page int) []domain.PiggyBankTransaction {
	transactions := piggyBankTransService.piggyBankTransRepo.FindAll(piggyBankId, page)
	if transactions != nil {
		return transactions
	}
	return []domain.PiggyBankTransaction{}
}

func (piggyBankTransService *piggyBankTransactionService) GetTotalAmount(piggyBankId string) float32 {
	var total float32
	piggyBankTransactions := piggyBankTransService.piggyBankTransRepo.FindAmount(piggyBankId)

	for _, transcation := range piggyBankTransactions {
		total += transcation.Amount
	}

	return total
}

func NewPiggyBankTransactionService(piggyBankTransRepo repository.PiggyBankTransactionRepository) PiggyBankTransactionService {
	return &piggyBankTransactionService{
		piggyBankTransRepo: piggyBankTransRepo,
	}
}
