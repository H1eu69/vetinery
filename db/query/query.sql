-- name: CreateVetineryBill :one
INSERT INTO Veterinary_bill (
 id, username , price , created_at 
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetAllVetineryBill :many
SELECT * FROM Veterinary_bill;