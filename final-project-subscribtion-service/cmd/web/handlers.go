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
	// best practice to renew the token which stored in the session each time you login or logout a user
	err := app.Session.RenewToken(r.Context())
	if err != nil {
		app.ErrLog.Println("error while renewing the token of the session : ", err)
	}

	// extract data from the form of the request
	err = r.ParseForm()
	if err != nil {
		app.ErrLog.Println("error while parsing the request data : ", err)

	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, err := app.models.UserRepo.GetByEmail(email)
	if err != nil {
		app.ErrLog.Println("error while fetching user by email : ", err)
		// put something in the session
		app.Session.Put(r.Context(), "error", "Invalid Credentials.")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check if the password is the same as the hashed stored password
	isPassMatches, err := user.CheckMatchedPasswords(password)
	if err != nil || !isPassMatches {
		app.ErrLog.Println("error while signing the user in : ", err)
		// put something in the session
		app.Session.Put(r.Context(), "error", "Invalid Credentials.")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// log the user in
	app.Session.Put(r.Context(), "userID", user.Id)
	app.Session.Put(r.Context(), "user", user) // in order to register the user data type into the session we have to configure this in the  session intiailizaiton function
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("logged-in successfully !")
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
