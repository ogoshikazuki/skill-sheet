package entity

//go:generate go run github.com/matryer/moq -out project_mock.go . ProjectRepository

import "context"

type Project struct {
	Id         ID
	Name       string
	StartMonth YearMonth
	// EndMonth 現在進行形のプロジェクトの場合null相当となる
	EndMonth      YearMonth
	TechnologyIDs []ID
}

type ProjectOrderField int

const (
	PROJECT_ORDER_START_MONTH ProjectOrderField = iota
	PROJECT_ORDER_END_MONTH
)

type ProjectOrder struct {
	Field     ProjectOrderField
	Direction OrderDirection
}

type ProjectRepository interface {
	Search(context.Context, []ProjectOrder) ([]Project, error)
	Find(context.Context, ID) (Project, error)
}
