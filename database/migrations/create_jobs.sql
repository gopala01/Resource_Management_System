CREATE TABLE IF NOT EXISTS jobs (
    id UUID PRIMARY KEY,
    status TEXT NOT NULL,
    payload JSONB,
    retries INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);