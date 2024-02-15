package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/ogoshikazuki/skill-sheet/entity"
)

type basicInformationRepository struct {
	sqlHandler SqlHandler
}

func (repository basicInformationRepository) Find(ctx context.Context) (entity.BasicInformation, error) {
	rows, err := repository.sqlHandler.QueryContext(ctx, `
SELECT "birthday", "gender", "academic_background"
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
	var academicBackground string
	if err := rows.Scan(&birthday, &gender, &academicBackground); err != nil {
		return entity.BasicInformation{}, err
	}

	entityGender, err := convertGenderFromDbToEntity(gender)
	if err != nil {
		return entity.BasicInformation{}, err
	}

	return entity.BasicInformation{
		Birthday:           entity.NewDateFromTime(birthday),
		Gender:             entityGender,
		AcademicBackground: academicBackground,
	}, nil
}

func (r basicInformationRepository) Update(ctx context.Context, tx entity.Tx, input entity.UpdateBasicInformationInput) error {
	if !input.Birthday.IsUpdated && !input.Gender.IsUpdated && !input.AcademicBackground.IsUpdated {
		return nil
	}

	t := tx.(Tx)

	query := `UPDATE "basic_information" SET`
	argCount := 0
	var args []any

	if input.Birthday.IsUpdated {
		argCount++
		query += setUpdate(argCount, "birthday")
		args = append(args, input.Birthday.Birthday)
	}
	if input.Gender.IsUpdated {
		argCount++
		query += setUpdate(argCount, "gender")
		gender, err := convertGenderFromEntityToDB(input.Gender.Gender)
		if err != nil {
			return err
		}
		args = append(args, gender)
	}
	if input.AcademicBackground.IsUpdated {
		argCount++
		query += setUpdate(argCount, "academic_background")
		args = append(args, input.AcademicBackground.AcademicBackground)
	}

	if _, err := t.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
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

func convertGenderFromEntityToDB(gender entity.Gender) (string, error) {
	switch gender {
	case entity.Male:
		return "MALE", nil
	case entity.Female:
		return "FEMALE", nil
	default:
		return "", entity.NewInternalServerError(fmt.Errorf("unexpected gender: %d", gender))
	}
}
