-- name: GetRegistryByRegistrationID :one
SELECT * FROM registry
WHERE registration_id = $1 LIMIT 1;

-- name: CreateRegistry :one
INSERT INTO registry (
    registration_id, isbn, title, author, translator, print_amount, self_publish, partner, registration_date
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;
