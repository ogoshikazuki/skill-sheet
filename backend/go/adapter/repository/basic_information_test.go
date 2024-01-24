package repository_test

import (
	"context"
	"testing"

	"github.com/ogoshikazuki/skill-sheet/adapter/repository"
	"github.com/ogoshikazuki/skill-sheet/entity"
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
	}
}
