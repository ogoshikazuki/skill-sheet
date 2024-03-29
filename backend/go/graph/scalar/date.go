package scalar

import (
	"errors"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/ogoshikazuki/skill-sheet/entity"
)

func MarshalDate(d entity.Date) graphql.Marshaler {
	return graphql.WriterFunc((func(w io.Writer) {
		w.Write([]byte(strconv.Quote(string(d))))
	}))
}

func UnmarshalDate(v interface{}) (entity.Date, error) {
	s, ok := v.(string)
	if !ok {
		return "", errors.New("date must be a string")
	}

	date := entity.Date(s)

	if !date.IsValid() {
		return "", errors.New("date is invalid")
	}

	return date, nil
}
