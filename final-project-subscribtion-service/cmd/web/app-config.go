package main

import (
	"database/sql"
	"log"
	"net/http"
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

func (app *Config) is_authenticated(r *http.Request) bool {
	return app.Session.Exists(r.Context(), "userID")
}

type AppResponse struct {
	Response interface{}
	Status   int32
}
