------------------Model---------------------------

create table users (
	user_id serial not null primary key,
	firstname varchar(64) not null,
	lastname varchar(64) not null,
	balance decimal(18,2),
	email varchar(256) not null,
	password varchar(64) not null
);

comment on table users is 'Foydalanuvchilar';

create table categories (
	category_id serial not null primary key,
	name varchar(256)
);

comment on table categories is 'Kategoriyalar';

create table incomes (
	income_id serial not null primary key,
	comment text default null,
	amount decimal (18,2),
	created_at timestamp default current_timestamp,
	user_id int not null,
	category_id int references categories (category_id)
);

comment on table incomes is 'Kirimlar';

create table outgoings (
	outgoing_id serial not null primary key,
	comment text default null,
	amount decimal (18,2),
	created_at timestamp default current_timestamp,
	user_id int not null,
	category_id int references categories (category_id)
);

comment on table outgoings is 'Harajatlar';



----------------------Mock data----------------------------


insert into users (firstname, lastname, balance, email, password) values ('Jamshid', 'Ismoil', 526000.00, 'example1@gmail.com', '568796');
insert into users (firstname, lastname, balance, email, password) values ('Yusuf', 'Ismoil', 452000.00, 'example2@gmail.com', '434445');
insert into users (firstname, lastname, balance, email, password) values ('Muhammad', 'Umar', 162400.00, 'example3@gmail.com', '234745');

insert into categories (name) values ('daily spendings'), ('communal payments'), ('wages');

insert into incomes (amount, user_id, category_id) values (5000, 1,3);

insert into outgoings (amount, user_id, category_id) values(3000,1,1),(100000, 1,2),(14000, 1,1);

insert into incomes (amount, user_id, category_id) values (12000, 2,3);

insert into outgoings (amount, user_id, category_id) values(3000,2,1),(140000, 2,2),(12000, 2,1);
