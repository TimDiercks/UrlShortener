package db

import (
	"github.com/gofrs/uuid"
)

type ShortUrl struct {
	Id       uuid.UUID `db:"id"`
	UserId   uuid.UUID `db:"user_id"`
	ShortUrl string    `db:"short_url"`
	FullUrl  string    `db:"full_url"`
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
