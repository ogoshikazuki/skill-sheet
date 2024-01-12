package repository_test

import (
	"context"
	"testing"

	"github.com/ogoshikazuki/skill-sheet/adapter/repository"
	"github.com/ogoshikazuki/skill-sheet/entity"
)

func TestBasicInformationRepositoryFind(t *testing.T) {
	tests := map[string]struct {
		sqlHandler repository.SqlHandler
		expect     entity.BasicInformation
		returnsErr bool
	}{
		"Success": {
			sqlHandler: sqlHandler,
			expect: entity.BasicInformation{
				Birthday: "1991-07-01",
			},
			returnsErr: false,
		},
		"SqlHandlerReturnsErr": {
			sqlHandler: errSqlHandler,
			expect:     entity.BasicInformation{},
			returnsErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			repository := repository.NewBasicInformationRepository(tt.sqlHandler)
			ctx := context.Background()

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
