BEGIN;

CREATE TABLE IF NOT EXISTS projects(
    id bigserial PRIMARY KEY NOT NULL,
    project_name varchar(500) NOT NULL,
    description varchar(1000)
);

CREATE TABLE IF NOT EXISTS columns(
    id bigserial PRIMARY KEY,
    column_name varchar(255) NOT NULL UNIQUE,
    index integer NOT NULL,
    project_id bigserial REFERENCES projects (id)
    ON DELETE CASCADE ON UPDATE CASCADE NOT NULL
);

CREATE TABLE IF NOT EXISTS tasks(
    id bigserial PRIMARY KEY NOT NULL,
    name varchar(500) NOT NULL,
    description varchar(5000),
    index integer NOT NULL,
    column_id bigserial REFERENCES columns (id)
    ON DELETE CASCADE ON UPDATE CASCADE NOT NULL
);

CREATE TABLE IF NOT EXISTS comments(
    id bigserial PRIMARY KEY NOT NULL,
    comment_text varchar(5000) NOT NULL,
    task_id bigserial REFERENCES tasks (id)
    ON DELETE CASCADE ON UPDATE CASCADE NOT NULL
);

COMMIT;