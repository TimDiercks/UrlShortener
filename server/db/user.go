package db

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	Id        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Role      string    `db:"role"`
	UpdateAt  time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}

func (d *DB) InsertUser(args CreateShortUrlParams) (ShortUrl, error) {
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

/* TODO:
type UpdateUserParams struct {
	Id       uuid.UUID
	ShortUrl string
	FullUrl  string
}

func (d *DB) UpdateUser(args UpdateUserParams) (ShortUrl, error) {
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

func (d *DB) DeleteUserById(id uuid.UUID) error {
	const q = `
DELETE
FROM urls
WHERE id = $1
	`

	_, err := d.db.Exec(q, id)
	return err
}

func (d *DB) GetUserById(ShortUrl string) (string, error) {
	const q = `
SELECT full_url
FROM urls
WHERE short_url = $1
	`

	row := d.db.QueryRow(q, ShortUrl)
	var fullUrl string
	return fullUrl, row.Scan(&fullUrl)
}*/

func (d *DB) GetUserByApiKey(key string) (User, error) {
	const q = `
SELECT u.*
FROM users u 
INNER JOIN apikeys a ON u.id = a.user_id
WHERE a.key = $1
	`

	row := d.db.QueryRowx(q, key)
	var user User
	return user, row.StructScan(&user)
}
