//DB Connection example

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "admin:MyPassword@tcp(mydb.cgx7wslwcmx7.us-east-1.rds.amazonaws.com:3306)/db")

	if err != nil {
		panic(err)
	}

	defer db.Close()
	sql := "insert into dept values(1111, 'ITT','Pune')"

	fmt.Println("Connection open  ", db)
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}

	rows, _ := result.RowsAffected()
	fmt.Println(rows)

	//select query to db
	res, err := db.Query("SELECT * FROM dept")

	defer res.Close()

	if err != nil {
		panic(err)
	}

	for res.Next() {
		deptno, dname, loc := 1, "a", "a"
		res.Scan(&deptno, &dname, &loc)
		fmt.Println(deptno, ", ", dname, ", ", loc)
	}

}
