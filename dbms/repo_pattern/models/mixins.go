package models

import "time"

type DateTimeMixin struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}
