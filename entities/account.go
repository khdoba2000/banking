package entities

type CreateAccountReq struct {
	ID      string
	OwnerID string
}

type Account struct {
	ID           string
	CurrencyCode string
	OwnerID      string `json:"owner_id,omitempty"`
	Balance      uint64
}
