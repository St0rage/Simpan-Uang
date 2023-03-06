package mailer

import (
	"strconv"

	"github.com/St0rage/Simpan-Uang/config"
	"github.com/St0rage/Simpan-Uang/utils"
	"gopkg.in/gomail.v2"
)

type MailService interface {
	ResetPasswordMail(email string, password string) error
}

type mailService struct {
	cfg config.MailConfig
}

func NewMailService(config config.MailConfig) MailService {
	return &mailService{
		cfg: config,
	}
}

func (m *mailService) ResetPasswordMail(email string, password string) error {
	mailer := gomail.NewMessage()

	mailer.SetHeader("From", m.cfg.MailUsername)
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", "Simpan Uang Password Reset")
	mailer.SetBody("text/html", "<h4>Password Baru Anda</h4><p><strong>Email : </strong>"+email+"</p><p><strong>Password : </strong>"+password+"</p></br></br><h4>Harap ganti password anda ketika sudah login</h4>")

	port, _ := strconv.Atoi(m.cfg.MailPort)
	dialer := gomail.NewDialer(m.cfg.MailHost, port, m.cfg.MailUsername, m.cfg.MailPassword)

	err := dialer.DialAndSend(mailer)
	utils.PanicIfError(err)

	return nil
}
