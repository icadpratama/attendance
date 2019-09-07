package resolvers

import (
	"github.com/icadpratama/attendance/internal/gql"
	"github.com/icadpratama/attendance/internal/orm"
)

//go:generate go run github.com/99designs/gqlgen

type Resolver struct{
	ORM *orm.ORM
}

func (r *Resolver) Mutation() gql.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
