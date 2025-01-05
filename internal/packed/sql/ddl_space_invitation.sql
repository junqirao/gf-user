create table if not exists space_invitation
(
    id         int auto_increment,
    space      int         not null,
    `from`     varchar(50) not null,
    status     int         not null comment '0: create, 1: accept, 2: reject, 3: cancel',
    target     varchar(50) not null,
    comment    varchar(50) null,
    created_at datetime    null,
    primary key (id),
    unique index uk_space_target (space, target),
    index idx_from (`from`),
    index idx_target (target)
);

