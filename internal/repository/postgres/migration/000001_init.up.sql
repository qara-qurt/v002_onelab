CREATE TABLE transactionBook (
    id Serial PRIMARY KEY not null unique,
    user_id int not null,
    book_id int not null,
    price varchar(255) not null
);