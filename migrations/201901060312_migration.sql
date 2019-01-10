create database library with owner postgres encoding 'UTF8';

create table author (
	id serial primary key,
	first_name varchar(50) not null,
	lastname varchar(50) not null,
	email varchar(200) not null,
	created_at timestamp not null default now(),
	modified_at timestamp not null default now(),
	status smallint not null default 1,
	constraint uq_email unique (email)
);

create table book (
	id serial primary key,
	title varchar(300) not null,
	isbn varchar(13) not null,
	about varchar(500),
	edition smallint not null default 1,
	page_number smallint null,
	release_date timestamp not null,
	author_id bigint not null,
	created_at timestamp not null default now(),
	modified_at timestamp not null default now(),
	status smallint not null default 1,
	constraint uq_isbn unique (isbn),
	constraint fk_book_author_id_author_id foreign key (author_id) references author (id)
);

insert into author (first_name, last_name, email) values ('Mina', 'Andrawos', 'mina.andrawos@gmail.com');
insert into author (first_name, last_name, email) values ('Martin', 'Helmich', 'martin.helmich@gmail.com');
insert into author (first_name, last_name, email) values ('Quanyi', 'Ma', 'quanyi.ma@gmail.com');
insert into author (first_name, last_name, email) values ('Karl', 'Isenberg', 'karl.isenberg@gmail.com');
