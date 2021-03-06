CREATE TABLE settings (
  -- this is a singleton table, so no auto-incrementing keys
  id integer PRIMARY KEY,
  sheet varchar NOT NULL DEFAULT '',
  script varchar NOT NULL DEFAULT '',
  sync_once varchar NOT NULL DEFAULT '',
  updated_at timestamptz NOT NULL DEFAULT NOW()
);

