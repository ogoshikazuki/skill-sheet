package entity

import "context"

type Technology struct {
	ID   ID
	Name string
}

type TechnologyRepository interface {
	SearchByIDsWithMultipleErr(context.Context, []ID) ([]Technology, []error)
}
