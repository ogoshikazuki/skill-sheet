package repository

import (
	"context"

	"github.com/ogoshikazuki/skill-sheet/entity"
)

type technologyRepository struct {
	sqlHandler SqlHandler
}

func NewTechnologyRepository(sqlHandler SqlHandler) entity.TechnologyRepository {
	return &technologyRepository{
		sqlHandler: sqlHandler,
	}
}

func (t *technologyRepository) SearchByIDsWithMultipleErr(ctx context.Context, ids []entity.ID) ([]entity.Technology, []error) {
	query := `
SELECT "id", "name"
FROM "technologies"
WHERE "id" IN (` + makePlaceholders(len(ids)) + `)
`
	args := []interface{}{}
	for _, id := range ids {
		args = append(args, id)
	}
	rows, err := t.sqlHandler.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, []error{err}
	}
	defer rows.Close()

	technologies := make([]entity.Technology, len(ids))
	errs := make([]error, len(ids))

	indexes := map[entity.ID]int{}
	for i, id := range ids {
		indexes[id] = i
	}

	for rows.Next() {
		var technology entity.Technology
		err := rows.Scan(&technology.ID, &technology.Name)
		technologies[indexes[technology.ID]] = technology
		errs[indexes[technology.ID]] = err
	}

	return technologies, errs
}
