package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	config "github.com/amenshenin/go_auth/internal/configs"
	"github.com/amenshenin/go_auth/internal/repository"
	"github.com/amenshenin/go_auth/internal/schemas"
	"github.com/dgrijalva/jwt-go"
)

type Auth struct {
	cfg *config.Config
	r   *repository.Repository
}

func NewAuth(cfg *config.Config, r *repository.Repository) *Auth {
	return &Auth{
		cfg: cfg,
		r:   r,
	}
}

func (a *Auth) CreateUser(user *schemas.UserInput) (int, error) {
	err := user.ValidateUserInput()
	if err != nil {
		return 0, err
	}
	hash := a.generatePasswordHash(user.Password)
	user.Status = StatusDisabled
	id, err := a.r.Autorization.CreateUser(user, hash)
	if err != nil {
		return 0, err
	}
	return id, nil
}

type tokenClaims struct {
	UserId int `json:"Id"`
	jwt.StandardClaims
}

func (a *Auth) GenerateToken(user *schemas.UserInput) (string, error) {
	err := user.ValidateUserInput()
	if err != nil {
		return "", err
	}
	hash := a.generatePasswordHash(user.Password)
	user.Status = StatusActive
	u, err := a.r.Autorization.GetUser(user, hash)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		u.Id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(a.cfg.App.TockenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})
	return token.SignedString([]byte(a.cfg.App.SigningKey))
}

func (a *Auth) ParceToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid singing method")
		}
		return []byte(a.cfg.App.SigningKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token clames are not of type *tokenClaims")
	}
	return claims.UserId, nil
}

func (a *Auth) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(a.cfg.App.SigningKey)))
}
