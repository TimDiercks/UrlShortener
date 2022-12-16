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

// TODO: update user functionality

func (d *DB) DeleteUserById(id uuid.UUID) error {
	const q = `
DELETE
FROM users
WHERE id = $1
	`

	_, err := d.db.Exec(q, id)
	return err
}

func (d *DB) GetUserById(userId uuid.UUID) (User, error) {
	const q = `
SELECT *
FROM users
WHERE id = $1
	`

	row := d.db.QueryRowx(q, userId)
	var user User
	return user, row.StructScan(&user)
}

func (d *DB) GetShortUrlsByUserId(userId uuid.UUID) ([]ShortUrl, error) {
	const q = `
SELECT u.*
FROM urls u 
WHERE u.id = $1
	`

	row := d.db.QueryRowx(q, userId)
	var urls []ShortUrl
	return urls, row.StructScan(&urls)
}

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
