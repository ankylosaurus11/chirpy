-- name: CreateToken :one
INSERT INTO refresh_tokens (token, created_at, updated_at, user_id, expires_at, revoked_at)
VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3,
    NULL
)
RETURNING *;

-- name: GetToken :one
SELECT * FROM refresh_tokens WHERE token = $1;

-- name: UpdateToken :exec
UPDATE refresh_tokens
SET updated_at = $1, revoked_at = $2
WHERE token = $3;