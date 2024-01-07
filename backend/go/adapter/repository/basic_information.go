package repository

import (
	"context"
	"time"

	"github.com/ogoshikazuki/skill-sheet/entity"
)

type basicInformationRepository struct {
	sqlHandler Sqlhandler
}

func (repository basicInformationRepository) Find(ctx context.Context) (entity.BasicInformation, error) {
	rows, err := repository.sqlHandler.QueryContext(ctx, `
SELECT birthday
FROM basic_information
LIMIT 1
`)
	if err != nil {
		return entity.BasicInformation{}, err
	}
	defer rows.Close()

	rows.Next()

	var birthday time.Time
	if err := rows.Scan(&birthday); err != nil {
		return entity.BasicInformation{}, err
	}

	return entity.BasicInformation{Birthday: entity.NewDateFromTime(birthday)}, nil
}

func NewBasicInformationRepository(sqlHandler Sqlhandler) entity.BasicInformationRepository {
	return basicInformationRepository{sqlHandler: sqlHandler}
}
