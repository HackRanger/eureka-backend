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
					remarks, dienumber, cavnumber, companyname, email, address)
					VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			order.LotNumber,
			order.Sl,
			order.BolsterNumber, order.FirstExtReqWeight, order.SolidLeadPI,
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
			order.Address)
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
