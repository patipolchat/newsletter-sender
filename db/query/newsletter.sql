-- name: CreateNewsLetter :one
-- description: Create a new newsletter
-- input: header: String, body: String
-- output: Newsletter: Object
-- example: CreateNewsLetter("Newsletter 1", "This is the first newsletter")
INSERT INTO newsletters (header, body)
VALUES ($1, $2)
RETURNING *;
