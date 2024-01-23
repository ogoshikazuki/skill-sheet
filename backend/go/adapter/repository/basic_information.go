package repository

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/ogoshikazuki/skill-sheet/entity"
)

type basicInformationRepository struct {
	sqlHandler SqlHandler
}

func (repository basicInformationRepository) Find(ctx context.Context) (entity.BasicInformation, error) {
	rows, err := repository.sqlHandler.QueryContext(ctx, `
SELECT "birthday", "gender"
FROM "basic_information"
LIMIT 1
`)
	if err != nil {
		return entity.BasicInformation{}, err
	}
	defer rows.Close()

	rows.Next()

	var birthday time.Time
	var gender string
	if err := rows.Scan(&birthday, &gender); err != nil {
		return entity.BasicInformation{}, err
	}

	entityGender, err := convertGenderFromDbToEntity(gender)
	if err != nil {
		return entity.BasicInformation{}, err
	}

	return entity.BasicInformation{
		Birthday: entity.NewDateFromTime(birthday),
		Gender:   entityGender,
	}, nil
}

func NewBasicInformationRepository(sqlHandler SqlHandler) entity.BasicInformationRepository {
	return basicInformationRepository{sqlHandler: sqlHandler}
}

func convertGenderFromDbToEntity(gender string) (entity.Gender, error) {
	switch gender {
	case "MALE":
		return entity.Male, nil
	case "FEMALE":
		return entity.Female, nil
	default:
		return 0, errors.Newf("gender is invalid. got %s.", gender)
	}
}
