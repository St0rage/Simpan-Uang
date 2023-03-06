package service

import (
	"errors"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/St0rage/Simpan-Uang/model/domain"
	"github.com/St0rage/Simpan-Uang/model/web"
	"github.com/St0rage/Simpan-Uang/repository"
	"github.com/St0rage/Simpan-Uang/utils"
	"github.com/St0rage/Simpan-Uang/utils/authenticator"
	"github.com/St0rage/Simpan-Uang/utils/mailer"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUser(userId string) web.UserResponse
	Register(newUser *web.UserRegisterRequest) (map[string]string, error)
	Login(loginRequest *web.UserLoginRequest) (string, error)
	ForgotPassword(resetRequest *web.UserResetRequest) (map[string]string, error)
	ChangePassword(userId string, changePasswordRequest *web.UserChangePasswordRequest)
	UpdateAvatar(userId string, updateAvatarRequest *web.UserAvatarRequest) (map[string]string, error)
	UpdateUser(userId string, userUpdateRequest *web.UserUpdateRequest) (map[string]string, error)
	CheckAdmin(userId string) bool
	GetBalance(userId string) float32
}

type userService struct {
	userRepo         repository.UserRepository
	piggyBankService PiggyBankService
	wishlistService  WishlistService
	tokenServ        authenticator.AccessToken
	mailServ         mailer.MailService
}

func (userService *userService) GetUser(userId string) web.UserResponse {
	var userResponse web.UserResponse
	user := userService.userRepo.FindById(userId)

	userResponse.Id = user.Id
	userResponse.Name = user.Name
	userResponse.Email = user.Email
	userResponse.IsAdmin = user.IsAdmin
	userResponse.Balance = userService.GetBalance(userId)
	userResponse.Avatar = os.Getenv("DOMAIN") + "/api/user/resources/avatar/" + user.Avatar

	return userResponse
}

func (userService *userService) Register(newUser *web.UserRegisterRequest) (map[string]string, error) {
	emailExist := userService.userRepo.CheckEmail(newUser.Email)
	if emailExist {
		return map[string]string{
			"email": "email sudah digunakan",
		}, errors.New("error")
	}

	var user domain.User
	adminExist := userService.userRepo.CheckAdmin()
	if adminExist {
		user.IsAdmin = false
	} else {
		user.IsAdmin = true
	}

	bytes, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)

	user.Id = utils.GenerateId()
	user.Name = newUser.Name
	user.Email = newUser.Email
	user.Password = string(bytes)
	user.Avatar = "default.png"

	userService.userRepo.Save(&user)

	return nil, nil
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

func (userService *userService) ForgotPassword(resetRequest *web.UserResetRequest) (map[string]string, error) {
	user, err := userService.userRepo.FindByEmail(resetRequest.Email)
	if err != nil {
		return map[string]string{
			"email": "email tidak ditemukan",
		}, errors.New("error")
	}

	generatePassword := strconv.Itoa(rand.Intn(100000))
	password, _ := bcrypt.GenerateFromPassword([]byte(generatePassword), 14)

	user.Password = string(password)

	userService.userRepo.UpdatePassword(&user)

	err = userService.mailServ.ResetPasswordMail(user.Email, generatePassword)
	if err != nil {
		panic(err)
	}
	return nil, nil
}

func (userService *userService) UpdateAvatar(userId string, updateAvatarRequest *web.UserAvatarRequest) (map[string]string, error) {
	user := userService.userRepo.FindById(userId)

	index := strings.Index(updateAvatarRequest.Avatar, "/")
	if index != 10 {
		return map[string]string{
			"avatar": "format gambar tidak valid",
		}, errors.New("error")
	}
	dataImage := updateAvatarRequest.Avatar[0:index]

	if dataImage != "data:image" {
		return map[string]string{
			"avatar": "format gambar tidak valid",
		}, errors.New("error")
	}

	imageName, err := utils.DecodeImage(updateAvatarRequest.Avatar)
	if err != nil {
		return map[string]string{
			"avatar": "format gambar tidak valid",
		}, errors.New("error")
	}

	oldImage := user.Avatar
	user.Avatar = imageName

	userService.userRepo.UpdateAvatar(&user)

	if oldImage != "default.png" {
		os.Remove("./resources/avatar/" + oldImage)
	}

	return nil, nil
}

func (userService *userService) ChangePassword(userId string, changePasswordRequest *web.UserChangePasswordRequest) {
	user := userService.userRepo.FindById(userId)

	bytes, _ := bcrypt.GenerateFromPassword([]byte(changePasswordRequest.Password), 14)

	user.Password = string(bytes)

	userService.userRepo.UpdatePassword(&user)
}

func (userService *userService) UpdateUser(userId string, userUpdateRequest *web.UserUpdateRequest) (map[string]string, error) {
	user := userService.userRepo.FindById(userId)

	if userUpdateRequest.Email != user.Email {
		exist := userService.userRepo.CheckEmail(userUpdateRequest.Email)
		if exist {
			return map[string]string{
				"email": "email sudah digunakan",
			}, errors.New("")
		} else {
			user.Email = userUpdateRequest.Email
		}
	}

	user.Name = userUpdateRequest.Name

	userService.userRepo.Update(&user)

	return nil, nil
}

func (userService *userService) CheckAdmin(userId string) bool {
	return userService.userRepo.IsAdmin(userId)
}

func (userService *userService) GetBalance(userId string) float32 {
	piggyBankTotal := userService.piggyBankService.GetAllPiggyBankTotal(userId)
	wishlistTotal := userService.wishlistService.GetAllWishlistTotal(userId)

	return piggyBankTotal + wishlistTotal
}

func NewUserService(userRepo repository.UserRepository, piggyBankService PiggyBankService, wishlistService WishlistService, tokenServ authenticator.AccessToken, mailServ mailer.MailService) UserService {
	return &userService{
		userRepo:         userRepo,
		piggyBankService: piggyBankService,
		wishlistService:  wishlistService,
		tokenServ:        tokenServ,
		mailServ:         mailServ,
	}
}
