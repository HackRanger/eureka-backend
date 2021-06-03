package interfaces

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/eureka/domain"
)

type DbConnection struct {
	db *sql.DB
}

type DbDie DbConnection

func NewDbDie(db *sql.DB) *DbDie {
	dbDie := new(DbDie)
	dbDie.db = db
	return dbDie
}

type DbDieOrder DbConnection

func NewDbDieOrder(db *sql.DB) *DbDieOrder {
	dbDieOrder := new(DbDieOrder)
	dbDieOrder.db = db
	return dbDieOrder
}

func (dbDie *DbDie) GetAllDie() ([]domain.Die, error) {
	var allDie []domain.Die
	rows, err := dbDie.db.Query("select * from diemasterregistry")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var lotNumber int
		var dieNumber string
		var nosCavity int
		var dieType string
		var bolsterNo string
		var supplier string
		var dieSize string
		var orderDate string
		var landedDate string
		var price int
		err = rows.Scan(&lotNumber, &dieNumber, &nosCavity, &dieType, &bolsterNo, &supplier, &dieSize, &orderDate, &landedDate, &price)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(lotNumber, dieNumber, nosCavity, dieType, bolsterNo, supplier, dieSize, orderDate, landedDate, price)

		oDate, err := time.Parse(time.RFC3339, orderDate)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		lDate, err := time.Parse(time.RFC3339, landedDate)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		die := domain.Die{
			LotNumber:  lotNumber,
			DieNumber:  dieNumber,
			NosCavity:  nosCavity,
			DieType:    bolsterNo,
			Supplier:   supplier,
			DieSize:    dieSize,
			BolsterNo:  bolsterNo,
			OrderDate:  oDate,
			LandedDate: lDate,
			Price:      price,
		}
		allDie = append(allDie, die)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return allDie, nil
}

func (dbDieOrder *DbDieOrder) CreateDieOrder(orderLines []domain.DieOrderLine) error {
	// Create a new context, and begin a transaction
	ctx := context.Background()
	tx, err := dbDieOrder.db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, order := range orderLines {
		_, err = tx.ExecContext(ctx, `INSERT INTO dieorder(lotnumber, sl, bolsternumber, 
					firstextreqweight, solidleadpi, soliddiepi, solidbacker, portholedie,
					portholemandrel, portholebacker, description, size, kgs, sup, price,
					remarks, dienumber, cavnumber, companyname, email, address,orderDate)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21,$22)`,
			order.LotNumber,
			order.Sl,
			order.BolsterNumber,
			order.FirstExtReqWeight,
			order.SolidLeadPI,
			order.SolidDiePI,
			order.SolidBacker,
			order.PortholeDie,
			order.PortholeMandrel,
			order.PortholeBacker,
			order.Description,
			order.Size,
			order.Kgs,
			order.Sup,
			order.Price,
			order.Remarks,
			order.DieNumber,
			order.CavNumber,
			order.CompanyName,
			order.Email,
			order.Address,
			order.Date)
		if err != nil {
			// Incase we find any error in the query execution, rollback the transaction
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		tx.Rollback()
		return err
	}

	return nil
}

func (dbDieOrder *DbDieOrder) GetAllDieOrders() ([]domain.DieOrderLine, error) {
	var allDieOrders []domain.DieOrderLine
	rows, err := dbDieOrder.db.Query("select * from dieorder")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
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
			Date              time.Time
		)
		err = rows.Scan(&LotNumber,
			&Sl,
			&BolsterNumber, &FirstExtReqWeight, &SolidLeadPI,
			&SolidDiePI,
			&SolidBacker,
			&PortholeDie,
			&PortholeMandrel,
			&PortholeBacker,
			&Description,
			&Size,
			&Kgs,
			&Sup,
			&Price,
			&Remarks,
			&DieNumber,
			&CavNumber,
			&CompanyName,
			&Email,
			&Address,
			&Date)
		if err != nil {
			log.Fatal(err)
		}

		dieOrder := domain.DieOrderLine{
			Sl:                Sl,
			LotNumber:         LotNumber,
			BolsterNumber:     BolsterNumber,
			FirstExtReqWeight: FirstExtReqWeight,
			SolidLeadPI:       SolidLeadPI,
			SolidDiePI:        SolidDiePI,
			SolidBacker:       SolidBacker,
			PortholeDie:       PortholeDie,
			PortholeMandrel:   PortholeMandrel,
			PortholeBacker:    PortholeBacker,
			Description:       Description,
			Size:              Size,
			Kgs:               Kgs,
			Sup:               Sup,
			Price:             Price,
			Remarks:           Remarks,
			DieNumber:         DieNumber,
			CavNumber:         CavNumber,
			CompanyName:       CompanyName,
			Email:             Email,
			Address:           Address,
			Date:              Date,
		}
		allDieOrders = append(allDieOrders, dieOrder)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return allDieOrders, nil
}

func (dbDieOrder *DbDieOrder) GenerateLotNumber() (int, error) {
	rows, err := dbDieOrder.db.Query("select lotnumber from dieorder order by sl desc limit 1")
	if err != nil {
		log.Fatal(err)
		return -1, err
	}

	defer rows.Close()
	var lotNumber int
	for rows.Next() {

		err = rows.Scan(&lotNumber)
		if err != nil {
			log.Fatal(err)
			return -1, err
		}
	}
	return lotNumber, nil
}
