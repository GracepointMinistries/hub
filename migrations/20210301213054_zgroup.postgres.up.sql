CREATE TABLE zgroups (
  id serial PRIMARY KEY,
  name varchar UNIQUE NOT NULL,
  zoom_link varchar NOT NULL DEFAULT '', -- this is a dirty hack to avoid nullable go structs
  published boolean NOT NULL DEFAULT FALSE,
  archived boolean NOT NULL DEFAULT FALSE,
  updated_at timestamptz NOT NULL DEFAULT NOW(),
  created_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE zgroup_members (
  zgroup_id integer NOT NULL,
  user_id integer NOT NULL,
  CONSTRAINT fk_zgroup FOREIGN KEY (zgroup_id) REFERENCES zgroups (id),
  CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id),
  PRIMARY KEY (zgroup_id, user_id)
);

