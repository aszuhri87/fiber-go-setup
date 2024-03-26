CREATE TABLE users (
   id UUID primary key,
   name VARCHAR(255),
   username VARCHAR(255),
   password VARCHAR(255),
   created_at TIMESTAMP default now(),
   updated_at TIMESTAMP default now(),
   deleted_at TIMESTAMP default now()
);
