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

func setUpdateInput[T any](input map[string]interface{}, field string, f func(v T)) {
	if v, ok := input[field]; ok {
		p := (v.(*T))
		if p == nil {
			return
		}
		f(*p)
	}
}
