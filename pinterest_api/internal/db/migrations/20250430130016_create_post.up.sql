CREATE TABLE post (
    id VARCHAR(100) NOT NULL,
    user_id VARCHAR(100) NOT NULL,
    title VARCHAR(100),
    description VARCHAR(100),
    image VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (image),
    CONSTRAINT fk_user_post FOREIGN KEY (user_id) REFERENCES users (id)
);
