package domain

import (
	"time"
)

// DieMasterRepository : Operations on DieMaster entity
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

// DieMasterRepository : Operations on DieMaster entity
type DieOrderRepository interface {
	CreateDieOrder(orders []DieOrderLine) error
}

// Die : Domain entity representing a Die in aluminum extrusion
type DieOrderLine struct {
	LotNumber         int
	Sl                int
	BolsterNumber     string
	FirstExtReqWeight string
	SolidLeadPI       bool
	SolidDiePI        bool
	SolidBacker       bool
	PortholeDie       bool
	PortholeMandrel   bool
	PortholeBacker    bool
	Description       string
	Size              string
	Kgs               string
	Sup               string
	Price             string
	Remarks           string
	DieNumber         string
	CavNumber         int
	CompanyName       string
	Email             string
	Address           string
}
