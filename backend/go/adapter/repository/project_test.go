package repository_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/ogoshikazuki/skill-sheet/adapter/repository"
	"github.com/ogoshikazuki/skill-sheet/entity"
)

func TestProjectSearch(t *testing.T) {
	yearMonth201704, _ := entity.NewYearMonth(2017, 4)
	yearMonth201808, _ := entity.NewYearMonth(2018, 8)
	yearMonth202007, _ := entity.NewYearMonth(2020, 7)
	yearMonth202103, _ := entity.NewYearMonth(2021, 3)
	yearMonth202110, _ := entity.NewYearMonth(2021, 10)

	tests := map[string]struct {
		projectOrders    []entity.ProjectOrder
		sqlHandler       repository.SqlHandler
		expectedProjects []entity.Project
		returnsErr       bool
	}{
		"Normal": {
			projectOrders: []entity.ProjectOrder{
				{
					Field:     entity.PROJECT_ORDER_START_MONTH,
					Direction: entity.DESC,
				},
				{
					Field:     entity.PROJECT_ORDER_END_MONTH,
					Direction: entity.DESC,
				},
			},
			sqlHandler: sqlHandler,
			expectedProjects: []entity.Project{
				{
					Id:         3,
					Name:       "健診PHR開発プロジェクト",
					StartMonth: yearMonth202110,
				},
				{
					Id:         2,
					Name:       "オンライン商談システムの管理画面保守開発",
					StartMonth: yearMonth202007,
					EndMonth:   yearMonth202103,
				},
				{
					Id:         1,
					Name:       "人材紹介会社向けクラウド型業務管理システムのリニューアル",
					StartMonth: yearMonth201704,
					EndMonth:   yearMonth201808,
				},
			},
		},
		"ASC": {
			projectOrders: []entity.ProjectOrder{
				{
					Field:     entity.PROJECT_ORDER_START_MONTH,
					Direction: entity.ASC,
				},
				{
					Field:     entity.PROJECT_ORDER_END_MONTH,
					Direction: entity.ASC,
				},
			},
			sqlHandler: sqlHandler,
			expectedProjects: []entity.Project{
				{
					Id:         1,
					Name:       "人材紹介会社向けクラウド型業務管理システムのリニューアル",
					StartMonth: yearMonth201704,
					EndMonth:   yearMonth201808,
				},
				{
					Id:         2,
					Name:       "オンライン商談システムの管理画面保守開発",
					StartMonth: yearMonth202007,
					EndMonth:   yearMonth202103,
				},
				{
					Id:         3,
					Name:       "健診PHR開発プロジェクト",
					StartMonth: yearMonth202110,
				},
			},
		},
		"SqlHandlerReturnsErr": {
			sqlHandler: errSqlHandler,
			returnsErr: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()

			repo := repository.NewProjectRepository(tt.sqlHandler)

			projects, err := repo.Search(ctx, tt.projectOrders)

			if !reflect.DeepEqual(projects, tt.expectedProjects) {
				t.Errorf("tt.expectedProjects: %+v, projects: %+v", tt.expectedProjects, projects)
			}
			if (err == nil) == tt.returnsErr {
				t.Errorf("tt.returnsErr: %t, err: %s", tt.returnsErr, err)
			}
		})
	}
}
