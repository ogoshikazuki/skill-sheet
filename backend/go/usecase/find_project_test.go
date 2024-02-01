package usecase_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/ogoshikazuki/skill-sheet/entity"
	"github.com/ogoshikazuki/skill-sheet/usecase"
)

func TestFindProjectUsecase(t *testing.T) {
	var inputID entity.ID = 1
	startMonth, _ := entity.NewYearMonth(2024, 1)
	project := entity.Project{
		Id:         inputID,
		Name:       "project",
		StartMonth: startMonth,
	}

	ctx := context.Background()

	repository := &entity.ProjectRepositoryMock{
		FindFunc: func(contextMoqParam context.Context, id entity.ID) (entity.Project, error) {
			if id != inputID {
				t.Errorf("id: %d, inputID: %d", id, inputID)
			}
			return project, nil
		},
	}

	input := usecase.FindProjectInput{
		ID: inputID,
	}

	findProjectUsecase := usecase.NewFindProjectUsecase(repository)
	output, err := findProjectUsecase.Handle(ctx, input)
	if len(repository.FindCalls()) != 1 {
		t.Errorf("repository.FindCalls(): %d", len(repository.FindCalls()))
	}
	if !reflect.DeepEqual(output.Project, project) {
		t.Errorf("output: %+v, project: %+v", output, project)
	}
	if err != nil {
		t.Errorf("err: %s", err)
	}
}
