package storage

import (
	"log"

	"github.com/khdoba2000/banking/configs"
	"github.com/khdoba2000/banking/db"
	"github.com/khdoba2000/banking/storage/postgres"
	"github.com/khdoba2000/banking/storage/repo"
)

type Storage interface {
	Customer() repo.ICustomerStorage
	Account() repo.IAccountStorage
	Transaction() repo.ITransactionStorage
	// other storage interfaces
	//
}

type storage struct {
	customerRepo    repo.ICustomerStorage
	accountRepo     repo.IAccountStorage
	transactionRepo repo.ITransactionStorage
}

// New
func New(cfg *configs.Configuration) *storage {

	postgresDB, err := db.Init(cfg)
	if err != nil {
		log.Fatalf("error connecting to postgres database: %v", err)
	}
	return &storage{
		customerRepo:    postgres.NewCustomer(postgresDB),
		accountRepo:     postgres.NewAccount(postgresDB),
		transactionRepo: postgres.NewTransaction(postgresDB),
	}
}

// Customer returns customer repository
func (s storage) Customer() repo.ICustomerStorage {
	return s.customerRepo
}

// Account returns account repository
func (s storage) Account() repo.IAccountStorage {
	return s.accountRepo
}

// Transaction returns transaction repository
func (s storage) Transaction() repo.ITransactionStorage {
	return s.transactionRepo
}
