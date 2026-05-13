package auth

import (
	"context"
	"errors"
	"time"

	"toki/internal/domain/user"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("SUPER_SECRET_KEY")

type Usecase interface {
	Register(ctx context.Context, user *user.User) error
	Login(ctx context.Context, email, password string) (string, error)
}

type authUsecase struct {
	userRepo user.Repository
}

func NewUsecase(
	userRepo user.Repository,
) Usecase {
	return &authUsecase{
		userRepo: userRepo,
	}
}

func (u *authUsecase) Register(
	ctx context.Context,
	userData *user.User,
) error {

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(userData.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	userData.Password = string(hashedPassword)

	return u.userRepo.Create(ctx, userData)
}

func (u *authUsecase) Login(
	ctx context.Context,
	email string,
	password string,
) (string, error) {

	userData, err := u.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(userData.Password),
		[]byte(password),
	)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userData.ID,
		"role":    userData.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
