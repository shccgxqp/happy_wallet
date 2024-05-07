-- name: CreateShare :one
INSERT INTO share (
    expense_id,
    user_id,
    share_type,
    share_value
) VALUES (
$1, $2, $3 ,$4
) RETURNING *;