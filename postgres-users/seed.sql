CREATE TABLE users (
	name	varchar(50),
	age	integer,
	CONSTRAINT pk_user UNIQUE (name)
);

INSERT INTO users (name, age) values ('test1', 11);
INSERT INTO users (name, age) values ('test2', 15);
INSERT INTO users (name, age) values ('test3', 21);
INSERT INTO users (name, age) values ('test4', 33);

