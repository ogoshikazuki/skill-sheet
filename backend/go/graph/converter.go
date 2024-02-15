package graph

import (
	"github.com/ogoshikazuki/skill-sheet/entity"
	"github.com/ogoshikazuki/skill-sheet/graph/model"
	"github.com/ogoshikazuki/skill-sheet/graph/scalar"
)

func convertBasicInformationFromEntityToGraph(basicInformation entity.BasicInformation) *model.BasicInformation {
	var gender model.Gender
	switch basicInformation.Gender {
	case entity.Male:
		gender = model.GenderMale
	case entity.Female:
		gender = model.GenderFemale
	}
	return &model.BasicInformation{
		ID:                 scalar.NewID("BasicInformation", 0),
		Birthday:           basicInformation.Birthday,
		Gender:             gender,
		AcademicBackground: basicInformation.AcademicBackground,
	}
}
