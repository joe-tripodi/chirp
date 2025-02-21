-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, hashed_password)
VALUES (
  gen_random_uuid(),
  NOW(),
  NOW(),
  $1,
  $2
)
RETURNING *;

-- name: DeleteAllUsers :exec
DELETE FROM users;

-- name: GetUserByEmail :one
SELECT * from users
WHERE email = $1
LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET
  hashed_password = $1,
  email = $2,
  updated_at = NOW()
WHERE id = $3
RETURNING *;


-- name: GetUserById :one
SELECT EXISTS (SELECT 1 FROM users WHERE id = $1) AS user_exists;

-- name: UpgradeUserToRed :one
UPDATE users
SET
  is_chirpy_red = TRUE,
  updated_at = NOW()
WHERE $1 = id
RETURNING *;
