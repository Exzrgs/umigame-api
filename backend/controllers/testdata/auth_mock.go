package testdata

import (
	"umigame-api/models"
)

func (*serviceMock) RegisterUserService(auth models.Auth) error {
	return nil
}

func (*serviceMock) MailCheckService(uuid string) error {
	return nil
}

func (*serviceMock) LoginService(email string, password string) (string, error) {
	return "", nil
}
