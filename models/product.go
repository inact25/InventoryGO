package models

import (
	"database/sql"
	"log"
)

type Product struct {
	ProductID,
	BrandDesc,
	ProductDesc,
	CategoryID,
	DiscountID,
	ProductPrice string
}

func GetDataProduct(db *sql.DB, pageNo, totalPerPage int) ([]*Product, error) {
	//it is a good practice to always use the LIMIT clause with the ORDER BY clause to constraint the result rows in unique order.
	rows, err := db.Query(`
		select p.productID,b.brandDesc,p.productDesc, c.categoryDesc,p.productPrice, d.discountDesc from product p 
		inner join category c on p.categoryID = c.categoryID
		inner join discount d on p.discountID = d.discountID
		inner join brand b on p.brandCode = b.brandId order by p.productID
		limit ?,?
		`, pageNo, totalPerPage)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	defer rows.Close()

	products := make([]*Product, 0)

	for rows.Next() {
		//new => reserve 1 memory allocation with certain data type pb := new(Product)
		p := new(Product)
		//c := new(Category)
		err := rows.Scan(&p.ProductID, &p.BrandDesc, &p.ProductDesc, &p.CategoryID, &p.DiscountID, &p.ProductPrice)
		if err != nil {
			log.Fatalf("%v", err)
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func AddNewProduct(db *sql.DB, product *Product) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("%v", err)
		return "", err
	}
	stmt, err := db.Prepare("insert into product values	(?,?,?,?,?,?)")
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(product.ProductID, product.ProductDesc, product.CategoryID, product.ProductPrice, product.DiscountID, product.BrandDesc); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	return "", tx.Commit()
}

func DeleteProduct(db *sql.DB, product *Product) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("%v", err)
		return "", err
	}
	stmt, err := db.Prepare("delete from product where productID = ?")
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(product.ProductID); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	return "", tx.Commit()
}
