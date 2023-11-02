package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type dbConfig struct {
	User     string
	Pass     string
	Host     string
	Port     int
	Database string
}

func newDBConfig() (*dbConfig, error) {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")
	host := os.Getenv("MYSQL_HOST")
	port, err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	if err != nil {
		return nil, err
	}
	database := os.Getenv("MYSQL_DBNAME")

	return &dbConfig{
		user,
		pass,
		host,
		port,
		database,
	}, nil
}

func initializeDB(config *dbConfig) (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		config.User,
		config.Pass,
		config.Host,
		config.Port,
		config.Database,
	)

	return sql.Open("mysql", connectionString)
}

const ServerPort = 8080

func main() {
	config, err := newDBConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := initializeDB(config)
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
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", ServerPort)))
}
