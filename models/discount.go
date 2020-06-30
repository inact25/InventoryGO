package models

import (
	"database/sql"
	"log"
)

type Discount struct {
	DiscountID,
	DiscountDesc string
}

func GetDataDiscount(db *sql.DB, pageNo, totalPerPage int) ([]*Discount, error) {
	//it is a good practice to always use the LIMIT clause with the ORDER BY clause to constraint the result rows in unique order.
	rows, err := db.Query(`
		select * from discount
		limit ?,?
		`, pageNo, totalPerPage)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	defer rows.Close()

	discount := make([]*Discount, 0)

	for rows.Next() {
		//new => reserve 1 memory allocation with certain data type pb := new(Product)
		d := new(Discount)
		//c := new(Category)
		err := rows.Scan(&d.DiscountID, &d.DiscountDesc)
		if err != nil {
			log.Fatalf("%v", err)
			return nil, err

		}
		discount = append(discount, d)
	}

	return discount, nil
}

func AddNewDiscount(db *sql.DB, discount *Discount) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("%v", err)
		return "", err
	}
	stmt, err := db.Prepare("insert into discount values(?,?)")
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(discount.DiscountID, discount.DiscountDesc); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	return "", tx.Commit()
}

func DeleteDiscount(db *sql.DB, discount *Discount) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("%v", err)
		return "", err
	}
	stmt, err := db.Prepare("delete from discount where discountID = ?")
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(discount.DiscountID); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	return "", tx.Commit()
}
