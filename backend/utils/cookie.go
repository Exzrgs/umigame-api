package utils

import (
	"net/http"
	"os"
	"strconv"
	"time"
)

func GetCookie(uuid string) (*http.Cookie, error) {
	cookieSecure := os.Getenv("COOKIE_SECURE")
	s, err := strconv.ParseBool(cookieSecure)
	if err != nil {
		return nil, err
	}

	return &http.Cookie{
		Name:     "uuid",
		Value:    uuid,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		Secure:   s,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}, nil
}
