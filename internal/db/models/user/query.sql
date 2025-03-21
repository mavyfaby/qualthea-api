-- name: GetUserByID :one
SELECT * FROM users WHERE id = ? LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = ? LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users;

-- name: CreateUser :execresult
INSERT INTO users (first_name, last_name, birth_date, email, username, password) VALUES (?, ?, ?, ?, ?, ?);

-- name: UpdateUser :execresult
UPDATE users SET first_name = ?, last_name = ?, birth_date = ?, email = ?, username = ?, password = ? WHERE id = ?;

-- name: DeleteUser :execresult
DELETE FROM users WHERE id = ?;
