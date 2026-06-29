package events

import (
	"context"

	db "github.com/Ademayowa/eks-terraform-argocd/internal/database"
)

type Repository struct {
	db db.Querier
}

func NewRepository(db db.Querier) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Insert(ctx context.Context, e Event) (*Event, error) {
	err := r.db.QueryRow(ctx, `
		INSERT INTO events (title, location, date, description, created_at)
		VALUES ($1, $2, $3, $4, NOW())
		RETURNING id, created_at`,
		e.Title, e.Location, e.Date, e.Description,
	).Scan(&e.ID, &e.CreatedAt)

	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *Repository) GetAll(ctx context.Context) ([]Event, error) {
	rows, err := r.db.Query(ctx, `SELECT id, title, location, date, description, created_at FROM events ORDER BY date ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var evts []Event = make([]Event, 0)
	for rows.Next() {
		var e Event
		if err := rows.Scan(&e.ID, &e.Title, &e.Location, &e.Date, &e.Description, &e.CreatedAt); err != nil {
			return nil, err
		}
		evts = append(evts, e)
	}
	return evts, nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (*Event, error) {
	var e Event
	err := r.db.QueryRow(ctx, `
		SELECT id, title, location, date, description, created_at 
		FROM events WHERE id = $1`, id,
	).Scan(&e.ID, &e.Title, &e.Location, &e.Date, &e.Description, &e.CreatedAt)

	if err != nil {
		return nil, err
	}
	return &e, nil
}
