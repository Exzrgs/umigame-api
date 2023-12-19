package services

// import (
// 	"net/mail"
// 	"os"
// 	"golang.org/x/crypto/bcrypt"

// 	"github.com/google/uuid"

// 	"umigame-api/myerrors"
// 	"umigame-api/repositories"
// 	"umigame-api/utils"
// 	"umigame-api/models"
// )

// func (s *Service) RegisterUserService(auth models.Auth) error {
// 	if _, err := mail.ParseAddress(auth.Email); err != nil {
// 		err = myerrors.EmailInvalid.Wrap(err, "email invalid")
// 		return err
// 	}

// 	hash, err := utils.PasswordEncrypt(auth.Password)
// 	if err != nil {
// 		err = myerrors.EncryptFailed.Wrap(err, "internal server error")
// 		return err
// 	}
// 	auth.Hash = hash

// 	uuidObj, err := uuid.NewUUID()
// 	if err != nil {
// 		err = myerrors.GenUUIDFailed.Wrap(err, "internal server error")
// 		return err
// 	}
// 	mailAuthUuid := uuidObj.String()
// 	auth.Uuid = mailAuthUuid

// 	if err = repositories.RegisterUser(s.db, &auth); err != nil {
// 		return err
// 	}

// 	mailMessage := utils.MailMessage(s.port, mailAuthUuid)
// 	myMail := utils.Mail{
// 		Host:     "smtp.gmail.com",
// 		Port:     "587",
// 		From:     os.Getenv("EMAIL_ADDRESS"),
// 		Password: os.Getenv("EMAIL_PASSWORD"),
// 		To:       []string{auth.Email},
// 		Subject:  "メールアドレスの認証",
// 		Message:  mailMessage,
// 	}
// 	if err := myMail.SendMail(); err != nil {
// 		err = myerrors.SendMailFailed.Wrap(err, "internal server error")
// 		return err
// 	}

// 	return nil
// }

// func (s *Service) MailCheckService(uuid string) error {
// 	err := repositories.UpdateActivate(s.db, uuid)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s *Service) LoginService(email string, password string) (string, error) {
// 	auth, err := repositories.GetAuthInfo(s.db, email)
// 	if err != nil {
// 		return "", err
// 	}

// 	if !auth.ActivateFlag {
// 		err := myerrors.NotActivate.Wrap(NotActivate, "email is not valified")
// 		return "", err
// 	}

// 	if err := bcrypt.CompareHashAndPassword([]byte(auth.Hash), []byte(password)); err != nil {
// 		err = myerrors.InvalidPassword.Wrap(err, "password is not correnct")
// 		return "", err
// 	}

// 	return auth.Uuid, nil
// }
