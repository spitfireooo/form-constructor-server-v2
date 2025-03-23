CREATE TABLE users
(
    id serial not null unique,
    email varchar(255) not null unique,
    phone varchar(255),
    address varchar(255),
    password varchar(255) not null,
    nickname varchar(255),
    logo varchar(255),
    created_at timestamp with time zone default now() not null,
    updated_at timestamp with time zone default now() not null
);

CREATE TABLE tokens
(
    id serial not null unique,
    user_id int not null,
    token varchar(255) not null,

    foreign key (user_id) references users(id)
);

CREATE TABLE user_permissions
(
    id serial not null unique,
    user_id int not null,
    permission varchar(20) not null,

    foreign key (user_id) references users(id)
);

CREATE TABLE tags
(
    id serial not null unique,
    title varchar(255) not null,
    description text
);

CREATE TABLE user_tags
(
    id serial not null unique,
    user_id int not null,
    tag_id int not null,

    foreign key (user_id) references users(id),
    foreign key (tag_id) references tags(id)
);

CREATE TABLE forms
(
    id serial not null unique,
    title varchar(255) not null,
    slug varchar(255) not null unique,
    description text,
    logo varchar(255),
    author_id int not null,

    foreign key (author_id) references users(id)
);

CREATE TABLE fields
(
    id serial not null unique,
    form_id int not null,
    type varchar(255) not null,
    label varchar(255) not null,
    name varchar(255) not null,
    order_of int default 0,
    required boolean not null default false,

    foreign key (form_id) references forms(id)
);

CREATE TABLE field_variants
(
    id serial not null unique,
    field_id int not null,
    variant varchar(255) not null,
    name varchar(255) not null,

    foreign key (field_id) references fields(id)
);

CREATE TABLE field_multiply
(
    id serial not null unique,
    field_id int not null,
    is_multiply boolean not null,

    foreign key (field_id) references fields(id)
);

CREATE TABLE field_placeholder
(
    id serial not null unique,
    field_id int not null,
    placeholder varchar(255) not null,

    foreign key (field_id) references fields(id)
);

CREATE TABLE field_range
(
    id serial not null unique,
    field_id int not null,
    min int not null default 1,
    max int not null default 20,

    foreign key (field_id) references fields(id)
);

CREATE TABLE results
(
    id serial not null unique,
    field_id int not null,
    value varchar(255) not null,

    foreign key (field_id) references fields(id)
);

CREATE TABLE dialogs
(
    id serial not null unique,
    user1_id int not null,
    user2_id int not null,

    foreign key (user1_id) references users(id),
    foreign key (user2_id) references users(id)
);

CREATE TABLE messages
(
    id serial not null unique,
    dialog_id int not null,
    author_id int not null,

    foreign key (dialog_id) references dialogs(id),
    foreign key (author_id) references users(id)
);