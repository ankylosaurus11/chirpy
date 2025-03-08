-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, hashed_password)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING id, created_at, updated_at, email;

-- name: ResetUsers :exec
DELETE FROM users;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: UpdatePassword :one
UPDATE users
SET updated_at = $1, email = $2, hashed_password = $3
WHERE id = $4
RETURNING id, updated_at, email;

-- name: UpgradeRed :exec
UPDATE users
SET is_chirpy_red = $1
WHERE id = $2;