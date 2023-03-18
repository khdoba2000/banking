package auth

import (
	"context"

	"github.com/khdoba/banking/configs"
	"github.com/khdoba/banking/constants"
	"github.com/khdoba/banking/entities"
	"github.com/khdoba/banking/logger"
	"github.com/khdoba/banking/pkg/jwt"
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

func NewAuthController(log logger.LoggerI, storage storage.Storage) AuthController {
	return authController{
		log:     log,
		storage: storage,
		cfg:     configs.Config(),
	}
}

// Login ...
func (ac authController) Login(ctx context.Context, req entities.LoginReq) (*entities.LoginRes, error) {
	res, err := ac.storage.Authenitication().Login(ctx, req)
	if err != nil {
		ac.log.Error("calling Login gateway failed", logger.Error(err))
		return nil, err
	}

	tokenMetadata := map[string]string{
		"id":   res.ID.String(),
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

	res.Tokens = tokens

	return &res, nil
}

// Login ...
func (ac authController) Signup(ctx context.Context, req entities.SignupReq) (*entities.SignupRes, error) {
	res, err := ac.storage.Authenitication().Signup(ctx, req)
	if err != nil {
		ac.log.Error("calling Login gateway failed", logger.Error(err))
		return nil, err
	}

	tokenMetadata := map[string]string{
		"id":   res.ID.String(),
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

	res.Tokens = tokens
	return &res, nil
}
