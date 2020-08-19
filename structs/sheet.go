package structs

import (
	"time"
)

type Sheet struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	Image     string    `json:"image"`
	Note      int       `json:"note"`
	Styles    []uint8   `json:"styles"`
	Synopsis  string    `json:"synopsis"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
