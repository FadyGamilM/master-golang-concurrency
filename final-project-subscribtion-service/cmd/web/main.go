package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

const Port = "80"

func main() {
	// initiate the database connection and get instance of the connection pool
	db := init_db()
}

func init_db() *sql.DB {
	conn := connect_to_db()
	if conn == nil {
		log.Fatalln("Cannot connect to the database!")
	}

	return conn
}

func connect_to_db() *sql.DB {
	// num of trials to setup a connection
	count_of_trials := 0

	dsn := os.Getenv("DSN")

	for {
		conn, err := open_db(dsn)
		if err != nil {
			log.Println("postgres instance is not ready yet ...")
		} else {
			log.Println("connected to postgres instance successfully ! ")
			return conn
		}

		// if we didn't return thats means we can't connect yet, so we try again if we have available trials still
		if count_of_trials > 10 {
			log.Println("we tried to connect to database instance 10 times but there is something wrong !")
			return nil
		}
		log.Println("Waiting for one second and try to connect again ..")
		time.Sleep(1 * time.Second)
		// repeat the loop again
		continue
	}

}

func open_db(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = ping_db(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ping_db(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	} else {
		log.Println("Ping to the database successfully!")
		return nil
	}
}
