package manager

import "github.com/St0rage/Simpan-Uang/repository"

type RepositoryManager interface {
	// UserRepo
	UserRepo() repository.UserRepository
	// PiggyBankRepo
	PiggyBankRepo() repository.PiggyBankRepository
	PiggyBankTransRepo() repository.PiggyBankTransactionRepository
	// WishlistRepo
	WishlistRepo() repository.WishlistRepository
	WishlistTransRepo() repository.WishlistTransactionRepository

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

func (repo *repositoryManager) PiggyBankTransRepo() repository.PiggyBankTransactionRepository {
	return repository.NewPiggyBankTransactionRepository(repo.infra.DbConn())
}

// WhislistRepo()
func (repo *repositoryManager) WishlistRepo() repository.WishlistRepository {
	return repository.NewWishlistRepository(repo.infra.DbConn())
}

func (repo *repositoryManager) WishlistTransRepo() repository.WishlistTransactionRepository {
	return repository.NewWishlistTransactionRepository(repo.infra.DbConn())
}

func NewRepositoryManager(manager InfraManager) RepositoryManager {
	return &repositoryManager{infra: manager}
}
