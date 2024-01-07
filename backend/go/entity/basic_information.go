package entity

import "context"

type BasicInformation struct {
	Birthday Date
}

type BasicInformationRepository interface {
	Find(context.Context) (BasicInformation, error)
}
