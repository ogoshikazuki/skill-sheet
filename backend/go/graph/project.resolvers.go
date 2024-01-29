package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"

	"github.com/ogoshikazuki/skill-sheet/di"
	"github.com/ogoshikazuki/skill-sheet/entity"
	"github.com/ogoshikazuki/skill-sheet/graph/model"
	"github.com/ogoshikazuki/skill-sheet/graph/scalar"
	"github.com/ogoshikazuki/skill-sheet/usecase"
)

// Projects is the resolver for the projects field.
func (r *queryResolver) Projects(ctx context.Context, orderBy []*model.ProjectOrder) ([]*model.Project, error) {
	input := usecase.SearchProjectsInput{
		OrderBy: []entity.ProjectOrder{},
	}
	for _, projectOrder := range orderBy {
		var field entity.ProjectOrderField
		switch projectOrder.Field {
		case model.ProjectOrderFieldStartMonth:
			field = entity.PROJECT_ORDER_START_MONTH
		case model.ProjectOrderFieldEndMonth:
			field = entity.PROJECT_ORDER_END_MONTH
		}
		input.OrderBy = append(input.OrderBy, entity.ProjectOrder{
			Field:     field,
			Direction: convertDirectionFromGraphToEntity(projectOrder.Direction),
		})
	}

	output, err := di.Usecases.SearchProjectsUsecase.Handle(ctx, input)
	if err != nil {
		return nil, err
	}

	modelProjects := []*model.Project{}
	for _, project := range output.Projects {
		modelProjects = append(modelProjects, &model.Project{
			ID:         scalar.NewID("Project", project.Id),
			Name:       project.Name,
			StartMonth: project.StartMonth,
			EndMonth:   project.EndMonth,
		})
	}
	return modelProjects, nil
}
