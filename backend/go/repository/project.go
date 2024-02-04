package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/ogoshikazuki/skill-sheet/entity"
)

type projectRepository struct {
	sqlHandler SqlHandler
}

func NewProjectRepository(sqlHandler SqlHandler) entity.ProjectRepository {
	return projectRepository{sqlHandler: sqlHandler}
}

func (r projectRepository) Search(ctx context.Context, projectOrders []entity.ProjectOrder) ([]entity.Project, error) {
	technologyIDsQuery := `
SELECT "project_id", "technology_id"
FROM project_technology
`
	technologyIDsRows, err := r.sqlHandler.QueryContext(ctx, technologyIDsQuery)
	if err != nil {
		return nil, err
	}
	defer technologyIDsRows.Close()

	technologyIDs := map[entity.ID]([]entity.ID){}
	for technologyIDsRows.Next() {
		var projectID, technologyID entity.ID
		if err := technologyIDsRows.Scan(&projectID, &technologyID); err != nil {
			return nil, err
		}

		technologyIDs[projectID] = append(technologyIDs[projectID], technologyID)
	}

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
		var id entity.ID
		var name string
		var startMonth time.Time
		var endMonth sql.NullTime
		if err := rows.Scan(&id, &name, &startMonth, &endMonth); err != nil {
			return []entity.Project{}, err
		}

		startYearMonth, err := entity.NewYearMonth(startMonth.Year(), int(startMonth.Month()))
		if err != nil {
			return nil, entity.NewInternalServerError(err)
		}
		var endYearMonth entity.YearMonth
		if endMonth.Valid {
			endYearMonth, err = entity.NewYearMonth(endMonth.Time.Year(), int(endMonth.Time.Month()))
			if err != nil {
				return nil, entity.NewInternalServerError(err)
			}
		}
		projects = append(projects, entity.Project{
			Id:            id,
			Name:          name,
			StartMonth:    startYearMonth,
			EndMonth:      endYearMonth,
			TechnologyIDs: technologyIDs[id],
		})
	}

	return projects, nil
}

func (r projectRepository) Find(ctx context.Context, id entity.ID) (entity.Project, error) {
	query := `
SELECT "name", "start_month", "end_month"
FROM "projects"
WHERE "id" = $1
`
	rows, err := r.sqlHandler.QueryContext(ctx, query, id)
	if err != nil {
		return entity.Project{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return entity.Project{}, nil
	}

	var name string
	var startMonth time.Time
	var endMonth sql.NullTime
	if err := rows.Scan(&name, &startMonth, &endMonth); err != nil {
		return entity.Project{}, err
	}
	if err := rows.Close(); err != nil {
		return entity.Project{}, nil
	}

	startYearMonth, err := entity.NewYearMonth(startMonth.Year(), int(startMonth.Month()))
	if err != nil {
		return entity.Project{}, entity.NewInternalServerError(err)
	}
	var endYearMonth entity.YearMonth
	if endMonth.Valid {
		endYearMonth, err = entity.NewYearMonth(endMonth.Time.Year(), int(endMonth.Time.Month()))
		if err != nil {
			return entity.Project{}, entity.NewInternalServerError(err)
		}
	}

	technologyIDsQuery := `
SELECT "technology_id"
FROM "project_technology"
WHERE "project_id" = $1
`
	technologyIDsRows, err := r.sqlHandler.QueryContext(ctx, technologyIDsQuery, id)
	if err != nil {
		return entity.Project{}, nil
	}
	defer technologyIDsRows.Close()

	technologyIDs := []entity.ID{}
	for technologyIDsRows.Next() {
		var technologyID entity.ID
		if err := technologyIDsRows.Scan(&technologyID); err != nil {
			return entity.Project{}, nil
		}

		technologyIDs = append(technologyIDs, technologyID)
	}

	return entity.Project{
		Id:            id,
		Name:          name,
		StartMonth:    startYearMonth,
		EndMonth:      endYearMonth,
		TechnologyIDs: technologyIDs,
	}, nil
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
