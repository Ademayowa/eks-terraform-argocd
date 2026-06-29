package events

import (
	db "github.com/Ademayowa/eks-terraform-argocd/internal/database"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, databasePool db.Querier) {
	repo := NewRepository(databasePool)
	svc := NewService(repo)
	h := NewHandler(svc)

	// Events routes
	r.POST("/events", h.Create)
	r.GET("/events", h.List)
	r.GET("/events/:id", h.Get)
}
