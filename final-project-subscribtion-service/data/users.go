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
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const get_all_users_query = `SELECT id, email, first_name, last_name, password, is_active, is_admin, created_at, updated_at FROM users`
const get_user_by_id = `SELECT id, email, first_name, last_name, password, is_active, is_admin, created_at, updated_at FROM users WHERE id = $1`
const get_user_by_email = `SELECT id, email, first_name, last_name, password, is_active, is_admin, created_at, updated_at FROM users WHERE email = $1`
const get_plan_by_user_id = `
	SELECT p.id, p.plan_name, p.plan_amount, p.created_at, p.updated_at
	FROM users_plans as up 
	JOIN plans as p
	ON p.id = up.plan_id
	WHERE up.user_id = $1	
`

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
	Plan      *Plan
}

func (u *User) GetAll() ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, get_all_users_query)
	if err != nil {
		return nil, err
	}

	var users []*User
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.Id,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.Password,
			&user.IsActive,
			&user.IsAdmin,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (u *User) GetByEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows := db.QueryRowContext(ctx, get_user_by_email, email)

	var user User
	err := rows.Scan(
		&user.Id,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.IsActive,
		&user.IsAdmin,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) GetById(id int64) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows := db.QueryRowContext(ctx, get_user_by_id, id)

	var user User
	err := rows.Scan(
		&user.Id,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.IsActive,
		&user.IsAdmin,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	// get the plan if there are any ..
	planRow := db.QueryRowContext(ctx, get_plan_by_user_id, id)
	var plan Plan
	err = planRow.Scan(
		&plan.Id,
		&plan.PlanName,
		&plan.PlanAmount,
		&plan.CreatedAt,
		&plan.UpdatedAt,
	)
	if err != nil {
		return &user, nil
	}

	user.Plan = &plan
	return &user, nil
}

func (u *User) ResetPassword(newPass string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(newPass), 12)
	if err != nil {
		return err
	}

	query := `UPDATE users SET password = $1 WHERE id = $2`

	_, err = db.ExecContext(ctx, query, hashedPass, u.Id)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) CheckMatchedPasswords(newPass string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(newPass))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid pass
			return false, nil
		default:
			// another error than mismatching
			return false, err
		}
	}
	return true, nil
}
