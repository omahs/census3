// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: censuses.sql

package queries

import (
	"context"
	"database/sql"

	"github.com/vocdoni/census3/db/annotations"
)

const censusByID = `-- name: CensusByID :one
SELECT id, strategy_id, merkle_root, uri FROM Censuses
WHERE id = ?
LIMIT 1
`

func (q *Queries) CensusByID(ctx context.Context, id int64) (Censuses, error) {
	row := q.db.QueryRowContext(ctx, censusByID, id)
	var i Censuses
	err := row.Scan(
		&i.ID,
		&i.StrategyID,
		&i.MerkleRoot,
		&i.Uri,
	)
	return i, err
}

const censusByMerkleRoot = `-- name: CensusByMerkleRoot :one
SELECT id, strategy_id, merkle_root, uri FROM Censuses
WHERE merkle_root = ?
LIMIT 1
`

func (q *Queries) CensusByMerkleRoot(ctx context.Context, merkleRoot annotations.Hash) (Censuses, error) {
	row := q.db.QueryRowContext(ctx, censusByMerkleRoot, merkleRoot)
	var i Censuses
	err := row.Scan(
		&i.ID,
		&i.StrategyID,
		&i.MerkleRoot,
		&i.Uri,
	)
	return i, err
}

const censusByStrategyID = `-- name: CensusByStrategyID :many
SELECT id, strategy_id, merkle_root, uri FROM Censuses
WHERE strategy_id = ?
`

func (q *Queries) CensusByStrategyID(ctx context.Context, strategyID int64) ([]Censuses, error) {
	rows, err := q.db.QueryContext(ctx, censusByStrategyID, strategyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Censuses
	for rows.Next() {
		var i Censuses
		if err := rows.Scan(
			&i.ID,
			&i.StrategyID,
			&i.MerkleRoot,
			&i.Uri,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const censusByURI = `-- name: CensusByURI :one
SELECT id, strategy_id, merkle_root, uri FROM Censuses
WHERE uri = ?
LIMIT 1
`

func (q *Queries) CensusByURI(ctx context.Context, uri sql.NullString) (Censuses, error) {
	row := q.db.QueryRowContext(ctx, censusByURI, uri)
	var i Censuses
	err := row.Scan(
		&i.ID,
		&i.StrategyID,
		&i.MerkleRoot,
		&i.Uri,
	)
	return i, err
}

const censusesByStrategyIDAndBlockID = `-- name: CensusesByStrategyIDAndBlockID :many
SELECT c.id, c.strategy_id, c.merkle_root, c.uri FROM Censuses c
JOIN CensusBlocks cb ON c.id = cb.census_id
WHERE c.strategy_id = ? AND cb.block_id = ?
LIMIT ? OFFSET ?
`

type CensusesByStrategyIDAndBlockIDParams struct {
	StrategyID int64
	BlockID    int64
	Limit      int32
	Offset     int32
}

func (q *Queries) CensusesByStrategyIDAndBlockID(ctx context.Context, arg CensusesByStrategyIDAndBlockIDParams) ([]Censuses, error) {
	rows, err := q.db.QueryContext(ctx, censusesByStrategyIDAndBlockID,
		arg.StrategyID,
		arg.BlockID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Censuses
	for rows.Next() {
		var i Censuses
		if err := rows.Scan(
			&i.ID,
			&i.StrategyID,
			&i.MerkleRoot,
			&i.Uri,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const censusesByTokenID = `-- name: CensusesByTokenID :many
SELECT c.id, c.strategy_id, c.merkle_root, c.uri FROM Censuses AS c
JOIN StrategyTokens AS st ON c.strategy_id = st.strategy_id
WHERE st.token_id = ?
LIMIT ? OFFSET ?
`

type CensusesByTokenIDParams struct {
	TokenID annotations.Address
	Limit   int32
	Offset  int32
}

func (q *Queries) CensusesByTokenID(ctx context.Context, arg CensusesByTokenIDParams) ([]Censuses, error) {
	rows, err := q.db.QueryContext(ctx, censusesByTokenID, arg.TokenID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Censuses
	for rows.Next() {
		var i Censuses
		if err := rows.Scan(
			&i.ID,
			&i.StrategyID,
			&i.MerkleRoot,
			&i.Uri,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const createCensus = `-- name: CreateCensus :execresult
INSERT INTO Censuses (
    id,
    strategy_id,
    merkle_root,
    uri
)
VALUES (
    ?, ?, ?, ?
)
`

type CreateCensusParams struct {
	ID         int64
	StrategyID int64
	MerkleRoot annotations.Hash
	Uri        sql.NullString
}

func (q *Queries) CreateCensus(ctx context.Context, arg CreateCensusParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createCensus,
		arg.ID,
		arg.StrategyID,
		arg.MerkleRoot,
		arg.Uri,
	)
}

const createCensusBlock = `-- name: CreateCensusBlock :execresult
INSERT INTO CensusBlocks (
    census_id,
    block_id
)
VALUES (
    ?, ?
)
`

type CreateCensusBlockParams struct {
	CensusID int64
	BlockID  int64
}

func (q *Queries) CreateCensusBlock(ctx context.Context, arg CreateCensusBlockParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createCensusBlock, arg.CensusID, arg.BlockID)
}

const deleteCensus = `-- name: DeleteCensus :execresult
DELETE FROM Censuses
WHERE id = ?
`

func (q *Queries) DeleteCensus(ctx context.Context, id int64) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteCensus, id)
}

const deleteCensusBlock = `-- name: DeleteCensusBlock :execresult
DELETE FROM CensusBlocks
WHERE census_id = ? AND block_id = ?
`

type DeleteCensusBlockParams struct {
	CensusID int64
	BlockID  int64
}

func (q *Queries) DeleteCensusBlock(ctx context.Context, arg DeleteCensusBlockParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteCensusBlock, arg.CensusID, arg.BlockID)
}

const lastCensusID = `-- name: LastCensusID :one
SELECT strategy_id 
FROM Censuses 
ORDER BY strategy_id DESC
LIMIT 1
`

func (q *Queries) LastCensusID(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, lastCensusID)
	var strategy_id int64
	err := row.Scan(&strategy_id)
	return strategy_id, err
}

const paginatedCensuses = `-- name: PaginatedCensuses :many
SELECT id, strategy_id, merkle_root, uri FROM Censuses
ORDER BY id
LIMIT ? OFFSET ?
`

type PaginatedCensusesParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) PaginatedCensuses(ctx context.Context, arg PaginatedCensusesParams) ([]Censuses, error) {
	rows, err := q.db.QueryContext(ctx, paginatedCensuses, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Censuses
	for rows.Next() {
		var i Censuses
		if err := rows.Scan(
			&i.ID,
			&i.StrategyID,
			&i.MerkleRoot,
			&i.Uri,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCensus = `-- name: UpdateCensus :execresult
UPDATE Censuses
SET strategy_id = ?,
    merkle_root = ?,
    uri = ?
WHERE id = ?
`

type UpdateCensusParams struct {
	StrategyID int64
	MerkleRoot annotations.Hash
	Uri        sql.NullString
	ID         int64
}

func (q *Queries) UpdateCensus(ctx context.Context, arg UpdateCensusParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateCensus,
		arg.StrategyID,
		arg.MerkleRoot,
		arg.Uri,
		arg.ID,
	)
}

const updateCensusBlock = `-- name: UpdateCensusBlock :execresult
UPDATE CensusBlocks
SET census_id = ?,
    block_id = ?
WHERE census_id = ? AND block_id = ?
`

type UpdateCensusBlockParams struct {
	CensusID int64
	BlockID  int64
}

func (q *Queries) UpdateCensusBlock(ctx context.Context, arg UpdateCensusBlockParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateCensusBlock,
		arg.CensusID,
		arg.BlockID,
		arg.CensusID,
		arg.BlockID,
	)
}
