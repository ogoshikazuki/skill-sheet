package usecases

import (
	"context"

	"github.com/ogoshikazuki/skill-sheet/entities"
)

type (
	FindBasicInformationOutput struct {
		basicInformation entities.BasicInformation
	}

	FindBasicInformationUsecase interface {
		Handle(context.Context) (FindBasicInformationOutput, error)
	}

	findBasicInformationInteractor struct {
		repository entities.BasicInformationRepository
	}
)

func NewFindBasicInformationUsecase(repository entities.BasicInformationRepository) FindBasicInformationUsecase {
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
		basicInformation: basicInformation,
	}, nil
}
