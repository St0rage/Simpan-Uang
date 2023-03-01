package delivery

import (
	"fmt"

	"github.com/St0rage/Simpan-Uang/config"
	"github.com/St0rage/Simpan-Uang/delivery/controller"
	"github.com/St0rage/Simpan-Uang/delivery/middleware"
	"github.com/St0rage/Simpan-Uang/manager"
	"github.com/St0rage/Simpan-Uang/utils"
	"github.com/St0rage/Simpan-Uang/utils/authenticator"
	"github.com/St0rage/Simpan-Uang/utils/mailer"
	"github.com/gin-gonic/gin"
)

type Server struct {
	serviceManager manager.ServiceManager
	engine         *gin.Engine
	host           string
	tokenServ      authenticator.AccessToken
}

func (server *Server) Run() {
	server.initController()
	err := server.engine.Run(server.host)
	utils.PanicIfError(err)
}

func (server *Server) initController() {
	// Middleware
	authMdw := middleware.NewAuthMiddleware(server.tokenServ, server.serviceManager.PiggyBankService())

	// Controller
	controller.NewUserController(server.engine, server.serviceManager.UserService(), authMdw)
	controller.NewPiggyBankController(server.engine, server.serviceManager.PiggyBankService(), authMdw)
	controller.NewWishlistController(server.engine, server.serviceManager.WishlistService(), authMdw)
}

func NewServer() *Server {
	config := config.NewConfig()
	r := gin.Default()
	tokenServ := authenticator.NewAccessToken(config.TokenConfig)
	mailServ := mailer.NewMailService(config.MailConfig)
	infra := manager.NewInfraManager(config)
	repo := manager.NewRepositoryManager(infra)
	service := manager.NewServiceManager(repo, tokenServ, mailServ)

	if config.ApiHost == "" || config.ApiPort == "" {
		panic("No Host or Port define")
	}

	host := fmt.Sprintf("%s:%s", config.ApiHost, config.ApiPort)
	return &Server{
		serviceManager: service,
		engine:         r,
		host:           host,
		tokenServ:      tokenServ,
	}
}
