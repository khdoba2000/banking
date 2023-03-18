package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/khdoba/banking/configs"
	"github.com/khdoba/banking/constants"
	"github.com/khdoba/banking/entities"
	"github.com/khdoba/banking/logger"
	e "github.com/khdoba/banking/pkg/errors"
	"github.com/khdoba/banking/pkg/jwt"
	"github.com/khdoba/banking/pkg/security"
	"github.com/khdoba/banking/storage"
)

// AuthController
type AuthController interface {
	Login(ctx context.Context, model entities.LoginReq) (*entities.LoginRes, error)
	Signup(ctx context.Context, model entities.SignupReq) (*entities.SignupRes, error)
}

type authController struct {
	log     logger.LoggerI
	storage storage.Storage
	cfg     *configs.Configuration
}

// NewAuthController ...
func NewAuthController(log logger.LoggerI, storage storage.Storage) AuthController {
	return authController{
		log:     log,
		storage: storage,
		cfg:     configs.Config(),
	}
}

// Login ...
func (ac authController) Login(ctx context.Context, req entities.LoginReq) (*entities.LoginRes, error) {
	customer, err := ac.storage.Customer().GetByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		ac.log.Error("calling GetByPhoneNumber failed", logger.Error(err))
		if errors.Is(err, e.ErrCustomerNotExists) {
			return nil, e.NewError(http.StatusForbidden, "phoneNumber or password is wrong")
		}
		return nil, err
	}

	match, err := security.ComparePassword(customer.Password, req.Password)
	if err != nil {
		ac.log.Error("ComparePassword failed", logger.Error(err))
		return nil, err
	}

	if !match {
		ac.log.Info("ComparePassword failed")
		err := e.NewError(http.StatusForbidden, "phoneNumber or password is wrong")
		return nil, err
	}

	tokenMetadata := map[string]string{
		"id":   customer.ID,
		"role": constants.CustomerRole,
	}

	tokens := entities.Tokens{}
	tokens.AccessToken, err = jwt.GenerateNewJWTToken(tokenMetadata, constants.JWTAccessTokenExpireDuration, ac.cfg.JWTSecretKey)
	if err != nil {
		ac.log.Error("calling GenerateNewTokens failed", logger.Error(err))
		return nil, err
	}

	tokens.RefreshToken, err = jwt.GenerateNewJWTToken(tokenMetadata, constants.JWTRefreshTokenExpireDuration, ac.cfg.JWTSecretKey)
	if err != nil {
		ac.log.Error("calling GenerateNewTokens failed", logger.Error(err))
		return nil, err
	}

	return &entities.LoginRes{
		ID:          customer.ID,
		PhoneNumber: customer.PhoneNumber,
		Name:        customer.Name,
		Tokens:      tokens,
	}, nil
}

// Signup ...
func (ac authController) Signup(ctx context.Context, req entities.SignupReq) (*entities.SignupRes, error) {

	customerID := uuid.NewString()

	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		ac.log.Error("calling HashPassword failed", logger.Error(err))
		return nil, err
	}

	req.ID = customerID
	req.Password = hashedPassword

	_, err = ac.storage.Authenitication().Signup(ctx, req)
	if err != nil {
		ac.log.Error("calling Signup failed", logger.Error(err))
		return nil, err
	}

	tokenMetadata := map[string]string{
		"id":   customerID,
		"role": constants.CustomerRole,
	}

	tokens := entities.Tokens{}
	tokens.AccessToken, err = jwt.GenerateNewJWTToken(tokenMetadata, constants.JWTAccessTokenExpireDuration, ac.cfg.JWTSecretKey)
	if err != nil {
		ac.log.Error("calling GenerateNewTokens failed", logger.Error(err))
		return nil, err
	}

	tokens.RefreshToken, err = jwt.GenerateNewJWTToken(tokenMetadata, constants.JWTRefreshTokenExpireDuration, ac.cfg.JWTSecretKey)
	if err != nil {
		ac.log.Error("calling GenerateNewTokens failed", logger.Error(err))
		return nil, err
	}

	return &entities.SignupRes{
		ID:     customerID,
		Tokens: tokens,
	}, nil
}
