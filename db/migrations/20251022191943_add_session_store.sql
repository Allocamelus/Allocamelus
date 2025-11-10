-- migrate:up
CREATE TABLE Sessions (
  key VARCHAR(64) PRIMARY KEY,
  data BYTEA NOT NULL,
  expiration BIGINT NOT NULL -- Time
);

CREATE INDEX Sessions_created_idx ON Sessions (expiration);

-- migrate:down

DROP TABLE IF EXISTS Sessions;


