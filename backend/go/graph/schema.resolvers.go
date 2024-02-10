package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"fmt"

	"github.com/ogoshikazuki/skill-sheet/entity"
	"github.com/ogoshikazuki/skill-sheet/graph/model"
	"github.com/ogoshikazuki/skill-sheet/graph/scalar"
	"github.com/ogoshikazuki/skill-sheet/usecase"
)

// UpdateBasicInformation is the resolver for the updateBasicInformation field.
func (r *mutationResolver) UpdateBasicInformation(ctx context.Context, input model.UpdateBasicInformationInput) (*model.UpdateBasicInformationPayload, error) {
	panic(fmt.Errorf("not implemented: UpdateBasicInformation - updateBasicInformation"))
}

// BasicInformation is the resolver for the basicInformation field.
func (r *queryResolver) BasicInformation(ctx context.Context) (*model.BasicInformation, error) {
	output, err := r.findBasicInformationUsecase.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var gender model.Gender
	switch output.BasicInformation.Gender {
	case entity.Male:
		gender = model.GenderMale
	case entity.Female:
		gender = model.GenderFemale
	}
	return &model.BasicInformation{
		ID:                 scalar.NewID("BasicInformation", 0),
		Birthday:           output.BasicInformation.Birthday,
		Gender:             gender,
		AcademicBackground: output.BasicInformation.AcademicBackground,
	}, nil
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id scalar.ID) (model.Node, error) {
	typename := id.GetTypename()

	switch typename {
	case "Project":
		output, err := r.findProjectUsecase.Handle(ctx, usecase.FindProjectInput{
			ID: id.GetID(),
		})
		if err != nil {
			return nil, err
		}

		technologyIDs := []scalar.ID{}
		for _, technologyID := range output.Project.TechnologyIDs {
			technologyIDs = append(technologyIDs, scalar.NewID(
				"Technology",
				technologyID,
			))
		}

		return model.Project{
			ID:            id,
			Name:          output.Project.Name,
			StartMonth:    output.Project.StartMonth,
			EndMonth:      output.Project.EndMonth,
			TechnologyIDs: technologyIDs,
		}, nil
	}

	return nil, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
