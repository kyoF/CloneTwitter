CREATE TABLE likes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL UNIQUE,
    tweet_id CHAR(36) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (tweet_id, user_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
    FOREIGN KEY (tweet_id) REFERENCES tweets(tweet_id)
);

