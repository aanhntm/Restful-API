-- name: CreateOrder :one
INSERT INTO orders (
  user_name,
  product_name,
  amount
) VALUES (
  $1, $2, $3
)RETURNING *;

-- name: GetOneOrder :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;

-- name: GetManyOrders :many
SELECT * FROM orders
ORDER BY id
LIMIT $1
OFFSET $2;