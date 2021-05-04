package interfaces

import (
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
