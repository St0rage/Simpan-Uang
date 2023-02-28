package manager

import "github.com/St0rage/Simpan-Uang/repository"

type RepositoryManager interface {
	UserRepo() repository.UserRepository
	// PiggyBankRepo
	WishlistRepo() repository.WishlistRepository
}

type repositoryManager struct {
	infra InfraManager
}

func (repo *repositoryManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(repo.infra.DbConn())
}

// PiggyBankRepo()
func (repo *repositoryManager) WishlistRepo() repository.WishlistRepository {
	return repository.NewWishlistRepository(repo.infra.DbConn())
}

func NewRepositoryManager(manager InfraManager) RepositoryManager {
	return &repositoryManager{infra: manager}
}
