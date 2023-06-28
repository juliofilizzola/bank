CREATE TABLE accounts (
    id int auto_increment primary key,
    owner varchar(191) not null,
    balance BIGINT not null,
    currency varchar(191) not null,
    created_at timestamp not null DEFAULT current_timestamp()
);

create table entries (
    id int auto_increment primary key,
    account_id int not null,
    amount bigint not null,
    created_at timestamp not null DEFAULT current_timestamp()
);

create table transfers (
    id int auto_increment primary key,
    from_account_id int not null,
    to_account_id int not null,
    amount bigint not null,
    created_at timestamp not null DEFAULT current_timestamp()
);


alter table entries add foreign key (account_id) references accounts (id);
alter table transfers add foreign key (from_account_id) references accounts (id);
alter table transfers add foreign key (to_account_id) references accounts (id);

