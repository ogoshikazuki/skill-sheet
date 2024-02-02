package model

import (
	"github.com/ogoshikazuki/skill-sheet/entity"
	"github.com/ogoshikazuki/skill-sheet/graph/scalar"
)

type Project struct {
	ID            scalar.ID
	Name          string
	StartMonth    entity.YearMonth
	EndMonth      entity.YearMonth
	TechnologyIDs []scalar.ID
}

func (p Project) IsNode() {}

func (p Project) GetID() scalar.ID {
	return p.ID
}
