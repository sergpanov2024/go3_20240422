package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sergpanov2024/go3/hw2/api"
	db "github.com/sergpanov2024/go3/hw2/db/sqlc"
	"github.com/sergpanov2024/go3/hw2/util"
)

// github.com/sergpanov2024/go3_20240422/hw2

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Can not read config file", err)
	}

	pool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("Can not connect to db", err)
	}

	defer pool.Close()

	store := db.NewStore(pool)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Can not create server", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Can not start server", err)
	}

}
