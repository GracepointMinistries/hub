CREATE TABLE oauths (
  id serial PRIMARY KEY,
  provider varchar NOT NULL,
  provider_id varchar NOT NULL,
  updated_at timestamptz NOT NULL DEFAULT NOW(),
  created_at timestamptz NOT NULL DEFAULT NOW(),
  UNIQUE (provider, provider_id)
);

CREATE TABLE users (
  id serial PRIMARY KEY,
  name varchar NOT NULL,
  email varchar UNIQUE NOT NULL,
  blocked boolean NOT NULL DEFAULT FALSE,
  updated_at timestamptz NOT NULL DEFAULT NOW(),
  created_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE oauth_users (
  oauth_id integer NOT NULL,
  user_id integer NOT NULL,
  CONSTRAINT fk_oauth FOREIGN KEY (oauth_id) REFERENCES oauths (id),
  CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id),
  PRIMARY KEY (oauth_id, user_id)
);

