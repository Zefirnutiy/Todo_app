package main

import "github.com/golang-jwt/jwt/v5"


type User struct {
	Id 			int    `json:"id"`
	Name 		string `json:"name"`
	Email 		string `json:"email"`
	Password 	string `json:"pass"`
}

type UserLogin struct {
	Email 		string 	`json:"email"`
	Password 	string 	`json:"password"`
}

type TodoList struct {
	Id 			int		`json:"id"`
	Title 		string	`json:"title"`
	UserId 		int		`json:"userId"`
}

type Todo struct {
	Id 				int		`json:"id"` 
	Title 			string	`json:"title"`
	Description 	string	`json:"description"`
	IsReady 		bool	`json:"isReady"`
	ListId 			int		`json:"listId"`
}

type TodoUpdate struct {
	Id 				*int	`json:"id"` 
	Title 			*string	`json:"title"`
	Description 	*string	`json:"description"`
	IsReady 		*bool	`json:"isReady"`
	ListId 			*int	`json:"listId"`
}

type ID struct {
	ID			int `json:"id"`
}

type MyCustomClaim struct {
	UserId int `json:"userId"`
	jwt.RegisteredClaims
}