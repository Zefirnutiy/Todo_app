package main

import (
	"context"
	"errors"
	"os"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

func PostgresConnect() error {
	var err error

	conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	return err
}

func GetAllUsers() ([]User, error){
	rows, err := conn.Query(context.Background(), "SELECT * FROM public.user")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var user User
	var usersArray []User

	for rows.Next(){
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}

		usersArray = append(usersArray, user)
	}

	return usersArray, nil
}

func isUserExist(email string) bool {
	var em string
	err := conn.QueryRow(context.Background(), "SELECT email FROM public.user WHERE email=$1", email).Scan(&em)

	
	if err != nil {
		return false
	}
	
	return true
}

// мне не нравится такая архитектура, поэтому перепродумать ее нужно
// добавить создание дефолтного спика для добавления туда задач
func CreateUser(user User) error {

	if isUserExist(user.Email) {
		return errors.New("that user is exist")
	}

	var id int
	err := conn.QueryRow(context.Background(), "INSERT INTO public.user(name, email, password) VALUES($1, $2, $3) RETURNING id", user.Name, user.Email, user.Password).Scan(&id)

	return err
}

func DeleteUser(id int) error {
	_, err := conn.Exec(context.Background(), "DELETE FROM public.user WHERE id=$1", id)

	return err
}

func IsValidUserLogin(user UserLogin) error {

	_, err := conn.Exec(context.Background(), "SELECT id FROM public.user WHERE email=$1 AND password=$2", user.Email, user.Password)

	return err
}



func GetTodoById(id int) (Todo, error) {
	var todo Todo

	err := conn.QueryRow(context.Background(), "SELECT * FROM public.todo").Scan(&todo.Id, &todo.IsReady, &todo.ListId, &todo.Description, &todo.Title)

	return todo, err
}

func CreateTodo(todo Todo) error {
	_, err := conn.Exec(context.Background(), "INSERT INTO public.todo (title, description, is_ready, list_id) VALUES($1, $2, $3, $4)",
	todo.Title, todo.Description, todo.IsReady, todo.ListId)
	
	return err
}

func GetTodosByListId(listId int) ([]Todo, error){
	rows, err := conn.Query(context.Background(), "SELECT * FROM public.todo WHERE list_id=$1", listId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todo Todo
	var todoArray []Todo

	for rows.Next(){
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.IsReady, &todo.ListId)

		if err != nil {
			return nil, err
		}

		todoArray = append(todoArray, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return todoArray, nil
}

func DeleteTodo(todoId int) error {
	_, err := conn.Exec(context.Background(), "DELETE FROM public.todo WHERE id=$1", todoId)
	
	return err
}

func UpdateTodo(todo TodoUpdate) error {
	_, err := conn.Exec(context.Background(), GenerateSql(todo), &todo.Id)

	return err
}



func GetAllLists(userId int) ([]TodoList, error) {

	rows, err := conn.Query(context.Background(), "SELECT * from public.todo_list WHERE user_id=$1", userId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todoList TodoList
	var todoListMas = make([]TodoList, 0)

	for rows.Next() {
		err = rows.Scan(&todoList.Id, &todoList.Title, &todoList.UserId)

		if err != nil {
			return nil, err
		}

		todoListMas = append(todoListMas, todoList)
	}

	return todoListMas, nil
}

func CreateTodoList(todoList TodoList) error {
	_, err := conn.Exec(context.Background(), "INSERT INTO public.todo_list(title, user_id) VALUES($1, $2)", todoList.Title, todoList.UserId)

	return err
}

func DeleteList(listId int) error {
	_, err := conn.Exec(context.Background(), "DELETE FROM public.todo_list WHERE id=$1", listId)

	return err
}

func UpdateList(todoList TodoList) error {
	_, err := conn.Exec(context.Background(), "UPDATE public.todo_list SET title=$1 WHERE id=$2", todoList.Title, todoList.Id)
	
	return err
}