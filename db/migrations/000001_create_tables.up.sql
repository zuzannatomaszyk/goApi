CREATE TABLE IF NOT EXISTS users(
    userID SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL 
);

CREATE TABLE IF NOT EXISTS messages(
    messageID SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL REFERENCES users (username), 
    messageContent TEXT,
    createdAt NUMERIC DEFAULT extract(epoch from now())
);
