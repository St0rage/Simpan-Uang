package manager

import "github.com/St0rage/Simpan-Uang/repository"

type RepositoryManager interface {
  // UserRepo
	UserRepo() repository.UserRepository
  // PiggyBankRepo
	PiggyBankRepo() repository.PiggyBankRepository
	// WishlistRepo
	WishlistRepo() repository.WishlistRepository
}

type repositoryManager struct {
	infra InfraManager
}

// UserRepo
func (repo *repositoryManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(repo.infra.DbConn())
}

// PiggyBankRepo
func (repo *repositoryManager) PiggyBankRepo() repository.PiggyBankRepository {
	return repository.NewPiggyBankRepository(repo.infra.DbConn())
}

// WhislistRepo()
func (repo *repositoryManager) WishlistRepo() repository.WishlistRepository {
	return repository.NewWishlistRepository(repo.infra.DbConn())
}

func NewRepositoryManager(manager InfraManager) RepositoryManager {
	return &repositoryManager{infra: manager}
}
