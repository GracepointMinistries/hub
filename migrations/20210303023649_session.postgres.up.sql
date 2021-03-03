CREATE TABLE sessions (
  id serial PRIMARY KEY,
  ip varchar NOT NULL,
  user_id integer NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW(),
  CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE admin_sessions (
  id serial PRIMARY KEY,
  email varchar NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW()
);

