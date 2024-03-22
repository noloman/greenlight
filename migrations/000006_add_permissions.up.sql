CREATE TABLE IF NOT EXISTS permissions (id bigserial PRIMARY KEY, code TEXT NOT NULL);
CREATE TABLE IF NOT EXISTS users_permissions(
  user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
  permission_id bigint NOT NULL REFERENCES permissions ON DELETE CASCADE
);
INSERT INTO permissions (code)
VALUES ('movies:read'), ('movies:write');
