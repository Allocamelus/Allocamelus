-- migrate:up
CREATE INDEX Users_userName_idx ON Users (lower(userName) varchar_pattern_ops);
CREATE INDEX Users_email_idx ON Users (lower(email) varchar_pattern_ops);

-- migrate:down

DROP INDEX Users_userName_idx ON Users;
DROP INDEX Users_email_idx ON Users;