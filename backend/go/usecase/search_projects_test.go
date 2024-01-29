package usecase_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/ogoshikazuki/skill-sheet/entity"
	"github.com/ogoshikazuki/skill-sheet/usecase"
)

func TestSearchProjectsUsecase(t *testing.T) {
	projectOrders := []entity.ProjectOrder{
		{
			Field:     entity.PROJECT_ORDER_START_MONTH,
			Direction: entity.ASC,
		},
		{
			Field:     entity.PROJECT_ORDER_END_MONTH,
			Direction: entity.DESC,
		},
	}
	yearMonth202301, _ := entity.NewYearMonth(2023, 1)
	yearMonth202312, _ := entity.NewYearMonth(2023, 12)
	yearMonth202401, _ := entity.NewYearMonth(2024, 1)
	projects := []entity.Project{
		{
			Id:         1,
			Name:       "test project 1",
			StartMonth: yearMonth202301,
			EndMonth:   yearMonth202312,
		},
		{
			Id:         2,
			Name:       "test project 2",
			StartMonth: yearMonth202401,
			EndMonth:   nil,
		},
	}
	testErr := errors.New("test")

	tests := map[string]struct {
		input            usecase.SearchProjectsInput
		repositoryReturn []entity.Project
		repositoryErr    error
		expectedOutput   usecase.SearchProjectsOutput
		expectedErr      error
	}{
		"Success": {
			input: usecase.SearchProjectsInput{
				OrderBy: projectOrders,
			},
			repositoryReturn: projects,
			repositoryErr:    nil,
			expectedOutput: usecase.SearchProjectsOutput{
				Projects: projects,
			},
			expectedErr: nil,
		},
		"RepositoryReturnsError": {
			input: usecase.SearchProjectsInput{
				OrderBy: projectOrders,
			},
			repositoryReturn: nil,
			repositoryErr:    testErr,
			expectedOutput:   usecase.SearchProjectsOutput{},
			expectedErr:      testErr,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			repository := entity.ProjectRepositoryMock{
				SearchFunc: func(contextMoqParam context.Context, projectOrders []entity.ProjectOrder) ([]entity.Project, error) {
					if !reflect.DeepEqual(projectOrders, tt.input.OrderBy) {
						t.Errorf("projectOrders: %+v, tt.input.OrderBy: %+v", projectOrders, tt.input.OrderBy)
					}
					return tt.repositoryReturn, tt.repositoryErr
				},
			}

			ctx := context.Background()
			output, err := usecase.NewSearchProjectsUsecase(&repository).Handle(ctx, tt.input)
			if len(repository.SearchCalls()) != 1 {
				t.Errorf("len(repository.SearchCalls()): %d", len(repository.SearchCalls()))
			}
			if !reflect.DeepEqual(output, tt.expectedOutput) {
				t.Errorf("tt.expectedOutput: %+v, output: %+v", tt.expectedOutput, output)
			}
			if err != tt.expectedErr {
				t.Errorf("tt.expectedErr: %+v, err: %+v", tt.expectedErr, err)
			}
		})
	}
}
