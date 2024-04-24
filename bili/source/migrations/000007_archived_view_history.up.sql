create table archived_view_history
(
    bvid       varchar(64)       primary key ,
    create_time datetime(3) null,
    update_time datetime(3) null,
    delete_time datetime(3) null,

    title       text      null comment 'video title',
    cover       varchar(256)     null comment 'video cover',
    upper_mid   bigint           not null default 0  comment 'video upper mid',
    upper_name  varchar(256)      not null default 0  comment 'video upper name',
    face_name  varchar(256)      not null default 0  comment 'video upper face',
    duration    bigint           not null default 0  comment 'video duration',
    view_at    bigint           not null comment 'video view_at',
    videos bigint           not null default 0 comment 'videos',
    progress bigint           not null default 0 comment 'progress',
    is_finish tinyint(1)       not null default 0 comment 'is_finish',
    kid  bigint           comment '稿件avid',
    tag_name varchar(32)        comment 'video tag_name',

    resp        json             null
) comment '浏览历史记录';
