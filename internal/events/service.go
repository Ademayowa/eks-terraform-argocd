package events

import (
	"context"
	"fmt"
	"time"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// Create new event
func (s *Service) CreateEvent(ctx context.Context, req EventRequest) (*Event, error) {
	parsedDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, fmt.Errorf("invalid date format, use YYYY-MM-DD")
	}

	evt := Event{
		Title:       req.Title,
		Location:    req.Location,
		Date:        parsedDate,
		Description: req.Description,
	}

	return s.repo.Insert(ctx, evt)
}

// List all events
func (s *Service) ListEvents(ctx context.Context) ([]Event, error) {
	return s.repo.GetAll(ctx)
}

// Get a single event by ID
func (s *Service) GetEvent(ctx context.Context, id string) (*Event, error) {
	return s.repo.GetByID(ctx, id)
}
