package main

import (
	"database/sql"
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/FadyGamilM/subscriptionservice/data"
	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

const Port = "80"
const DSN = "host=localhost port=5432 user=concurrency password=concurrency dbname=concurrencyDB sslmode=disable timezone=UTC connect_timeout=5"
const REDIS_CS = "127.0.0.1:6379"

func main() {

	//! 1. initiate the database connection and get instance of the connection pool
	db := init_db()

	//! 2. setup sessions with redis as persist store
	session := init_session()

	//! 3. define channels

	//! 4. define waitgroups
	wg := sync.WaitGroup{}

	//! 5. create loggers for info and errors
	info_logger := log.New(os.Stdout, "INFO ➜ \t", log.Ldate|log.Ltime)
	error_logger := log.New(os.Stdout, "ERROR ➜ \t", log.Ldate|log.Ltime|log.Lshortfile)

	//! 6. setup the app config and initialize it
	app := Config{
		DB:      db,
		Session: session,
		InfoLog: info_logger,
		ErrLog:  error_logger,
		wait:    &wg,
		models:  data.New(db),
	}

	//! 7.setup the email

	//! 8. running a go routine in the background to handle the shutdown for any termination or interrupt to get gracefull shutdown
	go app.ListenForShutdown()

	//! 9. Listen for requests
	app.Serve()

}

// initialize the database and return the pool of connection
func init_db() *sql.DB {
	conn := connect_to_db()
	if conn == nil {
		log.Fatalln("Cannot connect to the database!")
	}

	return conn
}

// the connection logic to postgres db and returns the pool of connection
func connect_to_db() *sql.DB {
	// num of trials to setup a connection
	count_of_trials := 0

	// dsn := os.Getenv("DSN")
	dsn := DSN

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
		count_of_trials += 1
		// repeat the loop again
		continue
	}
}

// open a connection to a database using pgx driver
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

// ping to the instance to test the databse connection
func ping_db(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	} else {
		log.Println("Ping to the database successfully!")
		return nil
	}
}

func init_session() *scs.SessionManager {
	// register data types into the session
	gob.Register(data.User{})

	// setup the session
	session := scs.New()
	session.Store = redisstore.New(init_redis())
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	return session
}

func init_redis() *redis.Pool {
	redis_pool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", REDIS_CS)
		},
	}
	return redis_pool
}
