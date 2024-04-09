CREATE TABLE test.task_status (
	id int8 GENERATED ALWAYS AS IDENTITY( INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL,
	task_statu varchar(50) NULL,
	created_at timestamp DEFAULT now() NULL,
	updated_at timestamp DEFAULT now() NULL,
	is_deleted bool DEFAULT false NULL,
	create_user_id int8 NULL,
	update_user_id int8 NULL,
	CONSTRAINT task_status_pk PRIMARY KEY (id)
);

CREATE TABLE test.tasks (
	id int8 GENERATED ALWAYS AS IDENTITY( INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL,
	title varchar(200) NULL,
	description text NULL,
	status_id int8 NULL,
	created_at timestamp DEFAULT now() NULL,
	updated_at timestamp DEFAULT now() NULL,
	is_deleted bool DEFAULT false NULL,
	create_user_id int8 NULL,
	update_user_id int8 NULL,
	difficulty int8 NULL,
	duration int8 DEFAULT 1 NULL,
	CONSTRAINT fk_status_id FOREIGN KEY (status_id) REFERENCES test.task_status(id)
);



CREATE TABLE test.users (
	id int8 GENERATED ALWAYS AS IDENTITY( INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL,
	first_name varchar NULL,
	last_name varchar NULL,
	email varchar NULL,
	"password" varchar NULL,
	created_at timestamp DEFAULT now() NULL,
	updated_at timestamp DEFAULT now() NULL,
	is_deleted bool DEFAULT false NULL,
	create_user_id int8 NULL,
	update_user_id int8 NULL,
	user_level int8 NULL,
	CONSTRAINT users_pk PRIMARY KEY (id)
);