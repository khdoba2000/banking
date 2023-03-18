package storage

import (
	"log"

	"github.com/khdoba/banking/configs"
	"github.com/khdoba/banking/db"
	"github.com/khdoba/banking/storage/postgres"
	"github.com/khdoba/banking/storage/repo"
)

type Storage interface {
	Authenitication() repo.IAuthStorage
	Customer() repo.ICustomerStorage
	// other storage interfaces
	//
}

type storagePg struct {
	authRepo     repo.IAuthStorage
	customerRepo repo.ICustomerStorage
}

// New
func New(cfg *configs.Configuration) *storagePg {

	postgresDB, err := db.Init(cfg)
	if err != nil {
		log.Fatalf("error connecting to postgres database: %v", err)
	}
	return &storagePg{
		authRepo:     postgres.NewAuth(postgresDB),
		customerRepo: postgres.NewCustomer(postgresDB),
	}
}

// Authenitication returns authenication repository
func (s storagePg) Authenitication() repo.IAuthStorage {
	return s.authRepo
}

// Customer returns customer repository
func (s storagePg) Customer() repo.ICustomerStorage {
	return s.customerRepo
}
