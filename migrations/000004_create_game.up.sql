CREATE TABLE IF NOT EXISTS game
(
    gid                    VARCHAR NOT NULL CHECK (gid <> ''),
    status                 VARCHAR(255) NOT NULL CHECK (status <> ''),
    created_at             TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at             TIMESTAMP NOT NULL DEFAULT NOW()
);

 CREATE TRIGGER game_updated_at
    BEFORE UPDATE ON game
    FOR EACH ROW
    EXECUTE PROCEDURE on_update_timestamp();
