
CREATE TABLE save (
    id VARCHAR(100) NOT NULL,
    user_id VARCHAR(100) NOT NULL,
    post_id VARCHAR(100) NOT NULL,
    created_at BIGINT NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_save_user FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT fk_save_post FOREIGN KEY (post_id) REFERENCES post (id) ON DELETE CASCADE
);