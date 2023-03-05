package service

import (
	"errors"

	"github.com/St0rage/Simpan-Uang/model/domain"
	"github.com/St0rage/Simpan-Uang/model/web"
	"github.com/St0rage/Simpan-Uang/repository"
	"github.com/St0rage/Simpan-Uang/utils"
)

type PiggyBankService interface {
	CreatePiggyBank(userId string, newPiggyBank *web.PiggyBankCreateUpdateRequest) error
	GetAllPiggyBank(userId string) []web.PiggyBankReponse
	GetPiggyBankById(piggyBankId string) web.PiggyBankReponse
	UpdatePiggyBank(piggyBankId string, piggyBankUpdate *web.PiggyBankCreateUpdateRequest) error
	DeletePiggyBank(userId string, piggyBankId string) error
	GetMainPiggyBank(userId string) string
	GetPiggyBankUser(piggyBankId string) (string, error)
	GetAllPiggyBankTotal(userId string) float32
}

type piggyBankService struct {
	piggyBankRepo         repository.PiggyBankRepository
	piggyBankTransService PiggyBankTransactionService
}

func (piggyBankservice *piggyBankService) CreatePiggyBank(userId string, newPiggyBank *web.PiggyBankCreateUpdateRequest) error {
	var piggyBank domain.PiggyBank

	isMainPiggyBank := piggyBankservice.piggyBankRepo.CheckMainPiggyBank(userId)
	if isMainPiggyBank {

		isPiggyBankNameExist := piggyBankservice.piggyBankRepo.CheckPiggyBankName(newPiggyBank.PiggyBankName, userId)
		if isPiggyBankNameExist {
			return errors.New("nama tabungan sudah digunakan")
		}

		piggyBank.Type = false

	} else {
		piggyBank.Type = true
	}

	piggyBank.Id = utils.GenerateId()
	piggyBank.UserId = userId
	piggyBank.PiggyBankName = newPiggyBank.PiggyBankName

	piggyBankservice.piggyBankRepo.Save(&piggyBank)

	return nil
}

func (piggyBankService *piggyBankService) GetAllPiggyBank(userId string) []web.PiggyBankReponse {
	piggyBanks := piggyBankService.piggyBankRepo.FindAllByUserId(userId)

	piggyBankResponses := make([]web.PiggyBankReponse, len(piggyBanks))

	for i, piggyBank := range piggyBanks {
		piggyBankResponses[i].Id = piggyBank.Id
		piggyBankResponses[i].UserId = piggyBank.UserId
		piggyBankResponses[i].PiggyBankName = piggyBank.PiggyBankName
		piggyBankResponses[i].Type = piggyBank.Type
		piggyBankResponses[i].Total = piggyBankService.piggyBankTransService.GetTotalAmount(piggyBank.Id)
	}

	return piggyBankResponses
}

func (piggyBankService *piggyBankService) GetPiggyBankById(piggyBankId string) web.PiggyBankReponse {
	var piggyBankResponse web.PiggyBankReponse

	piggyBank := piggyBankService.piggyBankRepo.FindById(piggyBankId)

	piggyBankResponse.Id = piggyBank.Id
	piggyBankResponse.UserId = piggyBank.UserId
	piggyBankResponse.PiggyBankName = piggyBank.PiggyBankName
	piggyBankResponse.Type = piggyBank.Type
	piggyBankResponse.Total = piggyBankService.piggyBankTransService.GetTotalAmount(piggyBank.Id)

	return piggyBankResponse
}

func (piggyBankService *piggyBankService) UpdatePiggyBank(piggyBankId string, piggyBankUpdate *web.PiggyBankCreateUpdateRequest) error {
	piggyBank := piggyBankService.piggyBankRepo.FindById(piggyBankId)

	exist := piggyBankService.piggyBankRepo.CheckPiggyBankName(piggyBankUpdate.PiggyBankName, piggyBank.UserId)
	if exist {
		return errors.New("nama tabungan sudah digunakan")
	} else {
		piggyBank.PiggyBankName = piggyBankUpdate.PiggyBankName
	}

	piggyBankService.piggyBankRepo.Update(&piggyBank)

	return nil
}

func (piggyBankService *piggyBankService) DeletePiggyBank(userId string, piggyBankId string) error {
	total := piggyBankService.piggyBankTransService.GetTotalAmount(piggyBankId)
	mainPiggyBankId := piggyBankService.GetMainPiggyBank(userId)

	if mainPiggyBankId != piggyBankId {
		if total > 0 {
			var transferRequest web.TransferTransactionRequest
			transferRequest.Amount = total
			piggyBankService.piggyBankTransService.TransferTransaction(userId, mainPiggyBankId, &transferRequest)
		}

		piggyBankService.piggyBankRepo.Delete(piggyBankId)

		return nil
	} else {
		return errors.New("tabungan utama tidak dapat dihapus")
	}
}

func (piggyBankService *piggyBankService) GetMainPiggyBank(userId string) string {
	return piggyBankService.piggyBankRepo.FindMainPiggyBank(userId)
}

func (piggyBankService *piggyBankService) GetPiggyBankUser(piggyBankId string) (string, error) {
	return piggyBankService.piggyBankRepo.CheckPiggyBankUser(piggyBankId)
}

func (piggyBankService *piggyBankService) GetAllPiggyBankTotal(userId string) float32 {
	piggyBanks := piggyBankService.piggyBankRepo.FindAllByUserId(userId)
	piggyBanksTotal := make([]float32, len(piggyBanks))
	var total float32

	for i, piggyBank := range piggyBanks {
		piggyBanksTotal[i] = piggyBankService.piggyBankTransService.GetTotalAmount(piggyBank.Id)
	}

	for _, piggyBankTotal := range piggyBanksTotal {
		total += piggyBankTotal
	}

	return total
}

func NewPiggyBankService(piggyBankRepo repository.PiggyBankRepository, piggyBankTransService PiggyBankTransactionService) PiggyBankService {
	return &piggyBankService{
		piggyBankRepo:         piggyBankRepo,
		piggyBankTransService: piggyBankTransService,
	}
}
