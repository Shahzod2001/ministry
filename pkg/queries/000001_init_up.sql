create table cities
(
    id   serial primary key,
    name varchar(255) not null
);

create table universities
(
    id         serial primary key,
    name       varchar(255) not null,
    login      varchar(100) not null unique,
    password   varchar(100) not null,
    city_id    int references cities,
    is_active  boolean   default false,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp,
    deleted_at timestamp
);

create table academic_degrees
(
    id   serial primary key,
    name varchar(255) not null
);

create table academic_positions
(
    id   serial primary key,
    name varchar(255) not null
);

create table teacher_specialities
(
    id   serial primary key,
    name varchar(255) not null
);

create table teacher_types
(
    id   serial primary key,
    name varchar(255) not null
);

create table direction_specialities
(
    id   serial primary key,
    name varchar(255) not null
);

create table genders
(
    id   serial primary key,
    name varchar(255) not null
);

create table admins
(
    id          serial primary key,
    last_name   varchar(255),
    first_name  varchar(255) not null,
    middle_name varchar(255),
    login       varchar(100) not null unique,
    password    varchar(100) not null
);

insert into admins (last_name, first_name, middle_name, login, password)
values ('Админов', 'Админ', 'Админович', 'admin', '$2y$10$dkJd2LVyrwJbYanxmB1r4eFrh0S478PYCyffeBkDufLajnvvY9i1a');

create table teachers
(
    id                   serial primary key,
    last_name            varchar(255),
    first_name           varchar(255) not null,
    middle_name          varchar(255),
    birth_date           date,
    birth_place          varchar(255),
    gender               int references genders,
    university_id        int references universities,
    academic_degree_id   int references academic_degrees,
    academic_position_id int references academic_positions,
    spec_id              int references teacher_specialities,
    direction_spec_id    int references direction_specialities,
    type_id              int references teacher_types,
    job_title            varchar(100),
    other_job            varchar(100),
    from_year            int,
    to_year              int,
    is_active            boolean   default true,
    created_at           timestamp default CURRENT_TIMESTAMP,
    updated_at           timestamp,
    deleted_at           timestamp
);

create table uni_specializations
(
    id            serial primary key,
    name          varchar(150),
    code          varchar(30),
    university_id int references universities,
    from_year     int,
    to_year       int,
    created_at    timestamp default CURRENT_TIMESTAMP,
    updated_at    timestamp,
    deleted_at    timestamp
);

create table budjet
(
    id                  serial primary key,
    university_id       int references universities,
    overall             decimal,
    scientists_award    decimal,
    scientists_research decimal,
    science_events      decimal,
    books_print         decimal,
    articles_print      decimal,
    journals_print      decimal,
    award               decimal,
    tech_parks          decimal,
    grants              decimal
);

create table universities
(
    id         serial primary key,
    name       varchar(255) not null,
    login      varchar(100) not null unique,
    password   varchar(100) not null,
    city_id    int references cities,
    is_active  boolean   default true,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp,
    deleted_at timestamp
);

insert into cities (name)
values ('Душанбе'),
       ('Худжанд'),
       ('Кулоб'),
       ('Турсунзода'),
       ('Гиссар');

insert into academic_degrees (name)
values ('Кандидат наук'),
       ('Доктор наук'),
       ('Доцент'),
       ('Профессор'),
       ('Магистр'),
       ('Бакалавр');
insert into academic_positions (name)
values ('Старший преподаватель'),
       ('Преподаватель'),
       ('Ассистент'),
       ('Доцент'),
       ('Профессор');
insert into teacher_specialities (name)
values ('Информационные технологии'),
       ('Экономика'),
       ('Медицина'),
       ('Педагогика'),
       ('Юриспруденция');
insert into teacher_types (name)
values ('Штатный'),
       ('Внештатный');
insert into direction_specialities (name)
values ('Информационные технологии'),
       ('Экономика'),
       ('Медицина'),
       ('Педагогика'),
       ('Юриспруденция');

