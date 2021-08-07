CREATE TABLE todos
(
    id       uuid not null constraint PK_TODO primary key,
    title    varchar(255),
    complete bool,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

INSERT INTO todos VALUES ('47ad1ced-df3d-461b-b914-04213331cc36', 'Harry potter', false);
INSERT INTO todos VALUES ('294d6fe3-9cc9-4fa3-9eda-9f70d84e83a6', 'Star wars', true);