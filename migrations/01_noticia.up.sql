create table if not exists tb_noticia(
    id varchar(36) not null,
    titulo varchar(255) not null,
    created_at timestamp not null DEFAULT Now(),
    updated_at timestamp,
    deleted_at timestamp,
    constraint pk_noticia primary key (id)
)