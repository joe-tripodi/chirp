--  token TEXT PRIMARY KEY,
--  created_at TIMESTAMP NOT NULL,
--  updated_at TIMESTAMP NOT NULL,
--  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
--  expires_at TIMESTAMP NOT NULL,
--  revoked_at TIMESTAMP

-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (token, created_at, updated_at, user_id, expires_at, revoked_at)
VALUES (
  $1,
  NOW(),
  NOW(),
  $2,
  NOW() + INTERVAL '60 days',
  NULL
)
RETURNING *;

-- name: RevokeRefreshToken :exec
UPDATE refresh_tokens
SET 
  revoked_at = NOW(),
  updated_at = NOW()
WHERE token = $1;

-- name: GetRefreshToken :one
SELECT * FROM refresh_tokens
WHERE $1 = token
AND revoked_at IS NULL
AND expires_at > NOW()
LIMIT 1;

