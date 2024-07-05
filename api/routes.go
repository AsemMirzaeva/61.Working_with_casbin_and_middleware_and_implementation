package api

import (
	"log"
	"os"
	"task/internal/repos"
	"task/internal/storage"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"

	_ "task/api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New ...
// @title	Project: Swagger Intro
// @description This swagger UI was created in lesson
// @version 1.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func Routes() {
	router := gin.Default()

	db, err := storage.OpenSql(os.Getenv("driver_name"), os.Getenv("postgres_url"))
	if err != nil {
		log.Println(err)
	}

	repo := repos.NewTaskRepo(db)
	handler := NewTaskHandler(repo)

	router.POST("/tasks", handler.CreateTask)
	router.GET("/tasks", handler.GetTask)
	router.GET("/tasks/:id", handler.GetTaskById)
	router.PUT("/tasks/:id", handler.UpdateTask)
	router.DELETE("/tasks/:id", handler.DeleteTask)

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	log.Println("Server is listening on port", os.Getenv("server_url"))
	log.Fatal(router.Run(":8080"))
}
