DROP TABLE IF EXISTS requests;

create table requests (
    id serial primary key,
    host TEXT,
    path TEXT,
    method TEXT,
    headers TEXT,
    body TEXT,
    params TEXT NOT NULL,
    cookies TEXT
);
