package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

	err := PostgresConnect()

	if err != nil {
		panic(err.Error())
	}

	router := gin.Default()

	user := router.Group("/user")
	{
		user.GET("/", getAllUsers)
		user.DELETE("/delete/:id", deleteUser)

	}

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", signUp)
		auth.POST("/sign-in", signIn)
	}

	todo := router.Group("/todo")
	{
		todo.GET("/:todoId", getTodoByTodoId)
		todo.GET("/list/:id", getTodosByListId)
		todo.POST("/create", createTodo)
		todo.DELETE("/delete/:todoId", deleteTodo)
		todo.PUT("/update", updateTodo)

	}

	list := router.Group("/list")
	{
		//это следует изменить на получение айдишки через json
		list.GET("/:userId", getAllLists)
		list.POST("/create", createList)
		list.DELETE("/delete/:listId", deleteList)
		list.PUT("/update", updateList)

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
	userId := ctx.Param("id")

	err := DeleteUser(StringToInt(userId))

	if err != nil {
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
//нужно будет генерровать токен
func signIn(ctx *gin.Context){
	var userLogin UserLogin
	
	if err := ctx.BindJSON(&userLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	if err := IsValidUserLogin(userLogin); err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":"you're has been logined",
	})
	return
}


func getTodoByTodoId( ctx *gin.Context){
	todoId := ctx.Param("todoId")

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
	listId := ctx.Param("id")
	
	data, err := GetTodosByListId(StringToInt(listId))
	
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
	todoId := ctx.Param("todoId")
	
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
	listId := ctx.Param("listId")
	
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
	//нужно продумать получение айдишки пользователя
	userId := ctx.Param("userId")

	todoLists, err := GetAllLists(StringToInt(userId))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": todoLists,
	})
	return
}
