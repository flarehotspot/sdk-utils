-- name: CreateWalletTrns :one
INSERT INTO wallet_transactions (
  wallet_id, amount, new_balance, description
) 
VALUES 
  ($1, $2, $3, $4)
RETURNING *;


-- name: FindWalletTrns :one
SELECT 
  id, 
  wallet_id, 
  amount, 
  new_balance, 
  description, 
  created_at 
FROM 
  wallet_transactions 
WHERE 
  id = $1 
LIMIT 
  1;

