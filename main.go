package main

import (
	"test-coding/db"
	"test-coding/handlers"
	"test-coding/repositories"

	"github.com/gin-gonic/gin"
)

func main() {
    db.Init()
    r := gin.Default()

    clientRepo := &repositories.ClientRepository{DB: db.DB}
    clientHandler := &handlers.ClientHandler{Repo: clientRepo, RDB: db.RDB}

    r.POST("/clients", clientHandler.CreateClient)
    r.PUT("/clients/:slug", clientHandler.UpdateClient)
    r.DELETE("/clients/:slug", clientHandler.DeleteClient)
    r.GET("/clients/:slug", clientHandler.GetClient)

    r.Run(":8080")
}