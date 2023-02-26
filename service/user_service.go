package service

import (
	"math/rand"
	"strconv"

	"github.com/St0rage/Simpan-Uang/model/domain"
	"github.com/St0rage/Simpan-Uang/model/web"
	"github.com/St0rage/Simpan-Uang/repository"
	"github.com/St0rage/Simpan-Uang/utils"
	"github.com/St0rage/Simpan-Uang/utils/authenticator"
	"github.com/St0rage/Simpan-Uang/utils/mailer"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUser(userId string) (web.UserResponse, error)
	Register(newUser *domain.User) error
	Login(loginRequest *web.UserLoginRequest) (string, error)
	ForgotPassword(resetRequest *web.UserResetRequest) error
}

type userService struct {
	userRepo  repository.UserRepository
	tokenServ authenticator.AccessToken
	mailServ  mailer.MailService
}

func (userService *userService) GetUser(userId string) (web.UserResponse, error) {
	var userResponse web.UserResponse
	user, err := userService.userRepo.FindById(userId)
	if err != nil {
		return userResponse, err
	}

	userResponse.Id = user.Id
	userResponse.Name = user.Name
	userResponse.Email = user.Email

	return userResponse, nil
}

func (userService *userService) Register(newUser *domain.User) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	if err != nil {
		return err
	}

	newUser.Id = utils.GenerateId()
	newUser.Password = string(bytes)

	err = userService.userRepo.Save(newUser)
	if err != nil {
		return err
	}

	return nil
}

func (userService *userService) Login(loginRequest *web.UserLoginRequest) (string, error) {
	user, err := userService.userRepo.FindByEmail(loginRequest.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return "", err
	}

	token, err := userService.tokenServ.CreateAccessToken(&user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (userService *userService) ForgotPassword(resetRequest *web.UserResetRequest) error {
	user, err := userService.userRepo.FindByEmail(resetRequest.Email)
	if err != nil {
		return err
	}

	generatePassword := strconv.Itoa(rand.Intn(100000))
	password, _ := bcrypt.GenerateFromPassword([]byte(generatePassword), 14)

	user.Password = string(password)

	_ = userService.userRepo.UpdatePassword(&user)

	err = userService.mailServ.ResetPasswordMail(user.Email, generatePassword)
	if err != nil {
		panic(err)
	}

	return nil
}

func NewUserService(userRepo repository.UserRepository, tokenServ authenticator.AccessToken, mailServ mailer.MailService) UserService {
	return &userService{
		userRepo:  userRepo,
		tokenServ: tokenServ,
		mailServ:  mailServ,
	}
}
