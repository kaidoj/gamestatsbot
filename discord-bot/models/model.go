package models

import (
	"time"
)

//ModelInterface structure of model
type ModelInterface interface {
	Insert() error
	Update() error
}

//Model fields we always have
type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}
