package middleware

import (
	"context"

	"net/http"

	"github.com/cockroachdb/errors"
	"github.com/ogoshikazuki/skill-sheet/entity"
	"github.com/ogoshikazuki/skill-sheet/repository"
	"github.com/vikstrous/dataloadgen"
)

func Dataloader(sqlHandler repository.SqlHandler) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			loaders := newLoaders(sqlHandler)
			r = r.WithContext(context.WithValue(r.Context(), loadersKey, loaders))
			next.ServeHTTP(w, r)
		})
	}
}

type loaders struct {
	technologyLoader *dataloadgen.Loader[entity.ID, entity.Technology]
}

func newLoaders(sqlHandler repository.SqlHandler) *loaders {
	return &loaders{
		technologyLoader: dataloadgen.NewLoader(repository.NewTechnologyRepository(sqlHandler).SearchByIDsWithMultipleErr),
	}
}

func getLoadersFromContext(ctx context.Context) *loaders {
	return ctx.Value(loadersKey).(*loaders)
}

func GetTechnologyFromDataloader(ctx context.Context, id entity.ID) (entity.Technology, error) {
	technology, err := getLoadersFromContext(ctx).technologyLoader.Load(ctx, id)
	if err != nil {
		if err, ok := err.(dataloadgen.ErrorSlice); ok {
			return technology, errorSlice(err)
		}
		return technology, entity.ErrInternalAndLogStack(ctx, err)
	}
	return technology, nil
}

type errorSlice dataloadgen.ErrorSlice

func (e errorSlice) Error() string {
	combinedErr := errors.Join([]error(e)...)
	if combinedErr == nil {
		return "no error data"
	}
	return combinedErr.Error()
}
func (e errorSlice) Unwrap() []error {
	return e
}

func GetTechnologiesFromDataloader(ctx context.Context, ids []entity.ID) ([]entity.Technology, error) {
	technologies := []entity.Technology{}
	for _, id := range ids {
		technology, err := GetTechnologyFromDataloader(ctx, id)
		if err != nil {
			return nil, err
		}

		technologies = append(technologies, technology)
	}

	return technologies, nil
}
