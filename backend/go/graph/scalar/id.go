package scalar

import (
	"encoding/base64"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/ogoshikazuki/skill-sheet/entity"
)

const ID_SEPARATOR = ":"

type ID struct {
	typename string
	id       entity.ID
}

func NewID(typename string, id entity.ID) ID {
	return ID{
		typename: typename,
		id:       id,
	}
}

func (i ID) GetTypename() string {
	return i.typename
}

func (i ID) GetID() entity.ID {
	return i.id
}

func (i *ID) UnmarshalGQL(v interface{}) error {
	input, ok := v.(string)
	if !ok {
		return fmt.Errorf("ID must be a string")
	}

	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return err
	}

	splited := strings.Split(string(decoded), ID_SEPARATOR)
	if len(splited) != 2 {
		return fmt.Errorf("id must be in the following format: [typename%sid]", ID_SEPARATOR)
	}

	i.typename = splited[0]
	id, err := strconv.Atoi(splited[1])
	if err != nil {
		return err
	}
	i.id = entity.ID(id)

	return nil
}

func (i ID) MarshalGQL(w io.Writer) {
	beforeEncode := fmt.Sprintf("%s%s%d", i.typename, ID_SEPARATOR, i.id)
	encoded := base64.StdEncoding.EncodeToString([]byte(beforeEncode))

	w.Write([]byte(strconv.Quote(encoded)))
}
