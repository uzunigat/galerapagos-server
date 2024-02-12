package model

type ServiceValidator interface {
	ValidateStruct(s interface{}) error
}
