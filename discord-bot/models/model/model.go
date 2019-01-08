package models

import (
	"time"
)

type Model interface {
	Insert() error
	Update() error
}

type ModelImpl struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}
