-- +migrate Up
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL,
    description TEXT,
    image_url VARCHAR(512),
    release_year INTEGER,
    price INTEGER,
    total_page INTEGER,
    thickness VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by VARCHAR(255)
);

-- Insert default admin user if not exists
INSERT INTO users (username, password, created_by)
SELECT 'admin', 'admin123', 'system'
WHERE NOT EXISTS (SELECT 1 FROM users WHERE username = 'admin');

-- Insert default categories
INSERT INTO categories (name, created_by)
SELECT 'Technology', 'system'
WHERE NOT EXISTS (SELECT 1 FROM categories WHERE name = 'Technology');

INSERT INTO categories (name, created_by)
SELECT 'Science', 'system'
WHERE NOT EXISTS (SELECT 1 FROM categories WHERE name = 'Science');

-- +migrate Down
DROP TABLE IF EXISTS books;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS categories;
