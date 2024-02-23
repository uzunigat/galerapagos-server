CREATE TABLE IF NOT EXISTS player_game_relation
(
    gid                    VARCHAR UNIQUE NOT NULL CHECK (gid <> ''),
    player_gid             VARCHAR NOT NULL REFERENCES player(gid) ON DELETE CASCADE,
    game_gid               VARCHAR NOT NULL REFERENCES game(gid) ON DELETE CASCADE,
    created_at             TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at             TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE (player_gid, game_gid)
);

CREATE TRIGGER player_game_relation_updated_at BEFORE
UPDATE ON player_game_relation FOR EACH ROW EXECUTE PROCEDURE on_update_timestamp();