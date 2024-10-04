package service

import (
	"fmt"
	"ministry/internal/app/repository"
	"ministry/internal/model"
	"ministry/utils"
)

type IAuthorization interface {
	SignUp(univer *model.University) (int, error)
	SignIn(univer *model.University) (*model.TokenPair, error)
	AdminSignIn(admin *model.Admin) (*model.TokenPair, error)
}

type AuthService struct {
	rep repository.IAuthorization
}

func NewAuthService(rep repository.IAuthorization) *AuthService {
	return &AuthService{rep: rep}
}

func (a *AuthService) SignUp(univer *model.University) (int, error) {
	hashPassword, err := utils.GenerateHashedPassword(univer.Password)
	if err != nil {
		return 0, err
	}
	univer.Password = hashPassword
	return a.rep.SignUp(univer)
}

func (a *AuthService) SignIn(univer *model.University) (*model.TokenPair, error) {
	univerFromDB, err := a.rep.SignIn(univer)
	if err != nil {
		return nil, fmt.Errorf("error signing in: %w", err)
	}

	if ok := utils.CompareHashAndPassword(univerFromDB.Password, univer.Password); !ok {
		return nil, errPasswordMismatch
	}

	tokens, err := utils.GenerateToken(univerFromDB)
	if err != nil {
		return nil, fmt.Errorf("error generating tokens: %w", err)
	}

	//if err := a.rep.SaveRefreshToken(univerFromDB.ID, tokens.RefreshToken); err != nil {
	//	return nil, fmt.Errorf("error saving refresh token: %w", err)
	//}

	return tokens, nil
}

func (a *AuthService) AdminSignIn(admin *model.Admin) (*model.TokenPair, error) {
	adminFromDB, err := a.rep.AdminSignIn(admin)
	if err != nil {
		return nil, fmt.Errorf("error signing in: %w", err)
	}

	if ok := utils.CompareHashAndPassword(adminFromDB.Password, admin.Password); !ok {
		return nil, errPasswordMismatch
	}

	return utils.GenerateToken(admin)
}
