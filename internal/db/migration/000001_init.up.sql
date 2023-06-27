CREATE TABLE accounts (
    id bigserial PRIMARY KEY,
    owner varchar not null,
    balance bigint not null,
    currency varchar not null,
    created_at timestamp not null DEFAULT (now())
);

create table entries (
    id bigserial primary key,
    account_id bigint not null,
    amount bigint not null,
    created_at timestamp not null DEFAULT (now())
);

create table transfers (
    id bigserial primary key,
    from_account_id bigint not null,
    to_account_id bigint not null,
    amount bigint not null,
    created_at timestamp not null DEFAULT (now())
);


alter table entries add foreign key (account_id) references accounts (id);
alter table transfers add foreign key (from_account_id) references accounts (id);
alter table transfers add foreign key (to_account_id) references accounts (id);

create index on accounts (owner);
create index on entries (account_id);
create index on transfers (from_account_id);
create index on transfers (to_account_id);
create index on transfers (from_account_id, to_account_id);

comment on column entries.amount is 'can be negative or positive';
comment on column transfers.amount is 'must be positive';
