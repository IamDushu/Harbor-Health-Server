-- name: CreateUser :one 
INSERT INTO users (
    user_id,
    email, 
    first_name, 
    last_name,
    phone_number,
    image_url
) VALUES (
   $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: UpdateUser :one 
UPDATE users 
SET first_name = $1, last_name = $2, phone_number = $3, is_onboarded = $4
WHERE email = $5
RETURNING *;