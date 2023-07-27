package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

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

// starts a http server
func (app *Config) Serve() {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", Port),
		Handler: app.routes(),
	}
	app.InfoLog.Println("starting the server on port ", Port)
	err := server.ListenAndServe()
	if err != nil {
		app.ErrLog.Println("error while starting the server on port ", Port, " : ", err)
		app.ErrLog.Panic(err)
	}
}

// gracefull shutdown implementation
// this is running on the background and lsiten for something
func (app *Config) ListenForShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // block and wait for any interrupt or terminate signal to come
	app.shutdown()
	os.Exit(0)
}

func (app *Config) shutdown() {
	// cleanup tasks
	app.InfoLog.Println("Running cleanup tasks after all running go routines in the background finish ...")

	// wait for all other go-routines ot finish their running tasks
	app.wait.Wait()

	// shutdown and close everything
	app.InfoLog.Println("Closing channels and shutting down the server ...")
}
