package resolvers

import (
	"context"

	"github.com/icadpratama/attendance/internal/gql/models"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (bool, error) {
	panic("not implemented")
}

func (r *queryResolver) Users(ctx context.Context, userID *string) ([]*models.User, error) {
	records := []*models.User{
		&models.User{
			ID:     "ec17af15-e354-440c-a09f-69715fc8b595",
			Email:  "your@email.com",
			UserID: "UserID-1",
		},
	}
	return records, nil
}
