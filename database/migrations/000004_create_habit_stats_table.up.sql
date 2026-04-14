CREATE TABLE habit_stats (
    id SERIAL PRIMARY KEY,
    habit_id INT REFERENCES habits(id),
    success_rate FLOAT,
    streak INT,
    last_updated TIMESTAMP
);