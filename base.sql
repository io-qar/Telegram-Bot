
CREATE DATABASE tg_bot

CREATE TABLE users(id INTEGER PRIMARY KEY, username TEXT NOT NULL);

CREATE TABLE requests (id_request SERIAL PRIMARY KEY,request TEXT,id_user INTEGER,FOREIGN KEY (id_user) REFERENCES users(id));