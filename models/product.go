package models

import (
	"database/sql"
	"log"
)

type Product struct {
	ProductID,
	ProductDesc,
	CategoryID,
	DiscountID string
	ProductPrice int
}

func GetDataProduct(db *sql.DB, pageNo, totalPerPage int) ([]*Product, error) {
	//it is a good practice to always use the LIMIT clause with the ORDER BY clause to constraint the result rows in unique order.
	rows, err := db.Query(`
		select p.productID,p.productDesc, c.categoryDesc,p.productPrice, d.discountDesc from product p 
		inner join category c on p.categoryID = c.categoryID
		inner join discount d on p.discountID = d.discountID order by p.productID 
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
		err := rows.Scan(&p.ProductID, &p.ProductDesc, &p.CategoryID, &p.DiscountID, &p.ProductPrice)
		if err != nil {
			log.Fatalf("%v", err)
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
