package usecase

import (
	"context"

	"github.com/ogoshikazuki/skill-sheet/entity"
)

type (
	SearchProjectsInput struct {
		OrderBy []entity.ProjectOrder
	}

	SearchProjectsOutput struct {
		Projects []entity.Project
	}

	SearchProjectsUsecase struct {
		repository entity.ProjectRepository
	}
)

func NewSearchProjectsUsecase(repository entity.ProjectRepository) SearchProjectsUsecase {
	return SearchProjectsUsecase{
		repository: repository,
	}
}

func (s SearchProjectsUsecase) Handle(ctx context.Context, searchProjectsInput SearchProjectsInput) (SearchProjectsOutput, error) {
	projects, err := s.repository.Search(ctx, searchProjectsInput.OrderBy)
	if err != nil {
		return SearchProjectsOutput{}, err
	}

	return SearchProjectsOutput{
		Projects: projects,
	}, err
}
