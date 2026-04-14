CREATE TABLE habits (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(100),
    description TEXT,
    target_per_day INT DEFAULT 1,
    preferred_time TIME,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);