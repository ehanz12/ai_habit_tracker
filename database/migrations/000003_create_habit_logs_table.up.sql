CREATE TABLE habit_logs (
    id SERIAL PRIMARY KEY,
    habit_id INT REFERENCES habits(id) ON DELETE CASCADE,
    date DATE,
    completed BOOLEAN DEFAULT FALSE,
    note TEXT
);