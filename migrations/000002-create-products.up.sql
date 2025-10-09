-- +migrate Up
CREATE TABLE leaderboard (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    game_id VARCHAR(100) NOT NULL,
    score INT NOT NULL,
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);
