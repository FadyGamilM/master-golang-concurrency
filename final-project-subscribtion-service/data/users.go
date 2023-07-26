// CREATE TABLE IF NOT EXISTS users (
//     id BIGSERIAL PRIMARY KEY,
//     email VARCHAR NOT NULL,
//     first_name VARCHAR NOT NULL,
//     last_name VARCHAR NOT NULL,
//     password VARCHAR NOT NULL,
//     is_active BOOLEAN DEFAULT(TRUE),
//     is_admin BOOLEAN DEFAULT(FALSE),
//         created_at TIMESTAMP DEFAULT now() NOT NULL,
//     updated_at TIMESTAMP DEFAULT now() NOT NULL
// );

package data

import (
	"context"
	"time"
)

const get_all_users_query = `SELECT id, email, first_name, last_name, password, is_active, is_admin, created_at, updated_at FROM users`
const get_user_by_id = `SELECT id, email, first_name, last_name, password, is_active, is_admin, created_at, updated_at FROM users WHERE id = $1`
const get_user_by_email = `SELECT id, email, first_name, last_name, password, is_active, is_admin, created_at, updated_at FROM users WHERE email = $1`
const get_plan_by_user_id = 

type User struct {
	Id        int64
	Email     string
	FirstName string
	LastName  string
	Password  string
	IsActive  bool
	IsAdmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
	Plan *Plan
}

func (u *User) GetAll() ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, email, first_name, last_name, password, is_active, is_admin, created_at, updated_at FROM users`

	rows, err := 
}
