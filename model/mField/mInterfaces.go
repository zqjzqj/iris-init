package mField

type OwnerInterface interface {
	TableName() string
	GetOwnerID() uint64
	GetOwnerCategory() string
}
