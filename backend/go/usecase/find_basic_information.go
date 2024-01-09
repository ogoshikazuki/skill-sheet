package usecase

import (
	"context"

	"github.com/ogoshikazuki/skill-sheet/entity"
)

type (
	FindBasicInformationOutput struct {
		BasicInformation entity.BasicInformation
	}

	FindBasicInformationUsecase struct {
		repository entity.BasicInformationRepository
	}
)

func NewFindBasicInformationUsecase(repository entity.BasicInformationRepository) FindBasicInformationUsecase {
	return FindBasicInformationUsecase{
		repository: repository,
	}
}

func (f FindBasicInformationUsecase) Handle(context context.Context) (FindBasicInformationOutput, error) {
	basicInformation, err := f.repository.Find(context)
	if err != nil {
		return FindBasicInformationOutput{}, err
	}

	return FindBasicInformationOutput{
		BasicInformation: basicInformation,
	}, nil
}
