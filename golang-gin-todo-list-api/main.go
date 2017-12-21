package main

import (
	"log"
	"os"

	"golang-gin-todo-list-api/controllers"
	"golang-gin-todo-list-api/routes"
	"golang-gin-todo-list-api/utils/database"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.Connect(os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASS"), os.Getenv("PGDB"))
	if err != nil {
		log.Fatal("err", err)
	}
	todoController := controllers.NewTodoController(db)
	r := gin.Default()
	routes.CreateRoutes(r, todoController)
	r.Run()
}
