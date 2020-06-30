package models

import (
	"database/sql"
	"log"
)

type User struct {
	UserID,
	UserName,
	UserPass string
}

func GetDataUser(db *sql.DB, userName, userPass string) ([]*User, error) {
	//it is a good practice to always use the LIMIT clause with the ORDER BY clause to constraint the result rows in unique order.
	rows, err := db.Query(`
		select * from user where userName = ? and userPass = ?
		`, userName, userPass)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	defer rows.Close()

	user := make([]*User, 0)

	for rows.Next() {
		//new => reserve 1 memory allocation with certain data type pb := new(Product)
		u := new(User)
		//c := new(Category)
		err := rows.Scan(&u.UserID, &u.UserName, &u.UserPass)
		if err != nil {
			log.Fatalf("%v", err)
			return nil, err

		}
		user = append(user, u)
	}

	return user, nil
}

func GetAllDataUser(db *sql.DB, pageNo, totalPerPage int) ([]*User, error) {
	//it is a good practice to always use the LIMIT clause with the ORDER BY clause to constraint the result rows in unique order.
	rows, err := db.Query(`
		select userID, userName from user
		limit ?,?
		`, pageNo, totalPerPage)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	defer rows.Close()

	user := make([]*User, 0)

	for rows.Next() {
		//new => reserve 1 memory allocation with certain data type pb := new(Product)
		u := new(User)
		//c := new(Category)
		err := rows.Scan(&u.UserID, &u.UserName)
		if err != nil {
			log.Fatalf("%v", err)
			return nil, err

		}
		user = append(user, u)
	}

	return user, nil
}

func AddNewUser(db *sql.DB, user *User) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("%v", err)
		return "", err
	}
	stmt, err := db.Prepare("insert into user (userName,userPass) values (?,?)")
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.UserName, user.UserPass); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	return "", tx.Commit()
}

func DeleteUser(db *sql.DB, user *User) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("%v", err)
		return "", err
	}
	stmt, err := db.Prepare("delete from user where userName = ?")
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.UserName); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	return "", tx.Commit()
}
