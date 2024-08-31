package main

import (
	"context"
	"fmt"
	"github.com/Dennisblay/ordering-app-server/config"
	"github.com/Dennisblay/ordering-app-server/internal/api"
	db "github.com/Dennisblay/ordering-app-server/internal/database/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func main() {
	configEnv, err := config.LoadConfig(".")
	if err != nil {
		fmt.Println("Cannot load configurations", "error", err)

		return
	}
	fmt.Println("Configurations loaded")
	store, err := InitDb(configEnv.DBUrl)
	s, err := api.NewServer(store)
	if err != nil {
		log.Fatal("cannot start api:", err)
	}

	err = s.RunServer(configEnv.ServerAddress)
	if err != nil {
		log.Fatal("cannot start api:", err)
	}

}

func InitDb(DBUrl string) (db.Store, error) {
	conn, err := pgxpool.New(context.Background(), DBUrl)
	if err != nil {
		return nil, err
	}
	return db.NewStore(conn), nil
}
