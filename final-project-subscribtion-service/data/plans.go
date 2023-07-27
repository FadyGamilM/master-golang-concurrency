package data

import (
	"context"
	"time"
)

type Plan struct {
	Id         int64
	PlanName   string
	PlanAmount float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

const get_all_plans = `SELECT id, plan_name, plan_amount, created_at, updated_at FROM plans`

func (p *Plan) GetAll() ([]*Plan, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, get_all_plans)
	if err != nil {
		return nil, err
	}

	var plans []*Plan
	for rows.Next() {
		var plan Plan
		err := rows.Scan(
			&plan.Id,
			&plan.PlanName,
			&plan.PlanAmount,
			&plan.CreatedAt,
			&plan.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		plans = append(plans, &plan)
	}

	return plans, nil
}

func (u *User) SubscribeToPlan(user User, plan Plan) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO users_plans (user_id, plan_id) VALUES ($1, $2)`

	_, err := db.ExecContext(ctx, query, user.Id, plan.Id)
	if err != nil {
		return err
	}

	return nil
}
