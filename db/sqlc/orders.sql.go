// Code generated by sqlc. DO NOT EDIT.
// source: orders.sql

package db

import (
	"context"
	"database/sql"
)

const createOrder = `-- name: CreateOrder :one
INSERT INTO orders (
  user_name,
  product_name,
  amount
) VALUES (
  $1, $2, $3
)RETURNING id, user_name, product_name, amount
`

type CreateOrderParams struct {
	UserName    sql.NullString
	ProductName sql.NullString
	Amount      int32
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, createOrder, arg.UserName, arg.ProductName, arg.Amount)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.ProductName,
		&i.Amount,
	)
	return i, err
}

const getManyOrders = `-- name: GetManyOrders :many
SELECT id, user_name, product_name, amount FROM orders
ORDER BY id
LIMIT $1
OFFSET $2
`

type GetManyOrdersParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) GetManyOrders(ctx context.Context, arg GetManyOrdersParams) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, getManyOrders, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.UserName,
			&i.ProductName,
			&i.Amount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOneOrder = `-- name: GetOneOrder :one
SELECT id, user_name, product_name, amount FROM orders
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetOneOrder(ctx context.Context, id int32) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOneOrder, id)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.ProductName,
		&i.Amount,
	)
	return i, err
}
