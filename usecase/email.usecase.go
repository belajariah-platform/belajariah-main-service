package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/utils"
	"strings"

	"gopkg.in/gomail.v2"
)

type emailUsecase struct {
	mailConfig      *model.Config
	userRepository  repository.UserRepository
	emailRepository repository.EmailRepository
}

type EmailUsecase interface {
	SendEmail(email model.EmailBody)
}

func InitEmailUsecase(mailConfig *model.Config, userRepository repository.UserRepository, emailRepository repository.EmailRepository) EmailUsecase {
	return &emailUsecase{
		mailConfig,
		userRepository,
		emailRepository,
	}
}

func (emailUsecase *emailUsecase) SendEmail(email model.EmailBody) {
	var bodyTemp, subject string

	user, _ := emailUsecase.userRepository.GetUserData(email.UserCode)
	dataEmail := model.EmailBody{
		//general
		UserName:    user.FullName.String,
		ExpiredDate: email.ExpiredDate,
		//registration
		UserEmail:        email.UserEmail,
		VerificationCode: email.VerificationCode,
		//payment
		InvoiceNumber:     email.InvoiceNumber,
		PaymentMethod:     email.PaymentMethod,
		AccountName:       email.AccountName,
		AccountNumber:     email.AccountNumber,
		ClassName:         email.ClassName,
		ClassPrice:        email.ClassPrice,
		TotalConsultation: email.TotalConsultation,
		TotalWebinar:      email.TotalWebinar,
		TotalTransfer:     email.TotalTransfer,
		PromoDiscount:     email.PromoDiscount + "%",
		PromoPrice:        email.ClassPrice - email.TotalTransfer,
	}

	switch strings.ToLower(email.BodyTemp) {
	case "registration success":
		bodyTemp = utils.TemplateRegisterSuccess(dataEmail)
		subject = "Pendaftaran berhasil : selamat bergabung di Belajariah"
	case "change password":
		bodyTemp = utils.TemplateChangePassword(dataEmail)
		subject = "Password belajariahmu berhasil diubah"
	case "account verification":
		bodyTemp = utils.TemplateAccountVerification(dataEmail)
		subject = "Verifikasi akun"
	case "class has been expired":
		bodyTemp = utils.TemplateClassHasBeenExpired(dataEmail)
		subject = "Kelasmu telah berakhir"
	case "1 days before class expired":
		dataEmail.Count = 1
		bodyTemp = utils.TemplateBeforeClassExpired(dataEmail)
		subject = "1 hari lagi kelasmu akan berakhir"
	case "2 days before class expired":
		dataEmail.Count = 2
		bodyTemp = utils.TemplateBeforeClassExpired(dataEmail)
		subject = "2 hari lagi kelasmu akan berakhir"
	case "5 days before class expired":
		dataEmail.Count = 5
		bodyTemp = utils.TemplateBeforeClassExpired(dataEmail)
		subject = "5 hari lagi kelasmu akan berakhir"
	case "7 days before class expired":
		dataEmail.Count = 7
		bodyTemp = utils.TemplateBeforeClassExpired(dataEmail)
		subject = "7 hari lagi kelasmu akan berakhir"
	case "waiting for payment":
		bodyTemp = utils.TemplateWaitingPayment(dataEmail)
		subject = "Menunggu pembayaran"
	case "2 hour before payment expired":
		bodyTemp = utils.TemplatePayment2HoursBeforeExpired(dataEmail)
		subject = "2 jam lagi waktu pembayaranmu akan habis"
	case "payment success":
		bodyTemp = utils.TemplatePaymentSuccess(dataEmail)
		subject = "Pembayaran berhasil"
	case "payment failed":
		bodyTemp = utils.TemplatePaymentFailed(dataEmail)
		subject = "Pembayaran gagal"
	case "payment canceled":
		bodyTemp = utils.TemplatePaymentCanceled(dataEmail)
		subject = "Pembayaran dibatalkan"
	case "payment revised":
		bodyTemp = utils.TemplatePaymentRevised(dataEmail)
		subject = "Perbaiki pembayarnmu"
	}
	mail := gomail.NewMessage()
	mail.SetHeader("From", emailUsecase.mailConfig.Mail.SenderName)
	mail.SetHeader("To", "rikijenifer15@gmail.com")
	mail.SetAddressHeader("Cc", user.Email, "fredidermaga98@gmail.com")
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", bodyTemp)
	// mail.Attach("./sample.png")

	dialer := gomail.NewDialer(
		emailUsecase.mailConfig.Mail.SMTPHost,
		emailUsecase.mailConfig.Mail.SMTPPort,
		emailUsecase.mailConfig.Mail.AuthEmail,
		emailUsecase.mailConfig.Mail.AuthPassword,
	)

	err := dialer.DialAndSend(mail)
	if err != nil {
		utils.PushLogf("ERROR SEND EMAIL : ", err.Error())
	} else {
		utils.PushLogf("SUCCESS SEND EMAIL : ", mail)
	}
}
