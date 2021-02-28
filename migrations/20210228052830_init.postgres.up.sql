CREATE TABLE oauths (
	id serial PRIMARY KEY,
	provider varchar NOT NULL,
	provider_id varchar NOT NULL,
	created_at timestamptz NOT NULL DEFAULT NOW(),
	UNIQUE(provider, provider_id)
);

CREATE TABLE users (
	id serial PRIMARY KEY,
	email varchar UNIQUE NOT NULL,
	created_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE oauth_users (
	oauth_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	CONSTRAINT fk_oauth FOREIGN KEY(oauth_id) REFERENCES oauths(id),
	CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
	PRIMARY KEY(oauth_id, user_id)
);
