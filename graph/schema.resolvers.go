package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/testgql/gqlgen-todos/graph/generated"
	"github.com/testgql/gqlgen-todos/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", 10),
		UserID: input.UserID,
		// User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		ID:   input.ID,
		Name: input.Name,
		// User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.users = append(r.users, user)
	return user, nil
}

func (r *queryResolver) Todos(ctx context.Context, id *string, text *string, first *int, after *int, order *string) ([]*model.Todo, error) {
	if first != nil && after != nil && *after >= 0 && *first >= 0 {
		return r.todos[*after:*first], nil
	} else {
		return r.todos, nil
	}

}

func (r *queryResolver) Users(ctx context.Context, id *string, name *string, first *int, after *int, order *string) ([]*model.User, error) {
	if first != nil && after != nil && *after >= 0 && *first >= 0 {
		return r.users[*after:*first], nil
	} else {
		return r.users, nil
	}
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) User(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}
