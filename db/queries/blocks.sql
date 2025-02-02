-- name: ListBlocks :many
SELECT * FROM blocks
ORDER BY id;

-- name: BlockByID :one
SELECT * FROM blocks
WHERE id = ?
LIMIT 1;

-- name: BlockByTimestamp :one
SELECT * FROM blocks
WHERE timestamp = ?
LIMIT 1;

-- name: BlockByRootHash :one
SELECT * FROM blocks
WHERE root_hash = ?
LIMIT 1;

-- name: CreateBlock :execresult
INSERT INTO blocks (
    id,
    timestamp,
    root_hash
)
VALUES (
    ?, ?, ?
);

-- name: DeleteBlock :execresult
DELETE FROM blocks
WHERE id = ?;

-- name: UpdateBlock :execresult
UPDATE blocks
SET timestamp = sqlc.arg(timestamp),
    root_hash = sqlc.arg(root_hash)
WHERE id = sqlc.arg(id);

-- name: LastBlock :one
SELECT id FROM blocks 
ORDER BY id DESC 
LIMIT 1;