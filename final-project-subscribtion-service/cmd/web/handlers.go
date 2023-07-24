package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Id   int
	Name string
}

func (app *Config) HomeHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Id:   1,
		Name: "fady",
	}
	res := AppResponse{
		Response: &user,
		Status:   http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}

func (app *Config) LoginHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Id:   1,
		Name: "fady",
	}
	res := AppResponse{
		Response: &user,
		Status:   http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}

func (app *Config) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Id:   1,
		Name: "fady",
	}
	res := AppResponse{
		Response: &user,
		Status:   http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}

func (app *Config) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Id:   1,
		Name: "fady",
	}
	res := AppResponse{
		Response: &user,
		Status:   http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}

func (app *Config) BuySubscribtionHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Id:   1,
		Name: "fady",
	}
	res := AppResponse{
		Response: &user,
		Status:   http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}

func (app *Config) ActivateAccountHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Id:   1,
		Name: "fady",
	}
	res := AppResponse{
		Response: &user,
		Status:   http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
