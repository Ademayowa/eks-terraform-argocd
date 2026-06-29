package events

import "time"

type EventRequest struct {
	Title       string `json:"title" binding:"required"`
	Location    string `json:"location" binding:"required"`
	Date        string `json:"date" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// Event represents the database model
type Event struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Location    string    `json:"location"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
