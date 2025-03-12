package handlers

import (
	"context"
	"encoding/json"
	"test-coding/models"
	"test-coding/repositories"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type ClientHandler struct {
    Repo *repositories.ClientRepository
    RDB  *redis.Client
}

func (h *ClientHandler) CreateClient(c *gin.Context) {
    var client models.Client
    if err := c.ShouldBindJSON(&client); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.Repo.Create(&client); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx := context.Background()
    clientJSON, _ := json.Marshal(client)
    h.RDB.Set(ctx, client.Slug, clientJSON, 0)

    c.JSON(http.StatusCreated, client)
}

func (h *ClientHandler) UpdateClient(c *gin.Context) {
    slug := c.Param("slug")
    var client models.Client
    if err := c.ShouldBindJSON(&client); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.Repo.Update(&client); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx := context.Background()
    h.RDB.Del(ctx, slug)
    clientJSON, _ := json.Marshal(client)
    h.RDB.Set(ctx, client.Slug, clientJSON, 0)

    c.JSON(http.StatusOK, client)
}

func (h *ClientHandler) DeleteClient(c *gin.Context) {
    slug := c.Param("slug")
    if err := h.Repo.Delete(slug); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx := context.Background()
    h.RDB.Del(ctx, slug)

    c.JSON(http.StatusOK, gin.H{"message": "Client deleted"})
}

func (h *ClientHandler) GetClient(c *gin.Context) {
    slug := c.Param("slug")
    ctx := context.Background()

    val, err := h.RDB.Get(ctx, slug).Result()
    if err == nil {
        var client models.Client
        json.Unmarshal([]byte(val), &client)
        c.JSON(http.StatusOK, client)
        return
    }

    client, err := h.Repo.GetBySlug(slug)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
        return
    }

    clientJSON, _ := json.Marshal(client)
    h.RDB.Set(ctx, slug, clientJSON, 0)

    c.JSON(http.StatusOK, client)
}