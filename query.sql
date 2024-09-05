-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT id, username, email, profile FROM users ORDER BY username; 

-- name: CreateUser :exec
INSERT INTO users (username,password_hash, email, profile )
VALUES($1, $2, $3, $4)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
    set 
      username = $2,
      email = $3,
      profile = $4
WHERE id=$1;

-- name: DeleteUser :exec 
DELETE FROM users
WHERE id = $1;