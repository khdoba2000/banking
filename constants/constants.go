package constants

import "time"

const (
	PGForeignKeyViolationCode = "23503"
	PGUniqueKeyViolationCode  = "23505"
)

const (
	TestMode                      = "test"
	DebugMode                     = "debug"
	JWTRefreshTokenExpireDuration = time.Hour * 24
	JWTAccessTokenExpireDuration  = time.Minute * 30
	CustomerRoleInSignup          = "customer_in_signup"
	CustomerRole                  = "customer"
	VerifyCodeLength              = 6
)

const (
	IncomeTransactionID = iota
	ExpenseTransactionID
	TransferTransactionID
)
