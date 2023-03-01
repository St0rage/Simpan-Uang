package manager

import "github.com/St0rage/Simpan-Uang/repository"

type RepositoryManager interface {
	UserRepo() repository.UserRepository
	PiggyBankRepo() repository.PiggyBankRepository
	// WhislistRepo
}

type repositoryManager struct {
	infra InfraManager
}

func (repo *repositoryManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(repo.infra.DbConn())
}

func (repo *repositoryManager) PiggyBankRepo() repository.PiggyBankRepository {
	return repository.NewPiggyBankRepository(repo.infra.DbConn())
}

// WhislistRepo()

func NewRepositoryManager(manager InfraManager) RepositoryManager {
	return &repositoryManager{infra: manager}
}
