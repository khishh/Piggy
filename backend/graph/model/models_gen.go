// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Book struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

type BookInput struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

type User struct {
	ID        int     `json:"id"`
	Email     string  `json:"email"`
	LastName  string  `json:"last_name"`
	FirstName string  `json:"first_name"`
	Picture   *string `json:"picture"`
	Sub       string  `json:"sub"`
}

type UserInput struct {
	Email     string  `json:"email"`
	LastName  string  `json:"last_name"`
	FirstName string  `json:"first_name"`
	Picture   *string `json:"picture"`
	Sub       string  `json:"sub"`
}
