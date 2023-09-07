-- Create the GitHub events table
CREATE TABLE IF NOT EXISTS github_events (
    id serial PRIMARY KEY,
    event_type varchar(255),
    actor varchar(255),
    repo_url varchar(255),
    created_at timestamp
);

-- Create an index on the 'created_at' column for performance optimization
CREATE INDEX IF NOT EXISTS idx_created_at ON github_events(created_at);

-- Create a table for actors (GitHub users)
CREATE TABLE IF NOT EXISTS github_actors (
    id serial PRIMARY KEY,
    login varchar(255) UNIQUE
);

-- Create a table for repositories
CREATE TABLE IF NOT EXISTS github_repositories (
    id serial PRIMARY KEY,
    url varchar(255) UNIQUE
);

-- Create a table for emails
CREATE TABLE IF NOT EXISTS github_emails (
    id serial PRIMARY KEY,
    email varchar(255) UNIQUE
);
