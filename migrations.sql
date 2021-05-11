create table human_resources
(
	id serial not null,
	name varchar(256) not null,
	dob date,
	address varchar(256) not null,
	description varchar(512) not null,
	created_at date not null,
	updated_at date,
	visibility bool not null default true
);

create unique index human_resources_id_uindex
	on human_resources (id);

alter table human_resources
	add constraint table_name_pk
		primary key (id);