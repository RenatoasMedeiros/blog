package main

import (
	"log"

	env "github.com/RenatoasMedeiros/RenatoasMedeiros/internal"
	"github.com/RenatoasMedeiros/RenatoasMedeiros/internal/db"
	"github.com/RenatoasMedeiros/RenatoasMedeiros/internal/store"
)

const version = "0.0.1" // pre re

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost:5432/renatoasmedeiros?sslmode=disable"),
			maxOpenConns: env.GetInteger("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInteger("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}
	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	log.Println("Dabase connection pool established")
	store := store.NewStorage(db)
	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
