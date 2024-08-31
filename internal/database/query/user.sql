-- name: GetUsers :many
SELECT id, first_name, last_name, email, phone, address, created_at, updated_at
FROM "user"
order by id
limit $1
offset $2 ;

-- name: GetUserById :one
SELECT id, first_name, last_name, email, phone, address, created_at, updated_at
FROM "user"
WHERE id = $1
LIMIT 1;

-- name: GetUserByEmailOrPassword :one
SELECT id, first_name, last_name, email, phone, address, created_at, updated_at
FROM "user"
WHERE id = $1 or email = $2 or phone = $3
LIMIT 1;

-- name: GetUserByEmailAndPassword :one
SELECT id, first_name, last_name, email, phone, address, created_at, updated_at
FROM "user"
WHERE email = $1
  AND password_hash = $2
LIMIT 1;

-- name: CreateUser :one
INSERT INTO "user" (first_name, last_name, email, password_hash, phone, address)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;


-- name: UpdateUser :one
UPDATE "user"
SET first_name = $2,
    last_name  = $3,
    email = $4,
    phone = $5,
    address = $6,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: UpdateUserName :one
UPDATE "user"
SET first_name = $2,
    last_name  = $3,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: UpdateUserEmail :one
UPDATE "user"
SET email      = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: UpdateUserPhone :one
UPDATE "user"
SET phone      = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: UpdateUserAddress :one
UPDATE "user"
SET address    = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: UpdateUserPassword :one
UPDATE "user"
SET password_hash = $2,
    updated_at    = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE
FROM "user"
WHERE id = $1;
