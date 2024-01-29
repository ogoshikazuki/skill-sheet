package scalar

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/ogoshikazuki/skill-sheet/entity"
)

func MarshalYearMonth(yearMonth entity.YearMonth) graphql.Marshaler {
	return graphql.WriterFunc((func(w io.Writer) {
		w.Write([]byte(strconv.Quote(fmt.Sprintf("%02d-%02d", yearMonth.Year(), yearMonth.Month()))))
	}))
}

func UnmarshalYearMonth(v interface{}) (entity.YearMonth, error) {
	yearMonth, ok := v.(entity.YearMonth)
	if !ok {
		return nil, errors.New("invalid year month")
	}

	return yearMonth, nil
}
