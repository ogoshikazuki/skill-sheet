// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/ogoshikazuki/skill-sheet/entity"
	"github.com/ogoshikazuki/skill-sheet/graph/scalar"
)

type Node interface {
	IsNode()
	GetID() scalar.ID
}

type BasicInformation struct {
	ID                 scalar.ID   `json:"id"`
	AcademicBackground string      `json:"academicBackground"`
	Birthday           entity.Date `json:"birthday"`
	Gender             Gender      `json:"gender"`
}

type Mutation struct {
}

type ProjectOrder struct {
	Field     ProjectOrderField `json:"field"`
	Direction OrderDirection    `json:"direction"`
}

type Query struct {
}

type Technology struct {
	ID   scalar.ID `json:"id"`
	Name string    `json:"name"`
}

type UpdateBasicInformationPayload struct {
	BasicInformation *BasicInformation `json:"basicInformation"`
}

type Gender string

const (
	GenderFemale Gender = "FEMALE"
	GenderMale   Gender = "MALE"
)

var AllGender = []Gender{
	GenderFemale,
	GenderMale,
}

func (e Gender) IsValid() bool {
	switch e {
	case GenderFemale, GenderMale:
		return true
	}
	return false
}

func (e Gender) String() string {
	return string(e)
}

func (e *Gender) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Gender(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Gender", str)
	}
	return nil
}

func (e Gender) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OrderDirection string

const (
	OrderDirectionAsc  OrderDirection = "ASC"
	OrderDirectionDesc OrderDirection = "DESC"
)

var AllOrderDirection = []OrderDirection{
	OrderDirectionAsc,
	OrderDirectionDesc,
}

func (e OrderDirection) IsValid() bool {
	switch e {
	case OrderDirectionAsc, OrderDirectionDesc:
		return true
	}
	return false
}

func (e OrderDirection) String() string {
	return string(e)
}

func (e *OrderDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderDirection", str)
	}
	return nil
}

func (e OrderDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ProjectOrderField string

const (
	ProjectOrderFieldStartMonth ProjectOrderField = "START_MONTH"
	ProjectOrderFieldEndMonth   ProjectOrderField = "END_MONTH"
)

var AllProjectOrderField = []ProjectOrderField{
	ProjectOrderFieldStartMonth,
	ProjectOrderFieldEndMonth,
}

func (e ProjectOrderField) IsValid() bool {
	switch e {
	case ProjectOrderFieldStartMonth, ProjectOrderFieldEndMonth:
		return true
	}
	return false
}

func (e ProjectOrderField) String() string {
	return string(e)
}

func (e *ProjectOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProjectOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProjectOrderField", str)
	}
	return nil
}

func (e ProjectOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
