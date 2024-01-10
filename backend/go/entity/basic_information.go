package entity

//go:generate go run github.com/matryer/moq -out basic_information_mock.go . BasicInformationRepository

import "context"

type BasicInformation struct {
	Birthday Date
}

type BasicInformationRepository interface {
	Find(context.Context) (BasicInformation, error)
}
