package entities

type CreateAccountReq struct {
	ID      string
	OwnerID string
}

type Account struct {
	ID      string
	OwnerID string
	Balance uint64
}
