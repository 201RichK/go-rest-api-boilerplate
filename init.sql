CREATE USER x_user WITH PASSWORD 'db1234@x2024';
ALTER USER x_user WITH SUPERUSER;
ALTER ROLE x_user CREATEROLE CREATEDB;

CREATE DATABASE my-db;
GRANT ALL PRIVILEGES ON DATABASE my-db to x_user;