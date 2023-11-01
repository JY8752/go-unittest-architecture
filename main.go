package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func initializeDB() (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		"root",
		os.Getenv("MYSQL_ROOT_PASSWORD"),
		"localhost",
		3306,
		os.Getenv("MYSQL_DATABASE"),
	)

	return sql.Open("mysql", connectionString)
}

func main() {
	db, err := initializeDB()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = db.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	e := echo.New()

	InitializeRootHandler(db, e).RegisterAll()
	e.Logger.Fatal(e.Start(":8080"))
}
