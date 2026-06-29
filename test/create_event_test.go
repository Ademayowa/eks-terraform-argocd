package test

import (
	"bytes"
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

func TestCreateEvent(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error: %s", err)
	}
	defer mock.Close()
	db.Pool = mock

	parsedTime, _ := time.Parse("2006-01-02", "2026-08-15")

	mock.ExpectQuery(`(?i)INSERT\s+INTO\s+events`).
		WithArgs("Demo", "UK", parsedTime, "Lean").
		WillReturnRows(pgxmock.NewRows([]string{"id", "created_at"}).AddRow("mock-id", time.Now()))

	body := `{"title":"Demo","location":"UK","date":"2026-08-15","description":"Lean"}`
	req := httptest.NewRequest("POST", "/events", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	repo := events.NewRepository(mock)
	svc := events.NewService(repo)
	h := events.NewHandler(svc)

	r.POST("/events", h.Create)
	r.ServeHTTP(w, req)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("mock error: %s", err)
	}

	if w.Code != http.StatusCreated {
		t.Logf("Response Body: %s", w.Body.String())
	}

	assert.Equal(t, http.StatusCreated, w.Code)
}
