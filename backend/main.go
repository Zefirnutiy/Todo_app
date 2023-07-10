package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"strings"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

	if err := PostgresConnect(); err != nil {
		panic(err.Error())
	}

	router := gin.Default()

	user := router.Group("/user")
	{
		user.GET("/", getAllUsers)
		user.DELETE("/delete", deleteUser)

	}

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", signUp)
		auth.POST("/sign-in", signIn)
	}

	todo := router.Group("/todo")
	{
		todo.GET("/:id", CheckToken, getTodoByTodoId)
	 	todo.GET("/list/:id", CheckToken, getTodosByListId)
		todo.POST("/create", CheckToken, createTodo)
		todo.DELETE("/delete/:id", CheckToken, deleteTodo)
		todo.PUT("/update", CheckToken, updateTodo)

	}

	list := router.Group("/list")
	{
		list.GET("/", CheckToken, getAllLists)
		list.POST("/create", CheckToken, createList)
		list.DELETE("/delete/:id", CheckToken, deleteList)
		list.PUT("/update", CheckToken, updateList)

	}

	router.Run(":8080")
}

func getAllUsers(ctx *gin.Context) {
	users, err := GetAllUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})

	return
}

func deleteUser(ctx *gin.Context) {
	var userId ID

	if err := ctx.BindJSON(&userId); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}


	if err := DeleteUser(userId.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":"user has been deleted",
	})

	return
}

func signUp(ctx *gin.Context){
	var user User
	
	if err := ctx.BindJSON(&user); err != nil{ 
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	if err := CreateUser(user); err != nil {
		print("\n", err.Error())
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, user)
	return
}

func signIn(ctx *gin.Context){
	var userLogin UserLogin
	
	if err := ctx.BindJSON(&userLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	id, err := IsValidUserLogin(userLogin)
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	var token string
	token, err = GenerateToken(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":"you're has been logined",
		"token":token,
	})
	return
}


func getTodoByTodoId( ctx *gin.Context){
	todoId := ctx.Param("id")


	todo, err := GetTodoById(StringToInt(todoId))
	
	if err != nil {
		print("\n", err.Error())
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"todo": todo,
	})
	
	return
}

func getTodosByListId( ctx *gin.Context){
	todoListId := ctx.Param("id")


	data, err := GetTodosByListId(StringToInt(todoListId))
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"data":data,
	})

	return
}

func createTodo(ctx *gin.Context){
	var todo Todo
	if err := ctx.BindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	if err := CreateTodo(todo); err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	} 
		
	ctx.JSON(http.StatusOK, gin.H{
		"message":"todo has been created",
	})
		
	return
}
	
func deleteTodo(ctx *gin.Context){
	todoId := ctx.Param("id")

	if err := DeleteTodo(StringToInt(todoId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"message":"todo has been deleted",
	})
	
	return
}
	
func updateTodo(ctx *gin.Context){
	var todo TodoUpdate
	if err := ctx.BindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	if err := UpdateTodo(todo); err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":"ok",
	})

	return
}
	

func createList(ctx *gin.Context){
	var list TodoList
	
	if err := ctx.BindJSON(&list); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := CreateTodoList(list); err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"message": "todo list has been created",
	})
	return
}

func updateList(ctx *gin.Context){
	var list TodoList
	
	if err := ctx.BindJSON(&list); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	if err := UpdateList(list); err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":"ok",
	})
	return
}

func deleteList(ctx *gin.Context){
	listId := ctx.Param("id")
	
	if err := DeleteList(StringToInt(listId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
		} 
		
		ctx.JSON(http.StatusOK, gin.H{
			"message": "list has been deleted",
		})
		
		return
}
	
func getAllLists(ctx *gin.Context){
	
	
	authHeader := ctx.Request.Header["Authorization"][0]
	token := strings.Split(authHeader, " ")[1]

	userId, err := ParseToken(token)
	if err != nil  || userId == nil{
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	todoLists, err := GetAllLists(*userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": todoLists,
	})
	return
}
