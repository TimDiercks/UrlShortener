package db

import (
	"time"

	"github.com/gofrs/uuid"
)

type ShortUrl struct {
	Id        uuid.UUID `db:"id"`
	UserId    uuid.UUID `db:"user_id"`
	ShortUrl  string    `db:"short_url"`
	FullUrl   string    `db:"full_url"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type CreateShortUrlParams struct {
	UserId   uuid.UUID
	ShortUrl string
	FullUrl  string
}

func (d *DB) InsertShortUrl(args CreateShortUrlParams) (ShortUrl, error) {
	const q = `
INSERT INTO urls(
		user_id,
		short_url,
		full_url
) VALUES(
		$1,
		$2,
		$3
)
RETURNING *
	`

	row := d.db.QueryRowx(q, args.UserId, args.ShortUrl, args.FullUrl)
	var shortUrl ShortUrl
	return shortUrl, row.StructScan(&shortUrl)
}

type UpdateShortUrlParams struct {
	Id       uuid.UUID
	ShortUrl string
	FullUrl  string
}

func (d *DB) UpdateShortUrl(args UpdateShortUrlParams) (ShortUrl, error) {
	const q = `
UPDATE urls 
SET shorturl=$1,
	fullurl=$2,
	updated_at=now()
WHERE id=$3
RETURNING *
	`

	row := d.db.QueryRowx(q, args.ShortUrl, args.FullUrl, args.Id)
	var shortUrl ShortUrl
	return shortUrl, row.StructScan(&shortUrl)
}

func (d *DB) DeleteShortUrlById(id uuid.UUID) error {
	const q = `
DELETE
FROM urls
WHERE id = $1
	`

	_, err := d.db.Exec(q, id)
	return err
}

func (d *DB) GetFullUrlFromShortUrl(ShortUrl string) (string, error) {
	const q = `
SELECT full_url
FROM urls
WHERE short_url = $1
	`

	row := d.db.QueryRow(q, ShortUrl)
	var fullUrl string
	return fullUrl, row.Scan(&fullUrl)
}
