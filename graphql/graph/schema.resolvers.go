package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"gogingraphqleg/graph/generated"
	"gogingraphqleg/graph/model"
	"log"

	"github.com/google/uuid"
)

var users []*model.User

func init() {
	log.Println("Init - Users array to be created")
	users = make([]*model.User, 0)
	users = append(users, &model.User{ID: "1", FirstName: "Srinivasa", LastName: "Ramanujam", Dob: "12/27/1887"})
	users = append(users, &model.User{ID: "2", FirstName: "CV", LastName: "Raman", Dob: "11/7/1888"})
	users = append(users, &model.User{ID: "3", FirstName: "Subrahmanyan", LastName: "Chandrasekhar", Dob: "10/19/1910"})
	log.Println("Init - Users array has been created")
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	uuidValue := uuid.NewString()
	user := &model.User{
		ID:        uuidValue,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Dob:       input.Dob,
	}
	users = append(users, user)
	return user, nil
}

// RemoveUser is the resolver for the removeUser field.
func (r *mutationResolver) RemoveUser(ctx context.Context, input model.DeleteUser) (*model.User, error) {
	index := -1
	for i, user := range users {
		if user.ID == input.UserID {
			index = i
			break
		}
	}
	if index == -1 {
		return nil, errors.New("cannot find user you are looking for")
	}
	user := users[index]
	users = append(users[:index], users[index+1:]...)

	return user, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return users, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	index := -1
	for i, user := range users {
		if user.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		return nil, errors.New("cannot find user you are looking for")
	}
	user := users[index]

	return user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
