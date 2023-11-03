package test

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest/v3"
)

type MySQLContainer struct {
	Pool     *dockertest.Pool
	Resource *dockertest.Resource
	DB       *sql.DB
}

func RunMySQLContainer() *MySQLContainer {
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Second * 30
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	resource, err := pool.Run("mysql", "latest", []string{"MYSQL_ROOT_PASSWORD=secret"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	var db *sql.DB

	if err = pool.Retry(func() error {
		var err error
		db, err = sql.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/mysql?parseTime=true&multiStatements=true", resource.GetPort("3306/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	log.Println("MySQL container start🐳")

	return &MySQLContainer{
		Pool:     pool,
		Resource: resource,
		DB:       db,
	}
}

func (mc *MySQLContainer) Close() {
	// When you're done, kill and remove the container
	if err := mc.Pool.Purge(mc.Resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	if err := mc.DB.Close(); err != nil {
		log.Fatalf("Could not close db: %s", err)
	}
}
