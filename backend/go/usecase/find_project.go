package usecase

import (
	"context"

	"github.com/ogoshikazuki/skill-sheet/entity"
)

type (
	FindProjectInput struct {
		ID entity.ID
	}

	FindProjectOutput struct {
		Project entity.Project
	}

	FindProjectUsecase struct {
		repository entity.ProjectRepository
	}
)

func NewFindProjectUsecase(repository entity.ProjectRepository) FindProjectUsecase {
	return FindProjectUsecase{
		repository: repository,
	}
}

func (f FindProjectUsecase) Handle(ctx context.Context, input FindProjectInput) (FindProjectOutput, error) {
	project, err := f.repository.Find(ctx, input.ID)
	if err != nil {
		return FindProjectOutput{}, err
	}

	return FindProjectOutput{
		Project: project,
	}, nil
}
