create table users
(
    id         int auto_increment,
    nik        varchar(16) not null,
    role       varchar(255) not null,
    password   varchar(255) not null,
    created_at timestamp    null,
    updated_at timestamp    null,
    constraint users_pk
        primary key (id),
    constraint users_unique_nik
        unique (nik)
);

