-- name: CreateMember :one 
INSERT INTO members (
    member_id,
    user_id,
    gender, 
    date_of_birth, 
    insurance,
    address_line_one,
    address_line_two,
    accepted_terms
) VALUES (
   $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetMember :one
SELECT * FROM members
WHERE user_id = $1 LIMIT 1;