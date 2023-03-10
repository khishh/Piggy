package graph

import "github.com/khishh/personal-finance-app/pkg/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	BookRepository repository.BookRepository
	UserRepository repository.UserRepository
	ItemRepository repository.ItemRepository
}
