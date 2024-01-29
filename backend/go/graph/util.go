package graph

import (
	"github.com/ogoshikazuki/skill-sheet/entity"
	"github.com/ogoshikazuki/skill-sheet/graph/model"
)

func convertDirectionFromGraphToEntity(direction model.OrderDirection) entity.OrderDirection {
	switch direction {
	case model.OrderDirectionAsc:
		return entity.ASC
	case model.OrderDirectionDesc:
		return entity.DESC
	}
	return entity.ASC
}
