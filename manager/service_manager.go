package manager

import (
	"github.com/St0rage/Simpan-Uang/service"
	"github.com/St0rage/Simpan-Uang/utils/authenticator"
	"github.com/St0rage/Simpan-Uang/utils/mailer"
)

type ServiceManager interface {
  // UserService
	UserService() service.UserService
  // PiggyBankService
	PiggyBankService() service.PiggyBankService
	// WhislistService
	WishlistService() service.WishlistService
}

type serviceManager struct {
	repoManager RepositoryManager
	tokenServ   authenticator.AccessToken
	mailServ    mailer.MailService
}

// UserService
func (s *serviceManager) UserService() service.UserService {
	return service.NewUserService(s.repoManager.UserRepo(), s.tokenServ, s.mailServ)
}


// PiggyBankService
func (s *serviceManager) PiggyBankService() service.PiggyBankService {
	return service.NewPiggyBankService(s.repoManager.PiggyBankRepo())
}

// WhislistService
func (s *serviceManager) WishlistService() service.WishlistService {
	return service.NewWishlistService(s.repoManager.WishlistRepo())
}


func NewServiceManager(repoManager RepositoryManager, tokenServ authenticator.AccessToken, mailServ mailer.MailService) ServiceManager {
	return &serviceManager{
		repoManager: repoManager,
		tokenServ:   tokenServ,
		mailServ:    mailServ,
	}
}
