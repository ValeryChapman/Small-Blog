CREATE TABLE articles
(
    id         serial       not null unique,
    title      varchar(255) not null,
    text       varchar      not null,
    views      int          not null default 0,
    created_at timestamp    with time zone not null     
);