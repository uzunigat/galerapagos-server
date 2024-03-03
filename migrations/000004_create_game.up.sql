CREATE TABLE IF NOT EXISTS game (
    gid VARCHAR UNIQUE NOT NULL CHECK (gid <> ''),
    status VARCHAR(255) NOT NULL CHECK (status <> ''),
    raft_level INTEGER NOT NULL DEFAULT 0 CHECK (raft_level >= 0),
    water_resources INTEGER CHECK (water_resources >= 0),
    food_resources INTEGER CHECK (food_resources >= 0),
    weather_cards JSONB DEFAULT '[]',
    wreck_card_gids JSONB DEFAULT '[]',
    player_turns JSONB DEFAULT '[]',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TRIGGER game_updated_at BEFORE
UPDATE
    ON game FOR EACH ROW EXECUTE PROCEDURE on_update_timestamp();