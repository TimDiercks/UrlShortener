CREATE TABLE IF NOT EXISTS user (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name text,
    email text NOT NULL,
    password text NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    UNIQUE(email)
);

CREATE TABLE IF NOT EXISTS urls (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid NOT NULL REFERENCES user(id) ON DELETE CASCADE ON UPDATE CASCADE,
    shortened text NOT NULL,
    full text NOT NULL
);

CREATE TABLE IF NOT EXISTS apikeys (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid NOT NULL REFERENCES user(id) ON DELETE CASCADE ON UPDATE CASCADE,
    key text NOT NULL
);