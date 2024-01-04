package entities

import "context"

type (
	BasicInformation struct {
		Birthday Birthday
	}

	Birthday string

	BasicInformationRepository interface {
		Find(context.Context) (BasicInformation, error)
	}
)
