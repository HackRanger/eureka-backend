package domain

import (
	"time"
)

// BookRepository : Operations on Book entity
type DieMasterRepository interface {
	GetAllDie() ([]Die, error)
}

// Die : Domain entity representing a Die in aluminum extrusion
type Die struct {
	LotNumber  int
	DieNumber  string
	NosCavity  int
	DieType    string
	BolsterNo  string
	Supplier   string
	DieSize    string
	OrderDate  time.Time
	LandedDate time.Time
	Price      int
}
