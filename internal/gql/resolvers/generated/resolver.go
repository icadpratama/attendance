package resolvers

import (
	"context"

	"github.com/icadpratama/attendance/internal/gql"
	"github.com/icadpratama/attendance/internal/gql/models"
)

type Resolver struct{}

func (r *Resolver) Mutation() gql.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (bool, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context, userID *string) ([]*models.User, error) {
	panic("not implemented")
}
