package models

import (
	"database/sql"
	"log"
)

type Brand struct {
	BrandID,
	BrandDesc string
}

func GetDataBrand(db *sql.DB, pageNo, totalPerPage int) ([]*Brand, error) {
	//it is a good practice to always use the LIMIT clause with the ORDER BY clause to constraint the result rows in unique order.
	rows, err := db.Query(`
		select * from brand
		limit ?,?
		`, pageNo, totalPerPage)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	defer rows.Close()

	brand := make([]*Brand, 0)

	for rows.Next() {
		//new => reserve 1 memory allocation with certain data type pb := new(Product)
		b := new(Brand)
		//c := new(Category)
		err := rows.Scan(&b.BrandID, &b.BrandDesc)
		if err != nil {
			log.Fatalf("%v", err)
			return nil, err

		}
		brand = append(brand, b)
	}

	return brand, nil
}

func AddNewBrand(db *sql.DB, brand *Brand) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("%v", err)
		return "", err
	}
	stmt, err := db.Prepare("insert into brand values	(?,?)")
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(brand.BrandID, brand.BrandDesc); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	return "", tx.Commit()
}

func DeleteBrand(db *sql.DB, brand *Brand) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("%v", err)
		return "", err
	}
	stmt, err := db.Prepare("delete from brand where brandID = ?")
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(brand.BrandID); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	return "", tx.Commit()
}
