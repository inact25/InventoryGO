package models

import (
	"database/sql"
	"log"
)

type Transaction struct {
	TransactionCode,
	TransactionDate,
	BrandDesc,
	ProductDesc,
	TransactionQTY string
}

func GetDataTransaction(db *sql.DB, pageNo, totalPerPage int) ([]*Transaction, error) {
	//it is a good practice to always use the LIMIT clause with the ORDER BY clause to constraint the result rows in unique order.
	rows, err := db.Query(`
	select t.transactionCode, t.transactionDate,b.brandDesc, p.productDesc, t.transactionQTY from transaction t 
	inner join product p on t.productID = p.productID
	inner join brand b on b.brandId = p.brandCode
	limit ?,?
	`, pageNo, totalPerPage)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	defer rows.Close()

	transaction := make([]*Transaction, 0)

	for rows.Next() {
		//new => reserve 1 memory allocation with certain data type pb := new(Product)
		p := new(Transaction)
		//c := new(Category)
		err := rows.Scan(&p.TransactionCode, &p.TransactionDate, &p.BrandDesc, &p.ProductDesc, &p.TransactionQTY)
		if err != nil {
			log.Fatalf("%v", err)
			return nil, err
		}
		transaction = append(transaction, p)
	}

	return transaction, nil
}

func AddNewTransaction(db *sql.DB, transaction *Transaction) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("%v", err)
		return "", err
	}

	stmt, err := db.Prepare("insert into transaction (transactionCode, transactionDate, productID, transactionQTY) values ( ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(transaction.TransactionCode, transaction.TransactionDate, transaction.ProductDesc, transaction.TransactionQTY); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	return "", tx.Commit()
}

func DeleteTransaction(db *sql.DB, transaction *Transaction) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("%v", err)
		return "", err
	}
	stmt, err := db.Prepare("delete from transaction where transactionCode = ?")
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(transaction.TransactionCode); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	return "", tx.Commit()
}
