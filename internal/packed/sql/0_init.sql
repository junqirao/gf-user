create table if not exists account
(
    id         varchar(50)   not null comment 'account uuid',
    account    varchar(50)   not null comment 'unique account',
    password   varchar(50)   null comment 'password hash',
    type       int default 0 not null comment '0: normal, 1: app',
    status     int default 0 not null comment '0: normal, 1: frozen',
    name       varchar(20)   null,
    email      varchar(50)   null,
    avatar     varchar(50)   null,
    created_at datetime      null,
    update_at  datetime      null,
    extra      json          null comment 'extra',
    primary key (id),
    unique index uk_account (account),
    index idx_acc_type_index (account, type)
);

create table if not exists space
(
    id          int auto_increment,
    name        varchar(50)  not null comment 'space name',
    logo        varchar(50)  null comment 'space logo',
    owner       varchar(50)  null comment 'owner account id',
    description varchar(200) null,
    profile     json         null,
    update_at   datetime     null,
    created_at  datetime     null,
    primary key (id),
    unique index uk_space_name (name)
);

create table if not exists user
(
    id         int auto_increment comment 'user id',
    account    varchar(50) not null comment 'account id',
    space      int         not null comment 'space id',
    type       int         null comment '0: normal, 1: manager',
    name       varchar(50) null,
    update_at  datetime    null,
    created_at datetime    null,
    primary key (id),
    index idx_space_account (account, space)
);

