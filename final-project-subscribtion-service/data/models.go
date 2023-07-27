package data

import "database/sql"

// this type is used to make our db functions (repos) available to the entire app
type Models struct {
	UserRepo User
	PlanRepo Plan
}

var db *sql.DB

func New(dbBool *sql.DB) Models {
	db = dbBool
	return Models{
		UserRepo: User{},
		PlanRepo: Plan{},
	}
}
