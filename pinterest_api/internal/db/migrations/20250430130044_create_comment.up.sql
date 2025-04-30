
CREATE TABLE comment (
    id VARCHAR(100) NOT NULL,
    comment VARCHAR(100) NOT NULL,
    user_id VARCHAR(100) NOT NULL,
    post_id VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT fk_comment_user FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT fk_comment_post FOREIGN KEY (post_id) REFERENCES post (id) ON DELETE CASCADE
);
