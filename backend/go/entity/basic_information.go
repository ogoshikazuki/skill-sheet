package entity

//go:generate go run github.com/matryer/moq -out basic_information_mock.go . BasicInformationRepository

import "context"

type BasicInformation struct {
	Birthday           Date
	Gender             Gender
	AcademicBackground string
}

type (
	UpdateBasicInformationInput struct {
		Birthday           UpdateBirthdayInput
		Gender             UpdateGenderInput
		AcademicBackground UpdateAcademicBackgroundInput
	}
	UpdateBirthdayInput struct {
		Birthday  Date
		IsUpdated bool
	}
	UpdateGenderInput struct {
		Gender    Gender
		IsUpdated bool
	}
	UpdateAcademicBackgroundInput struct {
		AcademicBackground string
		IsUpdated          bool
	}
)

type BasicInformationRepository interface {
	Find(context.Context) (BasicInformation, error)
	Update(ctx context.Context, tx Tx, input UpdateBasicInformationInput) error
}
