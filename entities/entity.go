package entities

type Entity interface {
	GetID() string
	SetID(id string)
}

type Validator interface {
	Validate() error
}
