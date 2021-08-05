CREATE TABLE catalogs
(
    id         uuid         not null constraint PK_CATALOG primary key,
    name       varchar(255) not null,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE products
(
    id         uuid         not null constraint PK_PRODUCT primary key,
    amount     bigint       not null,
    number     varchar(255) not null,
    catalog_id uuid         not null constraint FK_CATALOG references catalogs,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

INSERT INTO catalogs VALUES ('784e0c11-3aa7-418d-9546-24d9183b73fe', 'Books', null, null, null);
INSERT INTO catalogs VALUES ('9e8c4c6e-bf4f-4a20-9bf0-4815c0f5a383', 'Games', null, null, null);

INSERT INTO products VALUES (gen_random_uuid (), 12, '1a2d', '784e0c11-3aa7-418d-9546-24d9183b73fe', null, null, null);
INSERT INTO products VALUES (gen_random_uuid (), 120, '1a2e', '784e0c11-3aa7-418d-9546-24d9183b73fe', null, null, null);

INSERT INTO products VALUES (gen_random_uuid (), 1, '2a2d', '9e8c4c6e-bf4f-4a20-9bf0-4815c0f5a383', null, null, null);
INSERT INTO products VALUES (gen_random_uuid (), 2, '3a2d', '9e8c4c6e-bf4f-4a20-9bf0-4815c0f5a383', null, null, null);