-- queries.sql
-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: CreateUser :exec
INSERT INTO users (name, phone_number) VALUES ($1, $2);

-- name: UpdateOtp :exec
UPDATE users SET otp = $2, otp_expiry_time = $3
WHERE id = $1;

-- name: GetUserByPhone :one
SELECT * FROM users WHERE phone_number = $1;