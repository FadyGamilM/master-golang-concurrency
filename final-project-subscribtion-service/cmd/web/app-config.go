package main

import (
	"database/sql"
	"log"
	"sync"

	"github.com/alexedwards/scs/v2"
)

// type to share the app config
type Config struct {
	Session *scs.SessionManager
	DB      *sql.DB
	InfoLog *log.Logger
	ErrLog  *log.Logger
	wait    *sync.WaitGroup
}
