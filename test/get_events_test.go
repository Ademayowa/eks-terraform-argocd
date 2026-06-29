package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	db "github.com/Ademayowa/eks-terraform-argocd/internal/database"
	events "github.com/Ademayowa/eks-terraform-argocd/internal/events"
	"github.com/gin-gonic/gin"
	"github.com/pashagolub/pgxmock/v2"
	"github.com/stretchr/testify/assert"
)

func TestListEvents(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error: %s", err)
	}
	defer mock.Close()
	db.Pool = mock

	mock.ExpectQuery(`(?i)SELECT\s+.*\s+FROM\s+events`).
		WillReturnRows(pgxmock.NewRows([]string{"id", "title", "location", "date", "description", "created_at"}).
			AddRow("mock-id-1", "Demo 1", "UK", time.Now(), "Lean", time.Now()).
			AddRow("mock-id-2", "Demo 2", "US", time.Now(), "Agile", time.Now()))

	req := httptest.NewRequest("GET", "/events", nil)
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	repo := events.NewRepository(mock)
	svc := events.NewService(repo)
	h := events.NewHandler(svc)

	r.GET("/events", h.List)
	r.ServeHTTP(w, req)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("mock error: %s", err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
}
