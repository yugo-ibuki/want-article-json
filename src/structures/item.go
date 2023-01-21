package structures

import "time"

type Item struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}
