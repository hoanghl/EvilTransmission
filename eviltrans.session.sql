CREATE EXTENSION vector;


CREATE TABLE items(
    id          bigserial   primary key,
    embedding   vector(3)
);

select * from items;