create table if not exists app
(
    id           varchar(50)  not null comment 'app uuid (app_id)',
    secret       varchar(50)  not null comment 'secret md5',
    name         varchar(50)  not null,
    space        int          not null,
    descriptions varchar(100) null,
    profile      json         null,
    expired_at   datetime     null,
    created_at   datetime     null,
    primary key (id),
    index idx_space (space)
);

