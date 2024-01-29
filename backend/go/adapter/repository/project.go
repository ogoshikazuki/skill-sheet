package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/ogoshikazuki/skill-sheet/entity"
)

type projectRepository struct {
	sqlHandler SqlHandler
}

func NewProjectRepository(sqlHandler SqlHandler) entity.ProjectRepository {
	return projectRepository{sqlHandler: sqlHandler}
}

func (r projectRepository) Search(ctx context.Context, projectOrders []entity.ProjectOrder) ([]entity.Project, error) {
	query := `
SELECT "id", "name", "start_month", "end_month"
FROM "projects"
`
	query = r.addOrderBy(query, projectOrders)

	rows, err := r.sqlHandler.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projects := []entity.Project{}
	for rows.Next() {
		var id uint
		var name string
		var startMonth time.Time
		var endMonth sql.NullTime
		if err := rows.Scan(&id, &name, &startMonth, &endMonth); err != nil {
			return []entity.Project{}, err
		}

		startYearMonth, err := entity.NewYearMonth(startMonth.Year(), int(startMonth.Month()))
		if err != nil {
			return nil, errors.WithStack(err)
		}
		var endYearMonth entity.YearMonth
		if endMonth.Valid {
			endYearMonth, err = entity.NewYearMonth(endMonth.Time.Year(), int(endMonth.Time.Month()))
			if err != nil {
				return nil, errors.WithStack(err)
			}
		}
		projects = append(projects, entity.Project{
			Id:         id,
			Name:       name,
			StartMonth: startYearMonth,
			EndMonth:   endYearMonth,
		})
	}

	return projects, nil
}

func (r projectRepository) addOrderBy(query string, projectOrders []entity.ProjectOrder) string {
	for i, projectOrder := range projectOrders {
		if i == 0 {
			query += ` ORDER BY`
		} else {
			query += `,`
		}
		switch projectOrder.Field {
		case entity.PROJECT_ORDER_START_MONTH:
			query += ` "start_month"`
		case entity.PROJECT_ORDER_END_MONTH:
			query += ` "end_month"`
		}
		switch projectOrder.Direction {
		case entity.ASC:
			query += ` ASC`
		case entity.DESC:
			query += ` DESC`
		}
	}
	return query
}
