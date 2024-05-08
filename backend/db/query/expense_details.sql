-- name: CreateExpenseDetail :one
INSERT INTO expense_details (
    expense_id,
    member_id,
    actual_amount,
    shared_amount
) VALUES (
$1, $2, $3 ,$4 
) RETURNING *;

-- name: GetExpenseDetail :one
SELECT * FROM expense_details WHERE expense_id = $1;

-- name: UpdateExpenseDetail :one
UPDATE expense_details 
SET actual_amount = $3, 
shared_amount = $4 
WHERE expense_id = $1 
AND member_id = $2
RETURNING *;