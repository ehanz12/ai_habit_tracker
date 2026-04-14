CREATE TABLE ai_recommendations (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    habit_id INT REFERENCES habits(id),
    message TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);