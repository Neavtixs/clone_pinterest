CREATE TABLE follow (
	id VARCHAR(100) NOT NULL,
	follower_id VARCHAR(100) NOT NULL,
	following_id VARCHAR(100) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (id),
	CONSTRAINT fk_follower_user FOREIGN KEY (follower_id) REFERENCES users (id),
	CONSTRAINT fk_following_user FOREIGN KEY (following_id) REFERENCES users (id)
);