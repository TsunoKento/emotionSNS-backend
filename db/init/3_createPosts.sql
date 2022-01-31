use emotion_sns;

CREATE TABLE posts (
    id INT AUTO_INCREMENT,
    user_id INT NOT NULL,
    content VARCHAR(255) NOT NULL,
    image VARCHAR(255),
    published_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    emotion_id INT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (emotion_id) REFERENCES emotions(id)
);

INSERT INTO posts (user_id, content, image, emotion_id) VALUES(1, "今日も猫は可愛い", "https://source.unsplash.com/IuJc2qh2TcA", 1);