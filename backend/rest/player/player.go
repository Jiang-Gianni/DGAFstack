package player

import "time"

type Player struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Options   map[string]string
	Tags      []string
}
