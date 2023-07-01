CREATE TABLE ACCOUNTS (
                          id INTEGER auto_increment primary key,
                          owner varchar(191) not null,
                          balance BIGINT not null,
                          currency varchar(191) not null,
                          created_at timestamp not null DEFAULT current_timestamp()
);

create table ENTRIES (
                         id INTEGER auto_increment primary key,
                         account_id INTEGER not null,
                         amount bigint not null,
                         created_at timestamp not null DEFAULT current_timestamp(),
                         FOREIGN KEY(account_id) REFERENCES ACCOUNTS(id)
);

create table TRANSFERS (
                           id INTEGER auto_increment primary key,
                           from_account_id INTEGER not null,
                           to_account_id INTEGER not null,
                           amount bigint not null,
                           created_at timestamp not null DEFAULT current_timestamp(),
                           FOREIGN KEY(from_account_id) REFERENCES ACCOUNTS(id),
                           FOREIGN KEY(to_account_id) REFERENCES ACCOUNTS(id)
);
