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
    uuid VARCHAR(100) UNIQUE,
    is_valid BOOLEAN NOT NULL,
    created_at DATETIME
);

CREATE TABLE IF NOT EXISTS activities (
    problem_id INTEGER UNSIGNED NOT NULL,
    user_id INTEGER UNSIGNED NOT NULL,
    is_solved BOOLEAN DEFAULT FALSE NOT NULL,
    is_liked BOOLEAN DEFAULT FALSE NOT NULL,
    FOREIGN KEY (problem_id) REFERENCES problems(id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    PRIMARY KEY (problem_id, user_id)
);

CREATE TABLE IF NOT EXISTS chats (
    id INTEGER UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    problem_id INTEGER UNSIGNED NOT NULL,
    user_id INTEGER UNSIGNED NOT NULL,
    question TEXT NOT NULL,
    answer INTEGER NOT NULL,
    created_at DATETIME NOT NULL,
    FOREIGN KEY (problem_id) REFERENCES problems(id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    INDEX user_problem_index (problem_id, user_id)
);

INSERT INTO problems (title, statement, answer, author, reference, reference_url, created_at)
VALUES (
    "test1", 
    "test1",
    "test1",
    "test1",
    "test1",
    "test1",
    "2023-01-28 12:00:00"
);

INSERT INTO problems (title, statement, answer, author, reference, reference_url, created_at)
VALUES (
    "test2", 
    "test2",
    "test2",
    "test2",
    "test2",
    "test2",
    "2023-01-28 12:00:00"
);

INSERT INTO problems (title, statement, answer, author, reference, reference_url, created_at)
VALUES (
    "test3", 
    "test3",
    "test3",
    "test3",
    "test3",
    "test3",
    "2023-01-28 12:00:00"
);

INSERT INTO users (name, email, password_hash, uuid, is_valid, created_at)
VALUES (
    "test1",
    "test1",
    "test1",
    "test1",
    TRUE,
    "2023-01-28 12:00:00"
);

INSERT INTO activities (problem_id, user_id, is_solved, is_liked)
VALUES (
    1,
    1,
    TRUE,
    TRUE
);

INSERT INTO activities (problem_id, user_id, is_solved, is_liked)
VALUES (
    2,
    1,
    TRUE,
    FALSE
);

INSERT INTO activities (problem_id, user_id, is_solved, is_liked)
VALUES (
    3,
    1,
    FALSE,
    TRUE
);