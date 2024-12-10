package user

import "time"

type User struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	SupabaseID string    `json:"supabase_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
