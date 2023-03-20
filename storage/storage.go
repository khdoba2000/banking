package storage

import (
	"log"

	"github.com/khdoba/banking/configs"
	"github.com/khdoba/banking/db"
	"github.com/khdoba/banking/storage/postgres"
	"github.com/khdoba/banking/storage/repo"
)

type Storage interface {
	Customer() repo.ICustomerStorage
	Account() repo.IAccountStorage
	// other storage interfaces
	//
}

type storagePg struct {
	customerRepo repo.ICustomerStorage
	accountRepo  repo.IAccountStorage
}

// New
func New(cfg *configs.Configuration) *storagePg {

	postgresDB, err := db.Init(cfg)
	if err != nil {
		log.Fatalf("error connecting to postgres database: %v", err)
	}
	return &storagePg{
		customerRepo: postgres.NewCustomer(postgresDB),
		accountRepo:  postgres.NewAccount(postgresDB),
	}
}

// Customer returns customer repository
func (s storagePg) Customer() repo.ICustomerStorage {
	return s.customerRepo
}

// Account returns account repository
func (s storagePg) Account() repo.IAccountStorage {
	return s.accountRepo
}
