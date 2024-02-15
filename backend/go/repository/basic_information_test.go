package repository_test

import (
	"context"
	"testing"

	"github.com/ogoshikazuki/skill-sheet/entity"
	"github.com/ogoshikazuki/skill-sheet/repository"
)

func TestBasicInformationRepositoryFind(t *testing.T) {
	academicBackground := "上智大学卒業"

	tests := map[string]struct {
		sqlHandler repository.SqlHandler
		gender     string
		expect     entity.BasicInformation
		returnsErr bool
	}{
		"SuccessMale": {
			sqlHandler: sqlHandler,
			gender:     "MALE",
			expect: entity.BasicInformation{
				Birthday:           "1991-07-01",
				Gender:             entity.Male,
				AcademicBackground: academicBackground,
			},
			returnsErr: false,
		},
		"SuccessFemale": {
			sqlHandler: sqlHandler,
			gender:     "FEMALE",
			expect: entity.BasicInformation{
				Birthday:           "1991-07-01",
				Gender:             entity.Female,
				AcademicBackground: academicBackground,
			},
			returnsErr: false,
		},
		"SqlHandlerReturnsErr": {
			sqlHandler: errSqlHandler,
			expect:     entity.BasicInformation{},
			returnsErr: true,
		},
		"InvalidGender": {
			sqlHandler: sqlHandler,
			gender:     "INVALID",
			expect:     entity.BasicInformation{},
			returnsErr: true,
		},
	}
	for name, tt := range tests {
		ctx := context.Background()
		if _, err := sqlHandler.ExecContext(ctx, `UPDATE "basic_information" SET "gender" = $1`, tt.gender); err != nil {
			panic(err)
		}
		t.Run(name, func(t *testing.T) {
			repository := repository.NewBasicInformationRepository(tt.sqlHandler)

			basicInformation, err := repository.Find(ctx)
			if tt.expect != basicInformation {
				t.Errorf("expect: %+v, actual: %+v", tt.expect, basicInformation)
			}
			if (err == nil) == tt.returnsErr {
				t.Errorf("actual: %s", err)
			}
		})
		if _, err := sqlHandler.ExecContext(ctx, `UPDATE "basic_information" SET "gender" = 'MALE'`); err != nil {
			t.Fatal(err)
		}
	}
}

func TestBasicInformationRepositoryUpdate(t *testing.T) {
	const (
		SEED_BIRTHDAY            = "1991-07-01T00:00:00Z"
		SEED_GENDER              = "MALE"
		SEED_ACADEMIC_BACKGROUND = "上智大学卒業"
	)

	transactionController := repository.NewTransactionController(sqlHandler)

	tests := map[string]struct {
		sqlHandler                 repository.SqlHandler
		input                      entity.UpdateBasicInformationInput
		expectedBirthday           string
		expectedGender             string
		expectedAcademicBackground string
		isErr                      bool
	}{
		"UpdateAll": {
			sqlHandler: sqlHandler,
			input: entity.UpdateBasicInformationInput{
				Birthday:           entity.UpdateBirthdayInput{Birthday: "2022-08-24", IsUpdated: true},
				Gender:             entity.UpdateGenderInput{Gender: entity.Female, IsUpdated: true},
				AcademicBackground: entity.UpdateAcademicBackgroundInput{AcademicBackground: "荏田高等学校卒業", IsUpdated: true},
			},
			expectedBirthday:           "2022-08-24T00:00:00Z",
			expectedGender:             "FEMALE",
			expectedAcademicBackground: "荏田高等学校卒業",
		},
		"OnlyUpdateBirthday": {
			sqlHandler: sqlHandler,
			input: entity.UpdateBasicInformationInput{
				Birthday: entity.UpdateBirthdayInput{Birthday: "2022-08-24", IsUpdated: true},
			},
			expectedBirthday:           "2022-08-24T00:00:00Z",
			expectedGender:             SEED_GENDER,
			expectedAcademicBackground: SEED_ACADEMIC_BACKGROUND,
		},
		"OnlyUpdateGender": {
			sqlHandler: sqlHandler,
			input: entity.UpdateBasicInformationInput{
				Gender: entity.UpdateGenderInput{Gender: entity.Female, IsUpdated: true},
			},
			expectedBirthday:           SEED_BIRTHDAY,
			expectedGender:             "FEMALE",
			expectedAcademicBackground: SEED_ACADEMIC_BACKGROUND,
		},
		"OnlyUpdateAcademicBackground": {
			sqlHandler: sqlHandler,
			input: entity.UpdateBasicInformationInput{
				AcademicBackground: entity.UpdateAcademicBackgroundInput{AcademicBackground: "荏田高等学校卒業", IsUpdated: true},
			},
			expectedBirthday:           SEED_BIRTHDAY,
			expectedGender:             SEED_GENDER,
			expectedAcademicBackground: "荏田高等学校卒業",
		},
		"NoUpdated": {
			sqlHandler:                 sqlHandler,
			input:                      entity.UpdateBasicInformationInput{},
			expectedBirthday:           SEED_BIRTHDAY,
			expectedGender:             SEED_GENDER,
			expectedAcademicBackground: SEED_ACADEMIC_BACKGROUND,
		},
		"SqlHandlerReturnsErr": {
			sqlHandler: errSqlHandler,
			input: entity.UpdateBasicInformationInput{
				Birthday: entity.UpdateBirthdayInput{IsUpdated: true},
			},
			isErr: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			repo := repository.NewBasicInformationRepository(tt.sqlHandler)

			ctx := context.Background()
			err := transactionController.Transaction(ctx, func(tx entity.Tx) error {
				return repo.Update(ctx, tx, tt.input)
			})

			if tt.isErr {
				if err == nil {
					t.Error()
				}
				return
			}

			if err != nil {
				t.Errorf("err: %+v", err)
			}

			rows, err := sqlHandler.QueryContext(ctx, `SELECT "birthday", "gender", "academic_background" FROM "basic_information"`)
			if err != nil {
				t.Fatal(err)
			}
			defer rows.Close()

			var (
				birthday           string
				gender             string
				academicBackground string
			)
			rows.Next()
			if err := rows.Scan(&birthday, &gender, &academicBackground); err != nil {
				t.Fatal(err)
			}
			if birthday != tt.expectedBirthday {
				t.Errorf("birthday: %s, tt.expectedBirthday: %s", birthday, tt.expectedBirthday)
			}
			if gender != tt.expectedGender {
				t.Errorf("gender: %s, tt.expectedgender: %s", gender, tt.expectedGender)
			}
			if academicBackground != tt.expectedAcademicBackground {
				t.Errorf("academicBackground: %s, tt.expectedAcademicBackground: %s", academicBackground, tt.expectedAcademicBackground)
			}

			if _, err := sqlHandler.ExecContext(ctx, `UPDATE "basic_information" SET "birthday" = $1, "gender" = $2, "academic_background" = $3`, SEED_BIRTHDAY, SEED_GENDER, SEED_ACADEMIC_BACKGROUND); err != nil {
				t.Fatal(err)
			}
		})
	}
}
