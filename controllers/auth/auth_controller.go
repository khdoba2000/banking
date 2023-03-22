package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/khdoba2000/banking/configs"
	"github.com/khdoba2000/banking/constants"
	"github.com/khdoba2000/banking/entities"
	e "github.com/khdoba2000/banking/errors"
	"github.com/khdoba2000/banking/logger"
	pkgerrors "github.com/khdoba2000/banking/pkg/errors"
	"github.com/khdoba2000/banking/pkg/jwt"
	"github.com/khdoba2000/banking/pkg/security"
	"github.com/khdoba2000/banking/storage"
)

// AuthController
type AuthController interface {
	Login(ctx context.Context, req entities.LoginReq) (*entities.LoginRes, error)
	Signup(ctx context.Context, req entities.SignupReq) (*entities.SignupRes, error)
	SendCode(ctx context.Context, req entities.SendCodeReq) error
	VerifyCode(ctx context.Context, req entities.VerifyCodeReq) (*entities.VerifyCodeRes, error)
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
			return nil, pkgerrors.NewError(http.StatusForbidden, "phoneNumber or password is wrong")
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
		err = pkgerrors.NewError(http.StatusForbidden, "phoneNumber or password is wrong")
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

	customer := entities.Customer{
		ID:          customerID,
		Name:        req.Name,
		Password:    hashedPassword,
		PhoneNumber: req.PhoneNumber,
	}

	err = ac.storage.Customer().Create(ctx, customer)
	if err != nil {
		ac.log.Error("calling Signup failed", logger.Error(err))
		// if errors.Is(err, e.ErrCustomerAlreadyExists) {
		// 	return nil, e.ErrCustomerAlreadyExists
		// }
		return nil, err
	}

	accountID := uuid.NewString()

	err = ac.storage.Account().Create(ctx, entities.CreateAccountReq{
		ID:      accountID,
		OwnerID: customerID,
	})
	if err != nil {
		ac.log.Error("calling Account.Create failed", logger.Error(err))
		// if errors.Is(err, e.ErrAccountAlreadyExists) {
		// 	return nil, e.ErrAccountAlreadyExists
		// }
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

// SendCode ...
func (ac authController) SendCode(ctx context.Context, req entities.SendCodeReq) error {

	_, err := security.GenerateRandomCode(constants.VerifyCodeLength)
	if err != nil {
		return err
	}
	// send code with sms service provider

	// save the code and the phoneNumber tuple to temporary storage with verified = false status

	return nil
}

// VerifyCode ...
func (ac authController) VerifyCode(ctx context.Context, req entities.VerifyCodeReq) (*entities.VerifyCodeRes, error) {

	// search the code and the phoneNumber tuple within temporary storage

	//if found and code is the same
	// give the tuple status "verified"
	// ......
	//and
	//return a temporary access token to get access to Signup API

	tokenMetadata := map[string]string{
		"phone_number": req.PhoneNumber,
		"role":         constants.CustomerRoleInSignup,
	}

	accessToken, err := jwt.GenerateNewJWTToken(tokenMetadata, constants.JWTAccessTokenExpireDuration, ac.cfg.JWTSecretKey)
	if err != nil {
		ac.log.Error("calling GenerateNewTokens failed", logger.Error(err))
		return nil, err
	}

	return &entities.VerifyCodeRes{AccessToken: accessToken}, nil

	//if found but the code is different
	// return errors.New("invalid code")

	//if not found return error
	// return errors.New("invalid code")
}
