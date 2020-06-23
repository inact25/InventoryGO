package models

import (
	"database/sql"
	"log"
)

type Report struct {
	TransactionDate,
	TransactionCode,
	ProductDesc,
	TransactionQTY,
	ProductPrice,
	SubTotal string
}

func GetDailyReport(db *sql.DB, date string, pageNo, totalPerPage int) ([]*Report, error) {
	//it is a good practice to always use the LIMIT clause with the ORDER BY clause to constraint the result rows in unique order.
	rows, err := db.Query(`
		select t.transactionDate,t.transactionCode,p.productDesc,t.transactionQTY, 
		p.productPrice, (p.productPrice * t.transactionQTY) as subtotal from transaction t 
		inner join product p on t.productID = p.productID where transactionDate like ?
		limit ?,?
		`, date, pageNo, totalPerPage)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	defer rows.Close()

	report := make([]*Report, 0)

	for rows.Next() {
		//new => reserve 1 memory allocation with certain data type pb := new(Product)
		r := new(Report)
		//c := new(Category)
		err := rows.Scan(&r.TransactionDate, &r.TransactionCode, &r.ProductDesc, &r.TransactionQTY, &r.ProductPrice, &r.SubTotal)
		if err != nil {
			log.Fatalf("%v", err)
			return nil, err
		}
		report = append(report, r)
	}

	return report, nil
}

func GetAllReport(db *sql.DB, pageNo, totalPerPage int) ([]*Report, error) {
	//it is a good practice to always use the LIMIT clause with the ORDER BY clause to constraint the result rows in unique order.
	rows, err := db.Query(`
		select t.transactionDate,t.transactionCode,p.productDesc,t.transactionQTY, 
		p.productPrice, (p.productPrice * t.transactionQTY) as subtotal from transaction t 
		inner join product p on t.productID = p.productID
		limit ?,?
		`, pageNo, totalPerPage)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	defer rows.Close()

	report := make([]*Report, 0)

	for rows.Next() {
		//new => reserve 1 memory allocation with certain data type pb := new(Product)
		r := new(Report)
		//c := new(Category)
		err := rows.Scan(&r.TransactionDate, &r.TransactionCode, &r.ProductDesc, &r.TransactionQTY, &r.ProductPrice, &r.SubTotal)
		if err != nil {
			log.Fatalf("%v", err)
			return nil, err
		}
		report = append(report, r)
	}

	return report, nil
}

func GetMonthReport(db *sql.DB, date string, pageNo, totalPerPage int) ([]*Report, error) {
	//it is a good practice to always use the LIMIT clause with the ORDER BY clause to constraint the result rows in unique order.
	rows, err := db.Query(`
		select t.transactionDate,t.transactionCode,p.productDesc,t.transactionQTY, 
		p.productPrice, (p.productPrice * t.transactionQTY) as subtotal from transaction t 
		inner join product p on t.productID = p.productID where transactionDate like ?
		limit ?,?
		`, date, pageNo, totalPerPage)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	defer rows.Close()

	report := make([]*Report, 0)

	for rows.Next() {
		//new => reserve 1 memory allocation with certain data type pb := new(Product)
		r := new(Report)
		//c := new(Category)
		err := rows.Scan(&r.TransactionDate, &r.TransactionCode, &r.ProductDesc, &r.TransactionQTY, &r.ProductPrice, &r.SubTotal)
		if err != nil {
			log.Fatalf("%v", err)
			return nil, err
		}
		report = append(report, r)
	}

	return report, nil
}
