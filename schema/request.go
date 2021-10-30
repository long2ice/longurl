package schema

import "time"

type Url struct {
	Path     string     `json:"path"`
	Url      string     `json:"url"`
	ExpireAt *time.Time `json:"expire_at"`
}
