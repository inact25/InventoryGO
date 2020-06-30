package models

import (
	"database/sql"
	"log"
)

type Category struct {
	CategoryID,
	CategoryDesc string
}

func GetDataCategory(db *sql.DB, pageNo, totalPerPage int) ([]*Category, error) {
	//it is a good practice to always use the LIMIT clause with the ORDER BY clause to constraint the result rows in unique order.
	rows, err := db.Query(`
		select * from category
		limit ?,?
		`, pageNo, totalPerPage)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	defer rows.Close()

	category := make([]*Category, 0)

	for rows.Next() {
		//new => reserve 1 memory allocation with certain data type pb := new(Product)
		c := new(Category)
		//c := new(Category)
		err := rows.Scan(&c.CategoryID, &c.CategoryDesc)
		if err != nil {
			log.Fatalf("%v", err)
			return nil, err

		}
		category = append(category, c)
	}

	return category, nil
}

func AddNewCategory(db *sql.DB, category *Category) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("%v", err)
		return "", err
	}
	stmt, err := db.Prepare("insert into category values	(?,?)")
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(category.CategoryID, category.CategoryDesc); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	return "", tx.Commit()
}

func DeleteCategory(db *sql.DB, category *Category) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("%v", err)
		return "", err
	}
	stmt, err := db.Prepare("delete from category where brandID = ?")
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(category.CategoryID); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	return "", tx.Commit()
}
