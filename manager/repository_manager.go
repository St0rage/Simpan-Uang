package manager

import "github.com/St0rage/Simpan-Uang/repository"

type RepositoryManager interface {
	UserRepo() repository.UserRepository
	// PiggyBankRepo
	// WhislistRepo
}

type repositoryManager struct {
	infra InfraManager
}

func (repo *repositoryManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(repo.infra.DbConn())
}

// PiggyBankRepo()
// WhislistRepo()

func NewRepositoryManager(manager InfraManager) RepositoryManager {
	return &repositoryManager{infra: manager}
}
