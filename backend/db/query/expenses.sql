-- name: CreateExpenses :one
INSERT INTO expenses (
    name,
    amount,
    user_id,
    group_code
) VALUES (
$1, $2, $3 ,$4
) RETURNING *;

-- name: GetExpense :one
SELECT 
    e.name AS expense_name,
    e.amount AS expense_amount,
    u.username AS payer_username,
    g.name AS group_name,
    s.share_type,
    s.share_value,
    u2.username AS shared_with_username
FROM 
    expenses e
JOIN 
    users u ON e.user_id = u.id
JOIN 
    e_group g ON e.group_code = g.code
JOIN 
    share s ON e.id = s.expense_id
JOIN 
    users u2 ON s.user_id = u2.id
WHERE 
    e.id = $1;


-- name: ListExpenses :many
SELECT 
    e.id AS expense_id,
    e.name AS expense_name,
    e.amount AS expense_amount,
    e.user_id AS payer_user_id,
    s.user_id AS sharing_user_id,
    s.share_type,
    s.share_value,
    e.created_at,
    e.updated_at
FROM 
    expenses AS e
LEFT JOIN 
    share AS s ON e.id = s.expense_id;