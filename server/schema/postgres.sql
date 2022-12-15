CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name text,
    email text NOT NULL,
    password text NOT NULL,
    role text NOT NULL DEFAULT 'user',
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    UNIQUE(email)
);

CREATE TABLE IF NOT EXISTS urls (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid,
    short_url text NOT NULL,
    full_url text NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    UNIQUE(short_url)
);

CREATE TABLE IF NOT EXISTS apikeys (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    key text NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    UNIQUE(user_id)
);

/* TESTDATA */
INSERT INTO users(id, name, email, password, role)
VALUES ('a5b8292d-2bfd-4859-b7a1-4c3f4306103b', 'admin', 'admin@example.com', '<HASHEDPASSWORD>', 'admin');

INSERT INTO apikeys(user_id, key)
VALUES ('a5b8292d-2bfd-4859-b7a1-4c3f4306103b', 'b890b3d18a222c50de9089c74a092de471b23dc69698e96402b6269a45c7878e');

INSERT INTO urls(user_id, short_url, full_url)
VALUES ('a5b8292d-2bfd-4859-b7a1-4c3f4306103b', 'gh', 'https://github.com');
INSERT INTO urls(user_id, short_url, full_url)
VALUES ('a5b8292d-2bfd-4859-b7a1-4c3f4306103b', 'goo', 'https://google.com');