package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/ogoshikazuki/skill-sheet/entity"
	"github.com/ogoshikazuki/skill-sheet/usecase"
)

func TestFindBasicInformationUsecase(t *testing.T) {
	birthday := entity.Date("1991-07-01")
	basicInformation := entity.BasicInformation{
		Birthday: birthday,
	}
	testErr := errors.New("test")

	tests := map[string]struct {
		findFunc       func(contextMoqParam context.Context) (entity.BasicInformation, error)
		expectedOutput usecase.FindBasicInformationOutput
		expectedErr    error
	}{
		"Success": {
			findFunc: func(contextMoqParam context.Context) (entity.BasicInformation, error) {
				return basicInformation, nil
			},
			expectedOutput: usecase.FindBasicInformationOutput{
				BasicInformation: entity.BasicInformation{
					Birthday: birthday,
				},
			},
			expectedErr: nil,
		},
		"RepositoryReturnsError": {
			findFunc: func(contextMoqParam context.Context) (entity.BasicInformation, error) {
				return entity.BasicInformation{}, testErr
			},
			expectedOutput: usecase.FindBasicInformationOutput{},
			expectedErr:    testErr,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			repository := entity.BasicInformationRepositoryMock{
				FindFunc: tt.findFunc,
			}

			ctx := context.Background()
			output, err := usecase.NewFindBasicInformationUsecase(&repository).Handle(ctx)
			if len(repository.FindCalls()) != 1 {
				t.Errorf("actual: %d", len(repository.FindCalls()))
			}
			if output != tt.expectedOutput {
				t.Errorf("expected: %+v, actual: %+v", tt.expectedOutput, output)
			}
			if err != tt.expectedErr {
				t.Errorf("expected: %+v, actual: %+v", tt.expectedErr, err)
			}
		})
	}
}
