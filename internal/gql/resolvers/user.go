package resolvers

import (
	"context"

	log "github.com/icadpratama/attendance/internal/logger"

	"github.com/icadpratama/attendance/internal/gql/models"
	tf "github.com/icadpratama/attendance/internal/gql/resolvers/transformations"
	dbm "github.com/icadpratama/attendance/internal/orm/models"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	return userCreateUpdate(r, input, false)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input models.UserInput) (*models.User, error) {
	return userCreateUpdate(r, input, true, id)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	return userDelete(r, id)
}

func (r *queryResolver) Users(ctx context.Context, id *string) (*models.Users, error) {
	return userList(r, id)
}

func userCreateUpdate(r *mutationResolver, input models.UserInput, update bool, ids ...string) (*models.User, error) {
	dbo, err := tf.GQLInputUserToDBUser(&input, update, ids...)
	if err != nil {
		return nil, err
	}

	db := r.ORM.DB.New().Begin()
	if !update {
		db = db.Create(dbo).First(dbo)
	} else {
		db = db.Model(&dbo).Update(dbo).First(dbo)
	}
	gql, err := tf.DBUserToGQLUser(dbo)
	if err != nil {
		db.RollbackUnlessCommitted()
		return nil, err
	}
	db = db.Commit()
	return gql, db.Error
}

func userDelete(r *mutationResolver, id string) (bool, error) {
	return false, nil
}

func userList(r *queryResolver, id *string) (*models.Users, error) {
	entity := "users"
	whereID := "id = ?"
	record := &models.Users{}
	dbRecords := []*dbm.User{}
	db := r.ORM.DB.New()
	if id != nil {
		db = db.Where(whereID, *id)
	}
	db = db.Find(&dbRecords).Count(&record.Count)
	for _, dbRec := range dbRecords {
		if rec, err := tf.DBUserToGQLUser(dbRec); err != nil {
			log.Errorfn(entity, err)
		} else {
			record.List = append(record.List, rec)
		}
	}
	return record, db.Error
}