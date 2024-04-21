create table if not exists public.expressions
(
    id            serial
        constraint expressions_pk
            primary key,
    user_id       integer                             not null,
    expression    text                                not null,
    result        bigint,
    is_processing boolean   default false             not null,
    is_time_limit boolean,
    is_valid      boolean,
    is_finished   boolean   default false,
    time_limit    integer   default 200               not null,
    created_at    timestamp default CURRENT_TIMESTAMP not null,
    finished_at   timestamp
);

comment on column public.expressions.user_id is 'user_id which expression belongs to';

comment on column public.expressions.expression is 'expression to calculate';

comment on column public.expressions.result is 'result of expression';

comment on column public.expressions.is_processing is 'is calculations in processing';

comment on column public.expressions.is_valid is 'is expression valid or not';

comment on column public.expressions.time_limit is 'time limit to calculate an expression';

alter table public.expressions
    owner to expression_user;

