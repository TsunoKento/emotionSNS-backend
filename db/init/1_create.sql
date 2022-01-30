CREATE DATABASE emotion_sns;

use emotion_sns;

CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT,
    third_party_id VARCHAR(30) NOT NULL,
    user_id VARCHAR(50) NOT NULL,
    name VARCHAR(255),
    image VARCHAR(255),
    email VARCHAR(255),
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

INSERT INTO users (third_party_id, user_id, name, image, email) VALUES("123456789012", "yamada_taroooo", "太郎Y", "https://source.unsplash.com/p-6S-filXzM", "tarou_yamada@mail.com")