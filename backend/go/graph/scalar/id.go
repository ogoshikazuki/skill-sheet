package scalar

import (
	"encoding/base64"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const ID_SEPARATOR = ":"

type ID struct {
	typename string
	id       uint
}

func NewID(typename string, id uint) ID {
	return ID{
		typename: typename,
		id:       id,
	}
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
	i.id = uint(id)

	return nil
}

func (i ID) MarshalGQL(w io.Writer) {
	beforeEncode := fmt.Sprintf("%s%s%d", i.typename, ID_SEPARATOR, i.id)
	encoded := base64.StdEncoding.EncodeToString([]byte(beforeEncode))

	w.Write([]byte(strconv.Quote(encoded)))
}
