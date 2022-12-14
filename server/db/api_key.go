package db

import (
	"time"

	"github.com/gofrs/uuid"
)

type ApiKey struct {
	Id        uuid.UUID `db:"id"`
	UserId    uuid.UUID `db:"user_id"`
	Key       string    `db:"key"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type CreateKeyParams struct {
	UserId uuid.UUID
	Key    string
}

func (d *DB) InsertKey(args CreateKeyParams) (ApiKey, error) {
	const q = `
INSERT INTO apiKeys(
		user_id,
		key
) VALUES(
		$1,
		$2
)
RETURNING *
	`

	row := d.db.QueryRowx(q, args.UserId, args.Key)
	var apiKey ApiKey
	return apiKey, row.StructScan(&apiKey)
}

type UpdateApiKeyParams struct {
	Id  uuid.UUID
	Key string
}

func (d *DB) UpdateApiKey(args UpdateApiKeyParams) (ApiKey, error) {
	const q = `
UPDATE apikeys 
SET key=$1,
	updated_at=now()
WHERE id=$2
RETURNING *
	`

	row := d.db.QueryRowx(q, args.Key, args.Id)
	var apiKey ApiKey
	return apiKey, row.StructScan(&apiKey)
}

func (d *DB) DeleteApiKeyById(id uuid.UUID) error {
	const q = `
DELETE
FROM apikeys
WHERE id = $1
	`

	_, err := d.db.Exec(q, id)
	return err
}

func (d *DB) VerifyToken(key string) error {
	const q = `
SELECT id
FROM apikeys
WHERE key = $1
	`

	row := d.db.QueryRow(q, key)
	err := row.Err()
	return err
}
