use emotion_sns;

CREATE TABLE emotions (
    id INT AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL UNIQUE,
    PRIMARY KEY (id)
);

INSERT INTO emotions (name) VALUES("Happy");
INSERT INTO emotions (name) VALUES("Angry");
INSERT INTO emotions (name) VALUES("Sad");
INSERT INTO emotions (name) VALUES("Funny");
INSERT INTO emotions (name) VALUES("Emotionless");