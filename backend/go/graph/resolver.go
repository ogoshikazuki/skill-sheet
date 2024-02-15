package graph

import "github.com/ogoshikazuki/skill-sheet/usecase"

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	findBasicInformationUsecase   usecase.FindBasicInformationUsecase
	findProjectUsecase            usecase.FindProjectUsecase
	searchProjectsUsecase         usecase.SearchProjectsUsecase
	updateBasicInformationUsecase usecase.UpdateBasicInformationUsecase
}

func NewResolverRoot(
	findBasicInformationUsecase usecase.FindBasicInformationUsecase,
	findProjectUsecase usecase.FindProjectUsecase,
	searchProjectsUsecase usecase.SearchProjectsUsecase,
	updateBasicInformationUsecase usecase.UpdateBasicInformationUsecase,
) ResolverRoot {
	return &Resolver{
		findBasicInformationUsecase:   findBasicInformationUsecase,
		findProjectUsecase:            findProjectUsecase,
		searchProjectsUsecase:         searchProjectsUsecase,
		updateBasicInformationUsecase: updateBasicInformationUsecase,
	}
}
