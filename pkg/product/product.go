package producto

import "time"

// model of product
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Models slice of Model
type Models []*Model

type Storage interface {
	Create(*Model) error
	Update(*Model) error
	GetAll() ([]Models, error)
	GetByID(uint) (*Model, error)
	Delete(uint) error
}
