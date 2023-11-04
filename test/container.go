package test

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

func init() {
	// コンテナが起動するまでPingを打ち続けるので接続できるまでエラーが出続けるので抑止
	mysql.SetLogger(mysql.Logger(log.New(io.Discard, "", 0)))
}

type MySQLContainer struct {
	Pool     *dockertest.Pool
	Resource *dockertest.Resource
	DB       *sql.DB
}

func RunMySQLContainer() (*MySQLContainer, error) {
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Minute * 1
	if err != nil {
		return nil, fmt.Errorf("could not construct pool: %w", err)
	}

	log.Println("connecting to Docker...")

	err = pool.Client.Ping()
	if err != nil {
		return nil, fmt.Errorf("could not connect to Docker: %w", err)
	}

	log.Println("success connecting to Docker🚀")

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "latest",
		Env: []string{
			"MYSQL_ROOT_PASSWORD=secret",
		},
	}, func(hc *docker.HostConfig) {
		hc.AutoRemove = true
		hc.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})

	if err != nil {
		return nil, fmt.Errorf("could not start resource: %w", err)
	}

	if err = resource.Expire(600); err != nil {
		return nil, fmt.Errorf("could not set container expire: %w", err)
	}

	var db *sql.DB

	log.Println("connecting to DB...")

	if err = pool.Retry(func() error {
		var err error
		db, err = sql.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/mysql?parseTime=true&multiStatements=true", resource.GetPort("3306/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		return nil, fmt.Errorf("could not connect to docker: %w", err)
	}

	log.Println("success connecting to DB🚀")
	log.Println("mysql container start🐳")

	return &MySQLContainer{
		Pool:     pool,
		Resource: resource,
		DB:       db,
	}, nil
}

func (mc *MySQLContainer) Close() {
	// When you're done, kill and remove the container
	if err := mc.Pool.Purge(mc.Resource); err != nil {
		log.Printf("could not purge resource: %s\n", err)
	}
	if err := mc.DB.Close(); err != nil {
		log.Printf("could not close db: %s\n", err)
	}
	log.Println("mysql container end🐳")
}
