create table if not exists public.users
(
    id       serial
        primary key,
    login    varchar(255) not null
        unique,
    password varchar(255) not null
);

alter table public.users
    owner to expression_user;