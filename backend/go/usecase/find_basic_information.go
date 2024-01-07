package usecase

import (
	"context"

	"github.com/ogoshikazuki/skill-sheet/entity"
)

type (
	FindBasicInformationOutput struct {
		BasicInformation entity.BasicInformation
	}

	FindBasicInformationUsecase interface {
		Handle(context.Context) (FindBasicInformationOutput, error)
	}

	findBasicInformationInteractor struct {
		repository entity.BasicInformationRepository
	}
)

func NewFindBasicInformationUsecase(repository entity.BasicInformationRepository) FindBasicInformationUsecase {
	return findBasicInformationInteractor{
		repository: repository,
	}
}

func (interactor findBasicInformationInteractor) Handle(context context.Context) (FindBasicInformationOutput, error) {
	basicInformation, err := interactor.repository.Find(context)
	if err != nil {
		return FindBasicInformationOutput{}, err
	}

	return FindBasicInformationOutput{
		BasicInformation: basicInformation,
	}, nil
}
