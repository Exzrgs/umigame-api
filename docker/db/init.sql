CREATE TABLE IF NOT EXISTS problems (
    id INTEGER UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    statement TEXT NOT NULL,
    answer TEXT NOT NULL,
    author VARCHAR(100) NOT NULL,
    reference VARCHAR(100) NOT NULL,
    reference_url VARCHAR(500) NOT NULL,
    created_at DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    id INTEGER UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(100) NOT NULL,
    uuid VARCHAR(100),
    is_valid BOOLEAN NOT NULL,
    created_at DATETIME
);

CREATE TABLE IF NOT EXISTS users_problems (
    problem_id INTEGER UNSIGNED AUTO_INCREMENT UNIQUE NOT NULL,
    user_id INTEGER UNSIGNED AUTO_INCREMENT UNIQUE NOT NULL,
    is_solved BOOLEAN NOT NULL,
    is_liked BOOLEAN NOT NULL,
    FOREIGN KEY problem_id REFERENCES problems(id),
    FOREIGN KEY user_id REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS chats (
    id INTEGER PRIMARY KEY,
    problem_id INTEGER UNIQUE NOT NULL,
    user_id INTEGER UNIQUE NOT NULL,
    question TEXT NOT NULL,
    answer TEXT NOT NULL,
    created_at DATETIME NOT NULL,
    FOREIGN KEY problem_id REFERENCES problems(id),
    FOREIGN KEY user_id REFERENCES users(id)
);