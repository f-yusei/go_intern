package object

import (
	"time"
)

type Timeline struct {
	ID        int       `json:"id,omitempty"`
	Account   Account   ``
	URL       *string   `json:"url,omitempty" db:"url"`
	Content   string    `json:"status"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
}

func NewTimeline(statuses []Status) *Timeline {
	return &Timeline{}
}
