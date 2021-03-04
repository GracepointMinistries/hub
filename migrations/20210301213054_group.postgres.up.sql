CREATE TABLE GROUPS (
  id serial PRIMARY KEY,
  name varchar UNIQUE NOT NULL,
  zoom_link varchar NOT NULL DEFAULT '', -- this is a dirty hack to avoid nullable go structs
  published boolean NOT NULL DEFAULT FALSE,
  archived boolean NOT NULL DEFAULT FALSE,
  updated_at timestamptz NOT NULL DEFAULT NOW(),
  created_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE group_members (
  group_id integer NOT NULL,
  user_id integer NOT NULL,
  CONSTRAINT fk_group FOREIGN KEY (group_id) REFERENCES GROUPS (id),
  CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id),
  PRIMARY KEY (group_id, user_id)
);

