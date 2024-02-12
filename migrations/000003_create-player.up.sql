CREATE TABLE IF NOT EXISTS player
(
    gid                    VARCHAR NOT NULL CHECK (gid <> ''),
    givenName              VARCHAR(255) NOT NULL CHECK (givenName <> ''),
    familyName             VARCHAR(128) NOT NULL CHECK (familyName <> ''),
    email                  VARCHAR(128) NOT NULL CHECK (email <> ''),
    password               VARCHAR(128) NOT NULL CHECK (password <> ''),
    created_at             TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at             TIMESTAMP NOT NULL DEFAULT NOW()
);

 CREATE TRIGGER player_updated_at
    BEFORE UPDATE ON player
    FOR EACH ROW
    EXECUTE PROCEDURE on_update_timestamp();
