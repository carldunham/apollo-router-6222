package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"main/graph/model"
)

// FindAnimalByID is the resolver for the findAnimalByID field.
func (r *entityResolver) FindAnimalByID(ctx context.Context, id string) (*model.Animal, error) {
	var hello string
	switch id {
	case "1":
		hello = "woof"
	case "2":
		hello = "meow"
	}
	return &model.Animal{
		ID:    id,
		Hello: hello,
	}, nil
}

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }