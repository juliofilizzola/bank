CREATE TABLE accounts (
    id BIGINT PRIMARY KEY,
    owner varchar(191) not null,
    balance BIGINT not null,
    currency varchar(191) not null,
    created_at timestamp not null DEFAULT (now())
);

create table entries (
    id BIGINT primary key,
    account_id bigint not null,
    amount bigint not null,
    created_at timestamp not null DEFAULT (now())
);

create table transfers (
    id bigint primary key,
    from_account_id bigint not null,
    to_account_id bigint not null,
    amount bigint not null,
    created_at timestamp not null DEFAULT (now())
);


alter table entries add foreign key (account_id) references accounts (id);
alter table transfers add foreign key (from_account_id) references accounts (id);
alter table transfers add foreign key (to_account_id) references accounts (id);

