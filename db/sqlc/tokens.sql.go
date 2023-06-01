// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: tokens.sql

package queries

import (
	"context"
	"database/sql"

	"github.com/vocdoni/census3/db/annotations"
)

const createToken = `-- name: CreateToken :execresult
INSERT INTO tokens (
    id,
    name,
    symbol,
    decimals,
    total_supply,
    creation_block,
    type_id,
    synced
)
VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?
)
`

type CreateTokenParams struct {
	ID            annotations.Address
	Name          sql.NullString
	Symbol        sql.NullString
	Decimals      sql.NullInt64
	TotalSupply   annotations.BigInt
	CreationBlock sql.NullInt32
	TypeID        int64
	Synced        bool
}

func (q *Queries) CreateToken(ctx context.Context, arg CreateTokenParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createToken,
		arg.ID,
		arg.Name,
		arg.Symbol,
		arg.Decimals,
		arg.TotalSupply,
		arg.CreationBlock,
		arg.TypeID,
		arg.Synced,
	)
}

const deleteToken = `-- name: DeleteToken :execresult
DELETE FROM tokens
WHERE id = ?
`

func (q *Queries) DeleteToken(ctx context.Context, id annotations.Address) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteToken, id)
}

const existsToken = `-- name: ExistsToken :one
SELECT EXISTS 
    (SELECT id 
    FROM tokens
    WHERE id = ?)
`

func (q *Queries) ExistsToken(ctx context.Context, id annotations.Address) (bool, error) {
	row := q.db.QueryRowContext(ctx, existsToken, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const listTokens = `-- name: ListTokens :many
SELECT id, name, symbol, decimals, total_supply, creation_block, type_id, synced FROM tokens
ORDER BY type_id, name
`

func (q *Queries) ListTokens(ctx context.Context) ([]Token, error) {
	rows, err := q.db.QueryContext(ctx, listTokens)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Token
	for rows.Next() {
		var i Token
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Symbol,
			&i.Decimals,
			&i.TotalSupply,
			&i.CreationBlock,
			&i.TypeID,
			&i.Synced,
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

const tokenByID = `-- name: TokenByID :one
SELECT id, name, symbol, decimals, total_supply, creation_block, type_id, synced FROM tokens
WHERE id = ?
LIMIT 1
`

func (q *Queries) TokenByID(ctx context.Context, id annotations.Address) (Token, error) {
	row := q.db.QueryRowContext(ctx, tokenByID, id)
	var i Token
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Symbol,
		&i.Decimals,
		&i.TotalSupply,
		&i.CreationBlock,
		&i.TypeID,
		&i.Synced,
	)
	return i, err
}

const tokenByName = `-- name: TokenByName :one
SELECT id, name, symbol, decimals, total_supply, creation_block, type_id, synced FROM tokens
WHERE name = ?
LIMIT 1
`

func (q *Queries) TokenByName(ctx context.Context, name sql.NullString) (Token, error) {
	row := q.db.QueryRowContext(ctx, tokenByName, name)
	var i Token
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Symbol,
		&i.Decimals,
		&i.TotalSupply,
		&i.CreationBlock,
		&i.TypeID,
		&i.Synced,
	)
	return i, err
}

const tokenBySymbol = `-- name: TokenBySymbol :one
SELECT id, name, symbol, decimals, total_supply, creation_block, type_id, synced FROM tokens
WHERE symbol = ?
LIMIT 1
`

func (q *Queries) TokenBySymbol(ctx context.Context, symbol sql.NullString) (Token, error) {
	row := q.db.QueryRowContext(ctx, tokenBySymbol, symbol)
	var i Token
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Symbol,
		&i.Decimals,
		&i.TotalSupply,
		&i.CreationBlock,
		&i.TypeID,
		&i.Synced,
	)
	return i, err
}

const tokensByStrategyID = `-- name: TokensByStrategyID :many
SELECT t.id, t.name, t.symbol, t.decimals, t.total_supply, t.creation_block, t.type_id, t.synced, st.strategy_id, st.token_id, st.min_balance, st.method_hash FROM tokens t
JOIN strategy_tokens st ON st.token_id = t.id
WHERE st.strategy_id = ?
ORDER BY t.name
`

type TokensByStrategyIDRow struct {
	ID            annotations.Address
	Name          sql.NullString
	Symbol        sql.NullString
	Decimals      sql.NullInt64
	TotalSupply   annotations.BigInt
	CreationBlock sql.NullInt32
	TypeID        int64
	Synced        bool
	StrategyID    int64
	TokenID       []byte
	MinBalance    []byte
	MethodHash    []byte
}

func (q *Queries) TokensByStrategyID(ctx context.Context, strategyID int64) ([]TokensByStrategyIDRow, error) {
	rows, err := q.db.QueryContext(ctx, tokensByStrategyID, strategyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TokensByStrategyIDRow
	for rows.Next() {
		var i TokensByStrategyIDRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Symbol,
			&i.Decimals,
			&i.TotalSupply,
			&i.CreationBlock,
			&i.TypeID,
			&i.Synced,
			&i.StrategyID,
			&i.TokenID,
			&i.MinBalance,
			&i.MethodHash,
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

const tokensByType = `-- name: TokensByType :many
SELECT id, name, symbol, decimals, total_supply, creation_block, type_id, synced FROM tokens
WHERE type_id = ?
ORDER BY name
`

func (q *Queries) TokensByType(ctx context.Context, typeID int64) ([]Token, error) {
	rows, err := q.db.QueryContext(ctx, tokensByType, typeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Token
	for rows.Next() {
		var i Token
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Symbol,
			&i.Decimals,
			&i.TotalSupply,
			&i.CreationBlock,
			&i.TypeID,
			&i.Synced,
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

const updateToken = `-- name: UpdateToken :execresult
UPDATE tokens
SET name = ?,
    symbol = ?,
    decimals = ?,
    total_supply = ?,
    creation_block = ?,
    type_id = ?,
    synced = ?
WHERE id = ?
`

type UpdateTokenParams struct {
	Name          sql.NullString
	Symbol        sql.NullString
	Decimals      sql.NullInt64
	TotalSupply   annotations.BigInt
	CreationBlock sql.NullInt32
	TypeID        int64
	Synced        bool
	ID            annotations.Address
}

func (q *Queries) UpdateToken(ctx context.Context, arg UpdateTokenParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateToken,
		arg.Name,
		arg.Symbol,
		arg.Decimals,
		arg.TotalSupply,
		arg.CreationBlock,
		arg.TypeID,
		arg.Synced,
		arg.ID,
	)
}

const updateTokenCreationBlock = `-- name: UpdateTokenCreationBlock :execresult
UPDATE tokens
SET creation_block = ?
WHERE id = ?
`

type UpdateTokenCreationBlockParams struct {
	CreationBlock sql.NullInt32
	ID            annotations.Address
}

func (q *Queries) UpdateTokenCreationBlock(ctx context.Context, arg UpdateTokenCreationBlockParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateTokenCreationBlock, arg.CreationBlock, arg.ID)
}

const updateTokenStatus = `-- name: UpdateTokenStatus :execresult
UPDATE tokens
SET synced = ?
WHERE id = ?
`

type UpdateTokenStatusParams struct {
	Synced bool
	ID     annotations.Address
}

func (q *Queries) UpdateTokenStatus(ctx context.Context, arg UpdateTokenStatusParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateTokenStatus, arg.Synced, arg.ID)
}
