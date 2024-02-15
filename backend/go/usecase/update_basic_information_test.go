package usecase_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/ogoshikazuki/skill-sheet/entity"
	"github.com/ogoshikazuki/skill-sheet/usecase"
)

func TestUpdateBasicInformationUsecase(t *testing.T) {
	t.Parallel()

	birthday := entity.Date("1991-07-01")
	gender := entity.Male
	academicBackground := "上智大学卒業"

	updateBasicInformationInput := entity.UpdateBasicInformationInput{
		Birthday: entity.UpdateBirthdayInput{
			Birthday: birthday,
		},
		Gender: entity.UpdateGenderInput{
			Gender: gender,
		},
		AcademicBackground: entity.UpdateAcademicBackgroundInput{
			AcademicBackground: academicBackground,
		},
	}
	basicInformation := entity.BasicInformation{
		Birthday:           birthday,
		Gender:             gender,
		AcademicBackground: academicBackground,
	}

	testErr := errors.New("test")

	tests := map[string]struct {
		repositoryUpdateErr            error
		repositoryfindOutput           entity.BasicInformation
		repositoryfindErr              error
		expectedRepositoryFindCallsLen int
		expectedBasicInformation       entity.BasicInformation
		expectedErr                    error
	}{
		"Normal": {
			repositoryfindOutput:           basicInformation,
			expectedRepositoryFindCallsLen: 1,
			expectedBasicInformation:       basicInformation,
		},
		"RepositoryUpdateReturnsErr": {
			repositoryUpdateErr:            testErr,
			expectedRepositoryFindCallsLen: 0,
			expectedErr:                    testErr,
		},
		"RepositoryFindReturnsErr": {
			repositoryfindErr:              testErr,
			expectedRepositoryFindCallsLen: 1,
			expectedErr:                    testErr,
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			repository := entity.BasicInformationRepositoryMock{
				UpdateFunc: func(ctx context.Context, tx entity.Tx, input entity.UpdateBasicInformationInput) error {
					if !reflect.DeepEqual(input, updateBasicInformationInput) {
						t.Errorf("input: %+v, updateBasicInformationInput: %+v", input, updateBasicInformationInput)
					}
					if input != updateBasicInformationInput {
						t.Error()
					}
					return tt.repositoryUpdateErr
				},
				FindFunc: func(contextMoqParam context.Context) (entity.BasicInformation, error) {
					return tt.repositoryfindOutput, tt.repositoryfindErr
				},
			}
			tx := entity.TxMock{}
			transactionController := entity.TransactionControllerMock{
				TransactionFunc: func(ctx context.Context, f func(tx entity.Tx) error) error {
					return f(tx)
				},
			}

			ctx := context.Background()

			basicInformation, err := usecase.NewUpdateBasicInformationUsecase(&repository, &transactionController).Handle(ctx, updateBasicInformationInput)
			transactionCalls := transactionController.TransactionCalls()
			if len(transactionCalls) != 1 {
				t.Errorf("transactionCalls: %+v", transactionCalls)
			}
			updateCalls := repository.UpdateCalls()
			if len(updateCalls) != 1 {
				t.Errorf("updateCalls: %+v", updateCalls)
			}
			findCalls := repository.FindCalls()
			if len(findCalls) != tt.expectedRepositoryFindCallsLen {
				t.Errorf("findCalls: %+v, tt.expectedRepositoryFindCallsLen: %d", findCalls, tt.expectedRepositoryFindCallsLen)
			}
			if !reflect.DeepEqual(basicInformation, tt.expectedBasicInformation) {
				t.Errorf("basicInformation: %+v, tt.expectedBasicInformation: %+v", basicInformation, tt.expectedBasicInformation)
			}
			if err != tt.expectedErr {
				t.Errorf("err: %+v, tt.expectedErr: %+v", err, tt.expectedErr)
			}
		})
	}
}
