CREATE TABLE IF NOT EXISTS wreck_card (
    gid VARCHAR UNIQUE NOT NULL CHECK (gid <> ''),
    name VARCHAR NOT NULL,
    description TEXT NOT NULL,
    type VARCHAR NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    is_discarded BOOLEAN NOT NULL DEFAULT TRUE,
    quantity INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (gid)
);

INSERT INTO
    wreck_card (
        gid,
        name,
        description,
        type,
        is_active,
        is_discarded,
        quantity
    )
VALUES
    (
        '1',
        'Axe',
        'Each turn, you can pick 2 pieces of wood without any risk if you choose wood action',
        'PERMANENT',
        TRUE,
        FALSE,
        1
    ),
    (
        '2',
        'Cat',
        'In case of shortage ONLY, it gives you 2 pieces of food',
        'SPECIAL',
        TRUE,
        TRUE,
        1
    ),
    (
        '3',
        'Sandwich',
        'Equivalent to 1 ration of food',
        'RESSOURCE',
        TRUE,
        TRUE,
        2
    ),
    (
        '4',
        'Taser',
        'Allows you to steal a permanent card who is in front of another shipwreck ',
        'SPECIAL',
        TRUE,
        TRUE,
        1
    ),
    (
        '5',
        'Conch',
        'During that turn, I`m the boss and no one can vote against me. This card can be played before or after a vote',
        'SPECIAL',
        TRUE,
        TRUE,
        1
    ),
    (
        '6',
        'Pendulum',
        'Allows me to impose an action for another player (fishing, search for some water or wood, take a new wreck card)',
        'SPECIAL',
        TRUE,
        TRUE,
        1
    ),
    (
        '7',
        'Club',
        'Gives 2 votes in each vote',
        'PERMANENT',
        TRUE,
        FALSE,
        1
    ),
    (
        '8',
        'Concave Plate',
        'If someone shots you, the shot ricochets off my neighbor on the left',
        'SPECIAL',
        TRUE,
        TRUE,
        1
    ),
    (
        '9',
        'Bottle of water',
        'Equivalent to 1 ration of water',
        'RESSOURCE',
        TRUE,
        TRUE,
        1
    ),
    (
        '10',
        'Whetstone',
        'Linked to the axe. Allows me to kill one of my neighbor. The axe is discarded after the use of this card',
        'SPECIAL',
        TRUE,
        TRUE,
        1
    );

CREATE TRIGGER wreck_card_updated_at BEFORE
UPDATE
    ON wreck_card FOR EACH ROW EXECUTE PROCEDURE on_update_timestamp();