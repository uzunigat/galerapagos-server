CREATE TABLE IF NOT EXISTS weather_card (
    gid VARCHAR UNIQUE NOT NULL CHECK (gid <> ''),
    name VARCHAR NOT NULL,
    water_level INTEGER NOT NULL CHECK (water_level >= 0),
    is_final_game BOOLEAN NOT NULL DEFAULT FALSE,
    quantity INTEGER NOT NULL DEFAULT 3,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

INSERT INTO
    weather_card (
        gid,
        name,
        water_level,
        is_final_game,
        quantity,
        created_at
    )
VALUES
    ('1', 'Sunny', 0, FALSE, 3, NOW()),
    ('2', 'Windy', 1, FALSE, 3, NOW()),
    ('3', 'Rainy', 2, FALSE, 3, NOW()),
    ('4', 'Stormy', 3, FALSE, 2, NOW()),
    ('5', 'Hurricane', 2, TRUE, 1, NOW());